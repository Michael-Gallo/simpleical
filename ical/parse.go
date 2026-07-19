// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package ical

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/michael-gallo/simpleical/internal/icalerr"
	"github.com/michael-gallo/simpleical/model"
)

// parserState represents the current parsing state using a single integer.
type parserState uint8

const (
	stateCalendar parserState = iota
	stateEvent
	stateTimezone
	stateTodo
	stateJournal
	stateFreebusy
	stateAlarm
	stateStandard
	stateDaylight
	stateFinished
)

// beginVCalendarLine is the content line that opens a VCALENDAR object.
const beginVCalendarLine = "BEGIN:VCALENDAR"

// componentCursor caches pointers to the active component so property handlers
// avoid repeated &slice[len-1] indexing on every line.
type componentCursor struct {
	event       *model.Event
	todo        *model.Todo
	journal     *model.Journal
	freeBusy    *model.FreeBusy
	timeZone    *model.TimeZone
	alarm       *model.Alarm
	alarmParent parserState // stateEvent or stateTodo; valid while in stateAlarm
	tzProp      *model.TimeZoneProperty
}

// ReadSingle takes an io.Reader containing iCalendar data and parses it into a Calendar.
// It asserts that the input contains exactly one VCALENDAR object; any content after
// END:VCALENDAR (including another BEGIN:VCALENDAR) returns ErrContentAfterEndBlock.
// Use Read to parse a stream that may contain multiple VCALENDAR objects.
func ReadSingle(reader io.Reader) (*model.Calendar, error) {
	calendars, err := Read(reader)
	if err != nil {
		return nil, err
	}
	if len(calendars) != 1 {
		return nil, icalerr.ErrContentAfterEndBlock
	}
	return calendars[0], nil
}

// Read takes an io.Reader containing iCalendar data and parses it into a
// slice of Calendars. Per RFC 5545 section 3.4, a single iCalendar stream may
// contain multiple sequential VCALENDAR objects. Scanner state and internal
// buffers are shared across calendars, so parsing N calendars costs no more
// than parsing each individually.
// Use ReadSingle to assert that the input contains exactly one VCALENDAR object.
func Read(reader io.Reader) ([]*model.Calendar, error) {
	// Reusable parameter map to avoid allocations on every property
	reusableParams := make(map[string]string, 2)
	scanner := bufio.NewScanner(reader)
	var pending string
	var hasPending bool
	var calendars []*model.Calendar

	for {
		line, ok := nextLogicalLine(scanner, &pending, &hasPending)
		if !ok {
			break
		}
		if line != beginVCalendarLine {
			if line == "" {
				return nil, icalerr.ErrInvalidCalendarEmptyLine
			}
			if len(calendars) > 0 {
				return nil, icalerr.ErrContentAfterEndBlock
			}
			return nil, icalerr.ErrInvalidCalendarFormatMissingBegin
		}

		calendar, err := parseOneCalendar(scanner, &pending, &hasPending, reusableParams)
		if err != nil {
			return nil, err
		}
		calendars = append(calendars, calendar)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading iCalendar data: %w", err)
	}
	if len(calendars) == 0 {
		return nil, icalerr.ErrNoCalendarFound
	}

	return calendars, nil
}

// parseOneCalendar parses a single VCALENDAR body, starting just after its
// BEGIN:VCALENDAR line, and returns as soon as END:VCALENDAR is validated.
// The scanner, line-unfolding state, and parameter map are owned by the caller
// so multiple calendars can be parsed from the same stream without
// reallocating them.
func parseOneCalendar(scanner *bufio.Scanner, pending *string, hasPending *bool, reusableParams map[string]string) (*model.Calendar, error) {
	calendar := &model.Calendar{}
	currentState := stateCalendar
	var cursor componentCursor

	for {
		line, ok := nextLogicalLine(scanner, pending, hasPending)
		if !ok {
			break
		}

		if line == "" {
			return nil, icalerr.ErrInvalidCalendarEmptyLine
		}

		clear(reusableParams)

		propertyName, params, value, err := parseIcalLineWithReusableMap(line, reusableParams)
		if err != nil {
			return nil, err
		}
		switch propertyName {
		case "BEGIN":
			if err := handleBeginBlock(value, &currentState, calendar, &cursor); err != nil {
				return nil, err
			}
		case "END":
			if err := handleEndBlock(value, &currentState, calendar, &cursor); err != nil {
				return nil, err
			}
			if currentState == stateFinished {
				return calendar, nil
			}
		default:
			if err := parsePropertyLine(propertyName, value, params, currentState, calendar, &cursor); err != nil {
				return nil, err
			}
		}
	}

	// Check for scanner errors
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading iCalendar data: %w", err)
	}

	// Input ended before END:VCALENDAR was seen
	return nil, icalerr.ErrInvalidCalendarFormatMissingEnd
}

