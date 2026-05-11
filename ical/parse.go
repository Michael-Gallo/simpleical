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

// Read takes an io.Reader containing iCalendar data and parses it into a Calendar.
func Read(reader io.Reader) (*model.Calendar, error) {
	calendar := &model.Calendar{}
	currentState := stateCalendar
	// Reusable parameter map to avoid allocations on every property
	reusableParams := make(map[string]string, 2)
	scanner := bufio.NewScanner(reader)

	if !scanner.Scan() {
		return nil, icalerr.ErrNoCalendarFound
	}

	line := strings.TrimRight(scanner.Text(), " ")
	if line != "BEGIN:VCALENDAR" {
		return nil, icalerr.ErrInvalidCalendarFormatMissingBegin
	}

	for scanner.Scan() {
		line := strings.TrimRight(scanner.Text(), " ")

		if line == "" {
			return nil, icalerr.ErrInvalidCalendarEmptyLine
		}

		// Clear the reusable parameter map before each use
		for k := range reusableParams {
			delete(reusableParams, k)
		}

		propertyName, params, value, err := parseIcalLineWithReusableMap(line, reusableParams)
		if err != nil {
			return nil, err
		}
		switch propertyName {
		case "BEGIN":
			if err := handleBeginBlock(value, &currentState, calendar); err != nil {
				return nil, err
			}
			continue
		case "END":
			if currentState == stateFinished {
				return nil, icalerr.ErrContentAfterEndBlock
			}
			if err := handleEndBlock(value, &currentState, calendar); err != nil {
				return nil, err
			}
			continue
		default:
			if currentState == stateFinished {
				return nil, icalerr.ErrContentAfterEndBlock
			}
			if err := parsePropertyLine(propertyName, value, params, currentState, calendar); err != nil {
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

// parsePropertyLine parses a single property line and adds it to the appropriate component based on current state.
func parsePropertyLine(propertyName string, value string, params map[string]string, currentState parserState, calendar *model.Calendar) error {
	// Route to appropriate parser based on current state
	switch currentState {
	case stateEventAlarm:
		currentAlarm := &calendar.Events[len(calendar.Events)-1].Alarms[len(calendar.Events[len(calendar.Events)-1].Alarms)-1]
		return parseAlarmProperty(propertyName, value, params, currentAlarm)
	case stateTodoAlarm:
		currentAlarm := &calendar.Todos[len(calendar.Todos)-1].Alarms[len(calendar.Todos[len(calendar.Todos)-1].Alarms)-1]
		return parseAlarmProperty(propertyName, value, params, currentAlarm)
	case stateEvent:
		return parseEventProperty(propertyName, value, params, &calendar.Events[len(calendar.Events)-1])
	case stateTimezone:
		return parseTimezoneProperty(propertyName, value, params, currentState, &calendar.TimeZones[len(calendar.TimeZones)-1])
	case stateTodo:
		return parseTodoProperty(propertyName, value, params, &calendar.Todos[len(calendar.Todos)-1])
	case stateJournal:
		return parseJournalProperty(propertyName, value, params, &calendar.Journals[len(calendar.Journals)-1])
	case stateFreebusy:
		return parseFreeBusyProperty(propertyName, value, params, &calendar.FreeBusys[len(calendar.FreeBusys)-1])
	case stateStandard, stateDaylight:
		// These are handled within timezone parsing
		return parseTimezoneProperty(propertyName, value, params, currentState, &calendar.TimeZones[len(calendar.TimeZones)-1])
	case stateFinished:
		return fmt.Errorf("%w: %s", icalerr.ErrPropertyWhenNotInCalendar, propertyName)
	case stateCalendar:
		return parseCalendarProperty(propertyName, value, params, calendar)
	}
	return nil
}

// handleBeginBlock processes BEGIN blocks and updates the parser state.
func handleBeginBlock(beginValue string, currentState *parserState, calendar *model.Calendar) error {
	switch beginValue {
	case string(model.SectionTokenVEvent):
		*currentState = stateEvent
		calendar.Events = append(calendar.Events, model.Event{})
	// We have already verified that the first line is a BEGIN:VCALENDAR, so this is an error case
	case string(model.SectionTokenVCalendar):
		return icalerr.ErrNestedBeginVCalendar
	case string(model.SectionTokenVTimezone):
		*currentState = stateTimezone
		calendar.TimeZones = append(calendar.TimeZones, model.TimeZone{})
	case string(model.SectionTokenVFreebusy):
		*currentState = stateFreebusy
		calendar.FreeBusys = append(calendar.FreeBusys, model.FreeBusy{})
	case string(model.SectionTokenVAlarm):
		// Determine which parent component to add the alarm to based on current state.
		switch *currentState {
		case stateEvent:
			if len(calendar.Events) == 0 {
				return fmt.Errorf("%w: VALARM", icalerr.ErrUnexpectedBeginBlock)
			}
			*currentState = stateEventAlarm
			calendar.Events[len(calendar.Events)-1].Alarms = append(calendar.Events[len(calendar.Events)-1].Alarms, model.Alarm{})
		case stateTodo:
			if len(calendar.Todos) == 0 {
				return fmt.Errorf("%w: VALARM", icalerr.ErrUnexpectedBeginBlock)
			}
			*currentState = stateTodoAlarm
			calendar.Todos[len(calendar.Todos)-1].Alarms = append(calendar.Todos[len(calendar.Todos)-1].Alarms, model.Alarm{})
		case stateJournal:
			return fmt.Errorf("%w: VALARM not supported inside VJOURNAL", icalerr.ErrUnexpectedBeginBlock)
		default:
			return fmt.Errorf("%w: VALARM must be inside VEVENT or VTODO", icalerr.ErrUnexpectedBeginBlock)
		}
	case string(model.SectionTokenVJournal):
		*currentState = stateJournal
		calendar.Journals = append(calendar.Journals, model.Journal{})
	case string(model.SectionTokenVTodo):
		*currentState = stateTodo
		calendar.Todos = append(calendar.Todos, model.Todo{})
	case string(model.SectionTokenVStandard):
		if *currentState != stateTimezone || len(calendar.TimeZones) == 0 {
			return fmt.Errorf("%w: STANDARD must be inside VTIMEZONE", icalerr.ErrUnexpectedBeginBlock)
		}
		*currentState = stateStandard
		calendar.TimeZones[len(calendar.TimeZones)-1].Standard = append(calendar.TimeZones[len(calendar.TimeZones)-1].Standard, model.TimeZoneProperty{})
	case string(model.SectionTokenVDaylight):
		if *currentState != stateTimezone || len(calendar.TimeZones) == 0 {
			return fmt.Errorf("%w: DAYLIGHT must be inside VTIMEZONE", icalerr.ErrUnexpectedBeginBlock)
		}
		*currentState = stateDaylight
		calendar.TimeZones[len(calendar.TimeZones)-1].Daylight = append(calendar.TimeZones[len(calendar.TimeZones)-1].Daylight, model.TimeZoneProperty{})
	default:
		return fmt.Errorf("%w: %s", icalerr.ErrTemplateInvalidStartBlock, beginValue)
	}
	return nil
}

// handleEndBlock processes END blocks and updates the parser state.
func handleEndBlock(endLineValue string, currentState *parserState, calendar *model.Calendar) error {
	switch endLineValue {
	case string(model.SectionTokenVEvent):
		if *currentState != stateEvent || len(calendar.Events) == 0 {
			return fmt.Errorf("%w: END:VEVENT", icalerr.ErrUnexpectedEndBlock)
		}
		if err := validateEvent(&calendar.Events[len(calendar.Events)-1]); err != nil {
			return err
		}
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
		if *currentState != stateTimezone || len(calendar.TimeZones) == 0 {
			return fmt.Errorf("%w: END:VTIMEZONE", icalerr.ErrUnexpectedEndBlock)
		}
		if err := validateTimeZone(&calendar.TimeZones[len(calendar.TimeZones)-1]); err != nil {
			return err
		}
		*currentState = stateCalendar
	case string(model.SectionTokenVFreebusy):
		if *currentState != stateFreebusy || len(calendar.FreeBusys) == 0 {
			return fmt.Errorf("%w: END:VFREEBUSY", icalerr.ErrUnexpectedEndBlock)
		}
		if err := validateFreeBusy(&calendar.FreeBusys[len(calendar.FreeBusys)-1]); err != nil {
			return err
		}
		*currentState = stateCalendar
	case string(model.SectionTokenVAlarm):
		// Validate alarm based on current state.
		switch *currentState { //nolint:exhaustive // it is an error condition to end a VALARM block if the state isn't an alarm state
		case stateEventAlarm:
			if len(calendar.Events) == 0 {
				return fmt.Errorf("%w: END:VALARM", icalerr.ErrUnexpectedEndBlock)
			}
			ev := &calendar.Events[len(calendar.Events)-1]
			if len(ev.Alarms) == 0 {
				return fmt.Errorf("%w: END:VALARM", icalerr.ErrUnexpectedEndBlock)
			}
			if err := validateAlarm(&ev.Alarms[len(ev.Alarms)-1]); err != nil {
				return err
			}
			*currentState = stateEvent // Return to parent state.
		case stateTodoAlarm:
			if len(calendar.Todos) == 0 {
				return fmt.Errorf("%w: END:VALARM", icalerr.ErrUnexpectedEndBlock)
			}
			todo := &calendar.Todos[len(calendar.Todos)-1]
			if len(todo.Alarms) == 0 {
				return fmt.Errorf("%w: END:VALARM", icalerr.ErrUnexpectedEndBlock)
			}
			if err := validateAlarm(&todo.Alarms[len(todo.Alarms)-1]); err != nil {
				return err
			}
			*currentState = stateTodo // Return to parent state.
		default:
			return fmt.Errorf("%w: END:VALARM", icalerr.ErrUnexpectedEndBlock)
		}
	case string(model.SectionTokenVJournal):
		if *currentState != stateJournal || len(calendar.Journals) == 0 {
			return fmt.Errorf("%w: END:VJOURNAL", icalerr.ErrUnexpectedEndBlock)
		}
		if err := validateJournal(&calendar.Journals[len(calendar.Journals)-1]); err != nil {
			return err
		}
		*currentState = stateCalendar
	case string(model.SectionTokenVTodo):
		if *currentState != stateTodo || len(calendar.Todos) == 0 {
			return fmt.Errorf("%w: END:VTODO", icalerr.ErrUnexpectedEndBlock)
		}
		if err := validateTodo(&calendar.Todos[len(calendar.Todos)-1]); err != nil {
			return err
		}
		*currentState = stateCalendar
	case string(model.SectionTokenVStandard):
		if *currentState != stateStandard || len(calendar.TimeZones) == 0 {
			return fmt.Errorf("%w: END:STANDARD", icalerr.ErrUnexpectedEndBlock)
		}
		*currentState = stateTimezone
	case string(model.SectionTokenVDaylight):
		if *currentState != stateDaylight || len(calendar.TimeZones) == 0 {
			return fmt.Errorf("%w: END:DAYLIGHT", icalerr.ErrUnexpectedEndBlock)
		}
		*currentState = stateTimezone
	default:
		return fmt.Errorf("%w: %s", icalerr.ErrTemplateInvalidEndBlock, endLineValue)
	}
	return nil
}
