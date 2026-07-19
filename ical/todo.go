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
	switch propertyName {
	case model.PropDTStamp:
		return setOnceTimePropertyWithParams(&todo.DTStamp, value, params, propertyName, todoLocation)
	case model.PropUID:
		return setOnceProperty(&todo.UID, value, propertyName, todoLocation)
	case model.PropClass:
		return setOnceProperty(&todo.Class, model.Class(value), propertyName, todoLocation)
	case model.PropCompleted:
		return setOnceTimePropertyWithParams(&todo.Completed, value, params, propertyName, todoLocation)
	case model.PropCreated:
		return setOnceTimePropertyWithParams(&todo.Created, value, params, propertyName, todoLocation)
	case model.PropDescription:
		return setOnceProperty(&todo.Description, value, propertyName, todoLocation)
	case model.PropDTStart:
		return setOnceTimePropertyWithParams(&todo.DTStart, value, params, propertyName, todoLocation)
	case model.PropDue:
		return setOnceTimePropertyWithParams(&todo.Due, value, params, propertyName, todoLocation)
	case model.PropDuration:
		if todo.Due != (time.Time{}) {
			return icalerr.ErrInvalidDurationPropertyDue
		}
		return setOnceDurationProperty(&todo.Duration, value, propertyName, todoLocation)

	case model.PropGeo:
		if todo.Geo != nil {
			return fmt.Errorf(icalerr.ErrDuplicatePropertyInComponentFormat, icalerr.ErrDuplicatePropertyInComponent, propertyName, todoLocation)
		}
		geo, err := parseGeo(value)
		if err != nil {
			return err
		}
		todo.Geo = &geo
	case model.PropLastModified:
		return setOnceTimePropertyWithParams(&todo.LastModified, value, params, propertyName, todoLocation)
	case model.PropLocation:
		return setOnceProperty(&todo.Location, value, propertyName, todoLocation)
	case model.PropOrganizer:
		organizer, err := parseOrganizer(value, params)
		if err != nil {
			return err
		}
		return setOnceProperty(&todo.Organizer, organizer, propertyName, todoLocation)
	case model.PropPercentComplete:
		return setOnceIntProperty(&todo.PercentComplete, value, propertyName, todoLocation)
	case model.PropPriority:
		return setOnceIntProperty(&todo.Priority, value, propertyName, todoLocation)
	case model.PropRecurrenceID:
		return setOnceTimePropertyWithParams(&todo.RecurrenceID, value, params, propertyName, todoLocation)
	case model.PropSequence:
		return setOnceIntProperty(&todo.Sequence, value, propertyName, todoLocation)
	case model.PropStatus:
		return setOnceProperty(&todo.Status, model.TodoStatus(value), propertyName, todoLocation)
	case model.PropSummary:
		return setOnceProperty(&todo.Summary, value, propertyName, todoLocation)
	case model.PropRRule:
		rule, err := rrule.ParseRRule(value)
		if err != nil {
			return fmt.Errorf("%w: %w", icalerr.ErrInvalidRRule, err)
		}
		return setOnceProperty(&todo.RRule, rule, propertyName, todoLocation)
	case model.PropTransp:
		return setOnceProperty(&todo.Transp, model.Transp(value), propertyName, todoLocation)
	case model.PropURL:
		return setOnceProperty(&todo.URL, value, propertyName, todoLocation)

	// Repeatable properties
	case model.PropAttach:
		attachment, err := parseAttachment(value, params)
		if err != nil {
			return err
		}
		todo.Attach = append(todo.Attach, *attachment)
		return nil
	case model.PropAttendee:
		attendee, err := parseAttendee(value, params)
		if err != nil {
			return err
		}
		todo.Attendees = append(todo.Attendees, *attendee)
	case model.PropCategories:
		todo.Categories = append(todo.Categories, strings.Split(value, ",")...)
	case model.PropComment:
		todo.Comment = append(todo.Comment, value)
	case model.PropContact:
		todo.Contacts = append(todo.Contacts, value)
	case model.PropExDate:
		return appendCommaSeparatedTimePropertyWithParams(&todo.ExceptionDates, value, params, propertyName, todoLocation)
	case model.PropRequestStatus:
		todo.RequestStatus = append(todo.RequestStatus, value)
	case model.PropRelatedTo:
		todo.Related = append(todo.Related, value)
	case model.PropResources:
		todo.Resources = append(todo.Resources, strings.Split(value, ",")...)
	case model.PropRDate:
		return appendCommaSeparatedTimePropertyWithParams(&todo.Rdate, value, params, propertyName, todoLocation)
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