// nextLogicalLine returns the next RFC 5545 content line after unfolding,
// with trailing spaces trimmed.
// Folded physical lines (CRLF followed by a single SPACE or HTAB) are joined
// by stripping that leading white-space and appending the remainder.
func nextLogicalLine(scanner *bufio.Scanner, pending *string, hasPending *bool) (string, bool) {
	var line string
	if *hasPending {
		line = *pending
		*hasPending = false
	} else if scanner.Scan() {
		line = scanner.Text()
	} else {
		return "", false
	}

	var b strings.Builder
	folded := false
	for scanner.Scan() {
		next := scanner.Text()
		if len(next) > 0 && (next[0] == ' ' || next[0] == '\t') {
			if !folded {
				b.Grow(len(line) + len(next))
				b.WriteString(line)
				folded = true
			}
			b.WriteString(next[1:])
			continue
		}
		*pending = next
		*hasPending = true
		break
	}
	if folded {
		return strings.TrimRight(b.String(), " "), true
	}
	return strings.TrimRight(line, " "), true
}

// parsePropertyLine parses a single property line and adds it to the appropriate component based on current state.
func parsePropertyLine(propertyName string, value string, params map[string]string, currentState parserState, calendar *model.Calendar, cursor *componentCursor) error {
	switch currentState {
	case stateAlarm:
		return parseAlarmProperty(propertyName, value, params, cursor.alarm)
	case stateEvent:
		return parseEventProperty(propertyName, value, params, cursor.event)
	case stateTimezone:
		return parseTimezoneProperty(propertyName, value, params, currentState, cursor.timeZone)
	case stateTodo:
		return parseTodoProperty(propertyName, value, params, cursor.todo)
	case stateJournal:
		return parseJournalProperty(propertyName, value, params, cursor.journal)
	case stateFreebusy:
		return parseFreeBusyProperty(propertyName, value, params, cursor.freeBusy)
	case stateStandard, stateDaylight:
		return parseTimeZonePropertySubComponent(propertyName, value, params, cursor.tzProp)
	case stateFinished:
		// Unreachable: parseOneCalendar returns as soon as stateFinished is reached; kept as a defensive default.
		return fmt.Errorf("%w: %s", icalerr.ErrPropertyWhenNotInCalendar, propertyName)
	case stateCalendar:
		return parseCalendarProperty(propertyName, value, params, calendar)
	}
	return nil
}

// handleBeginBlock processes BEGIN blocks and updates the parser state.
func handleBeginBlock(beginValue string, currentState *parserState, calendar *model.Calendar, cursor *componentCursor) error {
	switch model.SectionToken(beginValue) {
	case model.SectionTokenVEvent:
		*currentState = stateEvent
		calendar.Events = append(calendar.Events, model.Event{})
		cursor.event = &calendar.Events[len(calendar.Events)-1]
	// We have already verified that the first line is a BEGIN:VCALENDAR, so this is an error case
	case model.SectionTokenVCalendar:
		return icalerr.ErrNestedBeginVCalendar
	case model.SectionTokenVTimezone:
		*currentState = stateTimezone
		calendar.TimeZones = append(calendar.TimeZones, model.TimeZone{})
		cursor.timeZone = &calendar.TimeZones[len(calendar.TimeZones)-1]
	case model.SectionTokenVFreebusy:
		*currentState = stateFreebusy
		calendar.FreeBusys = append(calendar.FreeBusys, model.FreeBusy{})
		cursor.freeBusy = &calendar.FreeBusys[len(calendar.FreeBusys)-1]
	case model.SectionTokenVAlarm:
		switch *currentState { //nolint:exhaustive // it is an error condition to begin a VALARM block if we are not in an event or todo
		case stateEvent:
			if cursor.event == nil {
				return fmt.Errorf("%w: VALARM", icalerr.ErrUnexpectedBeginBlock)
			}
			cursor.event.Alarms = append(cursor.event.Alarms, model.Alarm{})
			cursor.alarm = &cursor.event.Alarms[len(cursor.event.Alarms)-1]
		case stateTodo:
			if cursor.todo == nil {
				return fmt.Errorf("%w: VALARM", icalerr.ErrUnexpectedBeginBlock)
			}
			cursor.todo.Alarms = append(cursor.todo.Alarms, model.Alarm{})
			cursor.alarm = &cursor.todo.Alarms[len(cursor.todo.Alarms)-1]
		case stateJournal:
			return fmt.Errorf("%w: VALARM not supported inside VJOURNAL", icalerr.ErrUnexpectedBeginBlock)
		default:
			return fmt.Errorf("%w: VALARM must be inside VEVENT or VTODO", icalerr.ErrUnexpectedBeginBlock)
		}
		cursor.alarmParent = *currentState
		*currentState = stateAlarm
	case model.SectionTokenVJournal:
		*currentState = stateJournal
		calendar.Journals = append(calendar.Journals, model.Journal{})
		cursor.journal = &calendar.Journals[len(calendar.Journals)-1]
	case model.SectionTokenVTodo:
		*currentState = stateTodo
		calendar.Todos = append(calendar.Todos, model.Todo{})
		cursor.todo = &calendar.Todos[len(calendar.Todos)-1]
	case model.SectionTokenVStandard:
		if *currentState != stateTimezone || cursor.timeZone == nil {
			return fmt.Errorf("%w: STANDARD must be inside VTIMEZONE", icalerr.ErrUnexpectedBeginBlock)
		}
		*currentState = stateStandard
		cursor.timeZone.Standard = append(cursor.timeZone.Standard, model.TimeZoneProperty{})
		cursor.tzProp = &cursor.timeZone.Standard[len(cursor.timeZone.Standard)-1]
	case model.SectionTokenVDaylight:
		if *currentState != stateTimezone || cursor.timeZone == nil {
			return fmt.Errorf("%w: DAYLIGHT must be inside VTIMEZONE", icalerr.ErrUnexpectedBeginBlock)
		}
		*currentState = stateDaylight
		cursor.timeZone.Daylight = append(cursor.timeZone.Daylight, model.TimeZoneProperty{})
		cursor.tzProp = &cursor.timeZone.Daylight[len(cursor.timeZone.Daylight)-1]
	default:
		return fmt.Errorf("%w: %s", icalerr.ErrTemplateInvalidStartBlock, beginValue)
	}
	return nil
}

