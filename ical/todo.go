package ical

import (
	"fmt"
	"time"

	"github.com/michael-gallo/simpleical/internal/icalerr"
	"github.com/michael-gallo/simpleical/model"
	"github.com/michael-gallo/simpleical/rrule"
)

const todoLocation = "Todo"

// parseTodoProperty parses a single iCalendar TODO property line and applies it to the provided Todo.
func parseTodoProperty(propertyName string, value string, params map[string]string, todo *model.Todo) error {
	switch propertyName {
	case model.PropDTStamp:
		return setOnceUTCTimePropertyWithParams(&todo.DTStamp, value, params, propertyName, todoLocation)
	case model.PropUID:
		return setOnceProperty(&todo.UID, value, propertyName, todoLocation)
	case model.PropClass:
		class, err := parseClass(value)
		if err != nil {
			return err
		}
		return setOnceProperty(&todo.Class, class, propertyName, todoLocation)
	case model.PropCompleted:
		return setOnceUTCTimePropertyWithParams(&todo.Completed, value, params, propertyName, todoLocation)
	case model.PropCreated:
		return setOnceUTCTimePropertyWithParams(&todo.Created, value, params, propertyName, todoLocation)
	case model.PropDescription:
		return setOnceTextProperty(&todo.Description, value, params, propertyName, todoLocation)
	case model.PropDTStart:
		return setOnceTimePropertyWithParams(&todo.DTStart, value, params, propertyName, todoLocation)
	case model.PropDue:
		if todo.Duration != 0 {
			return icalerr.ErrInvalidDurationPropertyDue
		}
		return setOnceTimePropertyWithParams(&todo.Due, value, params, propertyName, todoLocation)
	case model.PropDuration:
		if !todo.Due.IsZero() {
			return icalerr.ErrInvalidDurationPropertyDue
		}
		return setOncePositiveDurationProperty(&todo.Duration, value, propertyName, todoLocation)

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
		return setOnceUTCTimePropertyWithParams(&todo.LastModified, value, params, propertyName, todoLocation)
	case model.PropLocation:
		return setOnceTextProperty(&todo.Location, value, params, propertyName, todoLocation)
	case model.PropOrganizer:
		organizer, err := parseOrganizer(value, params)
		if err != nil {
			return err
		}
		return setOnceProperty(&todo.Organizer, organizer, propertyName, todoLocation)
	case model.PropPercentComplete:
		return setOnceBoundedInt(&todo.PercentComplete, value, 0, 100, icalerr.ErrInvalidPercentComplete, propertyName, todoLocation)
	case model.PropPriority:
		return setOnceBoundedInt(&todo.Priority, value, 0, 9, icalerr.ErrInvalidPriority, propertyName, todoLocation)
	case model.PropRecurrenceID:
		if err := setOnceTimePropertyWithParams(&todo.RecurrenceID, value, params, propertyName, todoLocation); err != nil {
			return err
		}
		rangeParam, err := parseRecurrenceIDRange(params)
		if err != nil {
			return err
		}
		todo.RecurrenceIDRange = rangeParam
		return nil
	case model.PropSequence:
		return setOnceIntProperty(&todo.Sequence, value, propertyName, todoLocation)
	case model.PropStatus:
		status, err := parseTodoStatus(value)
		if err != nil {
			return err
		}
		return setOnceProperty(&todo.Status, status, propertyName, todoLocation)
	case model.PropSummary:
		return setOnceTextProperty(&todo.Summary, value, params, propertyName, todoLocation)
	case model.PropRRule:
		rule, err := rrule.ParseRRule(value)
		if err != nil {
			return fmt.Errorf("%w: %w", icalerr.ErrInvalidRRule, err)
		}
		return setOnceProperty(&todo.RRule, rule, propertyName, todoLocation)
	case model.PropTransp:
		return icalerr.ErrTodoTranspNotAllowed
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
		return appendTextListProperty(&todo.Categories, value)
	case model.PropComment:
		return appendTextProperty(&todo.Comment, value, params)
	case model.PropContact:
		todo.Contacts = append(todo.Contacts, value)
	case model.PropExDate:
		return appendCommaSeparatedTimePropertyWithParams(&todo.ExceptionDates, value, params, propertyName, todoLocation)
	case model.PropRequestStatus:
		todo.RequestStatus = append(todo.RequestStatus, value)
	case model.PropRelatedTo:
		return appendRelatedToProperty(&todo.Related, value, params)
	case model.PropResources:
		return appendTextListProperty(&todo.Resources, value)
	case model.PropRDate:
		return appendRDateProperty(&todo.Rdate, value, params, propertyName, todoLocation)
	default:
		appendExtensionProperty(&todo.XProp, &todo.IANAProp, propertyName, value, params)
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
		if todo.DTStart.IsDate() && todo.Duration%(24*time.Hour) != 0 {
			return icalerr.ErrDateDurationMustBeDayOrWeek
		}
	}

	if !todo.DTStart.IsZero() && !todo.Due.IsZero() && todo.DTStart.IsDate() != todo.Due.IsDate() {
		return icalerr.ErrMismatchedDateValueTypes
	}

	return nil
}
