// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package ical

import (
	"bufio"
	"fmt"
	"io"
	"os"
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
	stateEventAlarm
	stateTodoAlarm
	stateStandard
	stateDaylight
	stateFinished
)

// FromFileName parses an iCalendar file from the given file path into a Calendar.
// It opens the file, parses its contents, and returns a Calendar.
// This is a convenience function that wraps Read.
// The file is automatically closed after parsing.
func FromFileName(filename string) (*model.Calendar, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return Read(file)
}

// FromString takes the string representation of an ICAL and parses it into a Calendar.
// It returns an error if the input is not a valid ICAL string.
// This is a convenience function that wraps Read.
func FromString(input string) (*model.Calendar, error) {
	// Handle empty input
	if input == "" {
		return nil, icalerr.ErrNoCalendarFound
	}

	// Use the reader-based parser for consistency
	reader := strings.NewReader(input)
	return Read(reader)
}

// componentCursor caches pointers to the active component so property handlers
// avoid repeated &slice[len-1] indexing on every line.
type componentCursor struct {
	event    *model.Event
	todo     *model.Todo
	journal  *model.Journal
	freeBusy *model.FreeBusy
	timeZone *model.TimeZone
	alarm    *model.Alarm
	tzProp   *model.TimeZoneProperty
}

// Read takes an io.Reader containing iCalendar data and parses it into a Calendar.
func Read(reader io.Reader) (*model.Calendar, error) {
	calendar := &model.Calendar{}
	currentState := stateCalendar
	var cursor componentCursor
	// Reusable parameter map to avoid allocations on every property
	reusableParams := make(map[string]string, 2)
	scanner := bufio.NewScanner(reader)
	var pending string
	var hasPending bool

	line, ok := nextLogicalLine(scanner, &pending, &hasPending)
	if !ok {
		if err := scanner.Err(); err != nil {
			return nil, fmt.Errorf("error reading iCalendar data: %w", err)
		}
		return nil, icalerr.ErrNoCalendarFound
	}
	line = strings.TrimRight(line, " ")
	if line != "BEGIN:VCALENDAR" {
		return nil, icalerr.ErrInvalidCalendarFormatMissingBegin
	}

	for {
		line, ok = nextLogicalLine(scanner, &pending, &hasPending)
		if !ok {
			break
		}
		line = strings.TrimRight(line, " ")

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
			continue
		case "END":
			if currentState == stateFinished {
				return nil, icalerr.ErrContentAfterEndBlock
			}
			if err := handleEndBlock(value, &currentState, calendar, &cursor); err != nil {
				return nil, err
			}
			continue
		default:
			if currentState == stateFinished {
				return nil, icalerr.ErrContentAfterEndBlock
			}
			if err := parsePropertyLine(propertyName, value, params, currentState, calendar, &cursor); err != nil {
				return nil, err
			}
			continue
		}
	}

	// Check for scanner errors
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading iCalendar data: %w", err)
	}

	// Verify that the last line was a END:VCALENDAR
	if currentState != stateFinished {
		return nil, icalerr.ErrInvalidCalendarFormatMissingEnd
	}

	return calendar, nil
}

// nextLogicalLine returns the next RFC 5545 content line after unfolding.
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
		return b.String(), true
	}
	return line, true
}

// parsePropertyLine parses a single property line and adds it to the appropriate component based on current state.
func parsePropertyLine(propertyName string, value string, params map[string]string, currentState parserState, calendar *model.Calendar, cursor *componentCursor) error {
	switch currentState {
	case stateEventAlarm:
		return parseAlarmProperty(propertyName, value, params, cursor.alarm)
	case stateTodoAlarm:
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
		// Unreachable from Read (stateFinished is guarded earlier); kept as a defensive default.
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
			*currentState = stateEventAlarm
			cursor.event.Alarms = append(cursor.event.Alarms, model.Alarm{})
			cursor.alarm = &cursor.event.Alarms[len(cursor.event.Alarms)-1]
		case stateTodo:
			if cursor.todo == nil {
				return fmt.Errorf("%w: VALARM", icalerr.ErrUnexpectedBeginBlock)
			}
			*currentState = stateTodoAlarm
			cursor.todo.Alarms = append(cursor.todo.Alarms, model.Alarm{})
			cursor.alarm = &cursor.todo.Alarms[len(cursor.todo.Alarms)-1]
		case stateJournal:
			return fmt.Errorf("%w: VALARM not supported inside VJOURNAL", icalerr.ErrUnexpectedBeginBlock)
		default:
			return fmt.Errorf("%w: VALARM must be inside VEVENT or VTODO", icalerr.ErrUnexpectedBeginBlock)
		}
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
		if err := validateEvent(cursor.event); err != nil {
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
		switch *currentState { //nolint:exhaustive // it is an error condition to end a VALARM block if the state isn't an alarm state
		case stateEventAlarm:
			if cursor.alarm == nil {
				return fmt.Errorf("%w: END:VALARM", icalerr.ErrUnexpectedEndBlock)
			}
			if err := validateAlarm(cursor.alarm); err != nil {
				return err
			}
			cursor.alarm = nil
			*currentState = stateEvent
		case stateTodoAlarm:
			if cursor.alarm == nil {
				return fmt.Errorf("%w: END:VALARM", icalerr.ErrUnexpectedEndBlock)
			}
			if err := validateAlarm(cursor.alarm); err != nil {
				return err
			}
			cursor.alarm = nil
			*currentState = stateTodo
		default:
			return fmt.Errorf("%w: END:VALARM", icalerr.ErrUnexpectedEndBlock)
		}
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
