package ical

import (
	"fmt"
	"strings"
	"time"

	"github.com/michael-gallo/simpleical/internal/icalerr"
	"github.com/michael-gallo/simpleical/model"
	"github.com/michael-gallo/simpleical/rrule"
)

const todoLocation = "Todo"

// parseTodoProperty parses a single iCalendar TODO property line and applies it to the provided Todo.
// It updates the appropriate field(s) on todo (including appending repeatable properties), performs type
// conversions, and enforces single-assignment rules. It also validates the Geo latitude;longitude format.
// An error is returned for invalid property names, duplicate
// assignments, or any parse failures.
func parseTodoProperty(propertyName string, value string, params map[string]string, todo *model.Todo) error {
	switch model.TodoToken(propertyName) {
	case model.TodoTokenDTStamp:
		return setOnceTimeProperty(&todo.DTStamp, value, propertyName, todoLocation)
	case model.TodoTokenUID:
		return setOnceProperty(&todo.UID, value, propertyName, todoLocation)
	case model.TodoTokenClass:
		return setOnceProperty(&todo.Class, model.TodoClass(value), propertyName, todoLocation)
	case model.TodoTokenCompleted:
		return setOnceTimeProperty(&todo.Completed, value, propertyName, todoLocation)
	case model.TodoTokenCreated:
		return setOnceTimeProperty(&todo.Created, value, propertyName, todoLocation)
	case model.TodoTokenDescription:
		return setOnceProperty(&todo.Description, value, propertyName, todoLocation)
	case model.TodoTokenDTStart:
		return setOnceTimeProperty(&todo.DTStart, value, propertyName, todoLocation)
	case model.TodoTokenDue:
		return setOnceTimeProperty(&todo.Due, value, propertyName, todoLocation)
	case model.TodoTokenDuration:
		if todo.Due != (time.Time{}) {
			return icalerr.ErrInvalidDurationPropertyDue
		}
		return setOnceDurationProperty(&todo.Duration, value, propertyName, todoLocation)

	case model.TodoTokenGeo:
		if todo.Geo != nil {
			return fmt.Errorf(icalerr.ErrDuplicatePropertyInComponentFormat, icalerr.ErrDuplicatePropertyInComponent, propertyName, todoLocation)
		}
		geo, err := parseGeo(value)
		if err != nil {
			return err
		}
		todo.Geo = geo
	case model.TodoTokenLastModified:
		return setOnceTimeProperty(&todo.LastModified, value, propertyName, todoLocation)
	case model.TodoTokenLocation:
		return setOnceProperty(&todo.Location, value, propertyName, todoLocation)
	case model.TodoTokenOrganizer:
		organizer, err := parseOrganizer(value, params)
		if err != nil {
			return err
		}
		return setOnceProperty(&todo.Organizer, organizer, propertyName, todoLocation)
	case model.TodoTokenPercentComplete:
		return setOnceIntProperty(&todo.PercentComplete, value, propertyName, todoLocation)
	case model.TodoTokenPriority:
		return setOnceIntProperty(&todo.Priority, value, propertyName, todoLocation)
	case model.TodoTokenRecurrenceID:
		return setOnceTimeProperty(&todo.RecurrenceID, value, propertyName, todoLocation)
	case model.TodoTokenSequence:
		return setOnceIntProperty(&todo.Sequence, value, propertyName, todoLocation)
	case model.TodoTokenStatus:
		return setOnceProperty(&todo.Status, model.TodoStatus(value), propertyName, todoLocation)
	case model.TodoTokenSummary:
		return setOnceProperty(&todo.Summary, value, propertyName, todoLocation)
	case model.TodoTokenRRule:
		rule, err := rrule.ParseRRule(value)
		if err != nil {
			return fmt.Errorf("%w: %w", icalerr.ErrInvalidRRule, err)
		}
		return setOnceProperty(&todo.RRule, rule, propertyName, todoLocation)
	case model.TodoTokenTransp:
		return setOnceProperty(&todo.Transp, model.TodoTransp(value), propertyName, todoLocation)
	case model.TodoTokenURL:
		return setOnceProperty(&todo.URL, value, propertyName, todoLocation)

	// Repeatable properties
	case model.TodoTokenAttach:
		attachment, err := parseAttachment(value, params)
		if err != nil {
			return err
		}
		todo.Attach = append(todo.Attach, *attachment)
		return nil
	case model.TodoTokenAttendee:
		attendee, err := parseAttendee(value, params)
		if err != nil {
			return err
		}
		todo.Attendees = append(todo.Attendees, *attendee)
	case model.TodoTokenCategories:
		todo.Categories = append(todo.Categories, strings.Split(value, ",")...)
	case model.TodoTokenComment:
		todo.Comment = append(todo.Comment, value)
	case model.TodoTokenContact:
		todo.Contacts = append(todo.Contacts, value)
	case model.TodoTokenExceptionDates:
		return appendTimeProperty(&todo.ExceptionDates, value, propertyName, todoLocation)
	case model.TodoTokenRequestStatus:
		todo.RequestStatus = append(todo.RequestStatus, value)
	case model.TodoTokenRelated:
		todo.Related = append(todo.Related, value)
	case model.TodoTokenResources:
		todo.Resources = append(todo.Resources, strings.Split(value, ",")...)
	case model.TodoTokenRdate:
		return appendTimeProperty(&todo.Rdate, value, propertyName, todoLocation)
	default:
		return fmt.Errorf("%w: %s", icalerr.ErrInvalidTodoProperty, propertyName)
	}
	return nil
}

// validateTodo ensures that all required values are present for a todo.
// Per RFC 5545: If 'duration' appears in a 'todoprop', then 'dtstart' MUST also appear,
// and 'due' and 'duration' MUST NOT occur in the same 'todoprop'.
// DTstamp and uid are always required.
func validateTodo(todo *model.Todo) error {
	if todo.UID == "" {
		return icalerr.ErrMissingTodoUIDProperty
	}
	if todo.DTStamp.IsZero() {
		return icalerr.ErrMissingTodoDTStampProperty
	}

	if todo.Duration != 0 {
		if todo.DTStart.IsZero() {
			return icalerr.ErrDurationRequiresDTStart
		}
		if !todo.Due.IsZero() {
			return icalerr.ErrInvalidDurationPropertyDue
		}
	}

	return nil
}