// handleEndBlock processes END blocks and updates the parser state.
func handleEndBlock(endLineValue string, currentState *parserState, calendar *model.Calendar, cursor *componentCursor) error {
	switch endLineValue {
	case string(model.SectionTokenVEvent):
		if *currentState != stateEvent || cursor.event == nil {
			return fmt.Errorf("%w: END:VEVENT", icalerr.ErrUnexpectedEndBlock)
		}
		if err := validateEvent(cursor.event, calendar.Method); err != nil {
			return err
		}
		cursor.event = nil
		*currentState = stateCalendar
	case string(model.SectionTokenVCalendar):
		if *currentState != stateCalendar {
			return fmt.Errorf("%w: END:VCALENDAR", icalerr.ErrUnexpectedEndBlock)
		}
		if err := validateCalendar(calendar); err != nil {
			return err
		}
		*currentState = stateFinished
	case string(model.SectionTokenVTimezone):
		if *currentState != stateTimezone || cursor.timeZone == nil {
			return fmt.Errorf("%w: END:VTIMEZONE", icalerr.ErrUnexpectedEndBlock)
		}
		if err := validateTimeZone(cursor.timeZone); err != nil {
			return err
		}
		cursor.timeZone = nil
		*currentState = stateCalendar
	case string(model.SectionTokenVFreebusy):
		if *currentState != stateFreebusy || cursor.freeBusy == nil {
			return fmt.Errorf("%w: END:VFREEBUSY", icalerr.ErrUnexpectedEndBlock)
		}
		if err := validateFreeBusy(cursor.freeBusy); err != nil {
			return err
		}
		cursor.freeBusy = nil
		*currentState = stateCalendar
	case string(model.SectionTokenVAlarm):
		if *currentState != stateAlarm || cursor.alarm == nil {
			return fmt.Errorf("%w: END:VALARM", icalerr.ErrUnexpectedEndBlock)
		}
		if err := validateAlarm(cursor.alarm); err != nil {
			return err
		}
		cursor.alarm = nil
		*currentState = cursor.alarmParent
		cursor.alarmParent = 0
	case string(model.SectionTokenVJournal):
		if *currentState != stateJournal || cursor.journal == nil {
			return fmt.Errorf("%w: END:VJOURNAL", icalerr.ErrUnexpectedEndBlock)
		}
		if err := validateJournal(cursor.journal); err != nil {
			return err
		}
		cursor.journal = nil
		*currentState = stateCalendar
	case string(model.SectionTokenVTodo):
		if *currentState != stateTodo || cursor.todo == nil {
			return fmt.Errorf("%w: END:VTODO", icalerr.ErrUnexpectedEndBlock)
		}
		if err := validateTodo(cursor.todo); err != nil {
			return err
		}
		cursor.todo = nil
		*currentState = stateCalendar
	case string(model.SectionTokenVStandard):
		if *currentState != stateStandard || cursor.timeZone == nil {
			return fmt.Errorf("%w: END:STANDARD", icalerr.ErrUnexpectedEndBlock)
		}
		cursor.tzProp = nil
		*currentState = stateTimezone
	case string(model.SectionTokenVDaylight):
		if *currentState != stateDaylight || cursor.timeZone == nil {
			return fmt.Errorf("%w: END:DAYLIGHT", icalerr.ErrUnexpectedEndBlock)
		}
		cursor.tzProp = nil
		*currentState = stateTimezone
	default:
		return fmt.Errorf("%w: %s", icalerr.ErrTemplateInvalidEndBlock, endLineValue)
	}
	return nil
}
