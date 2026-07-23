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
	stateSkipComponent
	stateFinished
)

// beginVCalendarLine is the content line that opens a VCALENDAR object.
const beginVCalendarLine = "BEGIN:VCALENDAR"

// maxPhysicalLineBytes is the maximum size of a single physical line token.
// RFC 5545 folds at 75 octets, but unfolded logical lines may be longer.
const maxPhysicalLineBytes = 4 * 1024 * 1024

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

// skipTracker tracks nested unknown x-comp / iana-comp blocks so they can be ignored.
type skipTracker struct {
	depth       int
	returnState parserState
}

// calendarParseState holds per-calendar parse bookkeeping.
type calendarParseState struct {
	sawComponent bool
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
	reusableParams := make(map[string]string, 2)
	scanner := bufio.NewScanner(reader)
	scanner.Buffer(make([]byte, 64*1024), maxPhysicalLineBytes)
	var pending string
	var hasPending bool
	var calendars []*model.Calendar

	for {
		line, ok := nextLogicalLine(scanner, &pending, &hasPending)
		if !ok {
			break
		}
		if !strings.EqualFold(line, beginVCalendarLine) {
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
func parseOneCalendar(scanner *bufio.Scanner, pending *string, hasPending *bool, reusableParams map[string]string) (*model.Calendar, error) {
	calendar := &model.Calendar{}
	currentState := stateCalendar
	var cursor componentCursor
	var skip skipTracker
	var cps calendarParseState

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
		propertyName = strings.ToUpper(propertyName)
		switch propertyName {
		case "BEGIN":
			value = strings.ToUpper(value)
			if currentState == stateSkipComponent {
				skip.depth++
				continue
			}
			if err := handleBeginBlock(value, &currentState, calendar, &cursor, &skip, &cps); err != nil {
				return nil, err
			}
		case "END":
			value = strings.ToUpper(value)
			if currentState == stateSkipComponent {
				skip.depth--
				if skip.depth == 0 {
					currentState = skip.returnState
				}
				continue
			}
			if err := handleEndBlock(value, &currentState, calendar, &cursor); err != nil {
				return nil, err
			}
			if currentState == stateFinished {
				if err := validateCalendar(calendar, &cps); err != nil {
					return nil, err
				}
				if err := validateCalendarTZIDs(calendar); err != nil {
					return nil, err
				}
				return calendar, nil
			}
		default:
			if currentState == stateSkipComponent {
				continue
			}
			if err := parsePropertyLine(propertyName, value, params, currentState, calendar, &cursor); err != nil {
				return nil, err
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading iCalendar data: %w", err)
	}

	return nil, icalerr.ErrInvalidCalendarFormatMissingEnd
}

// nextLogicalLine returns the next RFC 5545 content line after unfolding.
// Trailing value whitespace is preserved.
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
		return b.String(), true
	}
	return line, true
}

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
	case stateSkipComponent:
		return nil
	case stateFinished:
		return fmt.Errorf("%w: %s", icalerr.ErrPropertyWhenNotInCalendar, propertyName)
	case stateCalendar:
		return parseCalendarProperty(propertyName, value, params, calendar)
	}
	return nil
}

func handleBeginBlock(beginValue string, currentState *parserState, calendar *model.Calendar, cursor *componentCursor, skip *skipTracker, cps *calendarParseState) error {
	token := model.SectionToken(beginValue)

	// Top-level calendar components may only begin from calendar state.
	switch token {
	case model.SectionTokenVEvent, model.SectionTokenVTodo, model.SectionTokenVJournal,
		model.SectionTokenVFreebusy, model.SectionTokenVTimezone:
		if *currentState != stateCalendar {
			return fmt.Errorf("%w: %s", icalerr.ErrComponentNotAllowedHere, beginValue)
		}
		cps.sawComponent = true
	case model.SectionTokenVCalendar:
		return icalerr.ErrNestedBeginVCalendar
	case model.SectionTokenVAlarm, model.SectionTokenVStandard, model.SectionTokenVDaylight:
		// Nested components; parent checks happen in the dispatch switch below.
	}

	switch token {
	case model.SectionTokenVEvent:
		*currentState = stateEvent
		calendar.Events = append(calendar.Events, model.Event{})
		cursor.event = &calendar.Events[len(calendar.Events)-1]
	case model.SectionTokenVCalendar:
		// Unreachable: rejected above. Kept for exhaustive switch coverage.
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
		switch *currentState { //nolint:exhaustive // VALARM is only valid inside VEVENT/VTODO; other states share the default error
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
		// RFC 5545: applications MUST ignore unrecognized x-comp / iana-comp.
		cps.sawComponent = true
		skip.returnState = *currentState
		skip.depth = 1
		*currentState = stateSkipComponent
	}
	return nil
}

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
		if *currentState != stateStandard || cursor.timeZone == nil || cursor.tzProp == nil {
			return fmt.Errorf("%w: END:STANDARD", icalerr.ErrUnexpectedEndBlock)
		}
		if err := validateObservance(cursor.tzProp); err != nil {
			return err
		}
		cursor.tzProp = nil
		*currentState = stateTimezone
	case string(model.SectionTokenVDaylight):
		if *currentState != stateDaylight || cursor.timeZone == nil || cursor.tzProp == nil {
			return fmt.Errorf("%w: END:DAYLIGHT", icalerr.ErrUnexpectedEndBlock)
		}
		if err := validateObservance(cursor.tzProp); err != nil {
			return err
		}
		cursor.tzProp = nil
		*currentState = stateTimezone
	default:
		return fmt.Errorf("%w: %s", icalerr.ErrTemplateInvalidEndBlock, endLineValue)
	}
	return nil
}
