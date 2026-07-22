package ical

import (
	"fmt"
	"net/url"
	"time"

	"github.com/michael-gallo/simpleical/internal/icalerr"
	"github.com/michael-gallo/simpleical/model"
	"github.com/michael-gallo/simpleical/rrule"
)

const eventLocation = "Event"

// parseEventProperty parses a single property line and adds it to the provided vevent.
func parseEventProperty(propertyName string, value string, params map[string]string, event *model.Event) error {
	switch propertyName {
	case model.PropDTStart:
		return setOnceTimePropertyWithParams(&event.Start, value, params, propertyName, eventLocation)
	case model.PropDTStamp:
		return setOnceUTCTimePropertyWithParams(&event.DTStamp, value, params, propertyName, eventLocation)

	// End and Duration are mutually exclusive
	case model.PropDTEnd:
		if event.Duration != 0 {
			return icalerr.ErrInvalidDurationPropertyDtend
		}
		return setOnceTimePropertyWithParams(&event.End, value, params, propertyName, eventLocation)
	case model.PropDuration:
		if !event.End.IsZero() {
			return icalerr.ErrInvalidDurationPropertyDtend
		}
		return setOncePositiveDurationProperty(&event.Duration, value, propertyName, eventLocation)
	case model.PropLastModified:
		return setOnceUTCTimePropertyWithParams(&event.LastModified, value, params, propertyName, eventLocation)

	case model.PropSummary:
		return setOnceTextProperty(&event.Summary, value, params, propertyName, eventLocation)
	case model.PropDescription:
		return setOnceTextProperty(&event.Description, value, params, propertyName, eventLocation)
	case model.PropLocation:
		return setOnceTextProperty(&event.Location, value, params, propertyName, eventLocation)
	case model.PropUID:
		return setOnceProperty(&event.UID, value, propertyName, eventLocation)
	case model.PropClass:
		class, err := parseClass(value)
		if err != nil {
			return err
		}
		return setOnceProperty(&event.Class, class, propertyName, eventLocation)
	case model.PropCreated:
		return setOnceUTCTimePropertyWithParams(&event.Created, value, params, propertyName, eventLocation)
	case model.PropPriority:
		return setOncePriorityProperty(&event.Priority, value, propertyName, eventLocation)
	case model.PropURL:
		return setOnceProperty(&event.URL, value, propertyName, eventLocation)
	case model.PropRecurrenceID:
		return setOnceTimePropertyWithParams(&event.RecurrenceID, value, params, propertyName, eventLocation)
	case model.PropContact:
		event.Contacts = append(event.Contacts, value)
		return nil

	case model.PropStatus:
		status, err := parseEventStatus(value)
		if err != nil {
			return err
		}
		return setOnceProperty(&event.Status, status, propertyName, eventLocation)
	case model.PropTransp:
		transp, err := parseTransp(value)
		if err != nil {
			return err
		}
		return setOnceProperty(&event.Transp, transp, propertyName, eventLocation)
	case model.PropSequence:
		return setOnceIntProperty(&event.Sequence, value, propertyName, eventLocation)
	case model.PropOrganizer:
		organizer, err := parseOrganizer(value, params)
		if err != nil {
			return err
		}
		return setOnceProperty(&event.Organizer, organizer, propertyName, eventLocation)
	case model.PropComment:
		return appendTextProperty(&event.Comment, value, params)
	case model.PropCategories:
		return appendTextListProperty(&event.Categories, value)
	case model.PropGeo:
		if event.Geo != nil {
			return fmt.Errorf(icalerr.ErrDuplicatePropertyInComponentFormat, icalerr.ErrDuplicatePropertyInComponent, propertyName, eventLocation)
		}
		geo, err := parseGeo(value)
		if err != nil {
			return err
		}
		event.Geo = &geo
	case model.PropRRule:
		rule, err := rrule.ParseRRule(value)
		if err != nil {
			return fmt.Errorf("%w: %w", icalerr.ErrInvalidRRule, err)
		}
		return setOnceProperty(&event.RRule, rule, propertyName, eventLocation)
	case model.PropAttach:
		attachment, err := parseAttachment(value, params)
		if err != nil {
			return err
		}
		event.Attach = append(event.Attach, *attachment)
		return nil
	case model.PropAttendee:
		attendee, err := parseAttendee(value, params)
		if err != nil {
			return err
		}
		event.Attendees = append(event.Attendees, *attendee)
		return nil
	case model.PropExDate:
		return appendCommaSeparatedTimePropertyWithParams(&event.ExceptionDates, value, params, propertyName, eventLocation)
	case model.PropRequestStatus:
		event.RequestStatus = append(event.RequestStatus, value)
	case model.PropRelatedTo:
		return appendRelatedToProperty(&event.Related, value, params)
	case model.PropResources:
		return appendTextListProperty(&event.Resources, value)
	case model.PropRDate:
		return appendRDateProperty(&event.Rdate, value, params, propertyName, eventLocation)
	default:
		appendExtensionProperty(&event.XProp, &event.IANAProp, propertyName, value, params)
	}
	return nil
}

// parseOrganizer parses a calendar line starting with ORGANIZER.
func parseOrganizer(value string, params map[string]string) (*model.Organizer, error) {
	organizer := &model.Organizer{}
	for propName, propValue := range params {
		switch propName {
		case model.ParamCN:
			organizer.CommonName = propValue
		case model.ParamDir:
			parsedURI, err := url.Parse(propValue)
			if err != nil {
				return nil, fmt.Errorf("%w: %w", icalerr.ErrInvalidOrganizer, err)
			}
			organizer.Directory = parsedURI
		case model.ParamLanguage:
			organizer.Language = propValue
		case model.ParamSentBy:
			parsedURI, err := url.Parse(propValue)
			if err != nil {
				return nil, fmt.Errorf("%w: %w", icalerr.ErrInvalidOrganizer, err)
			}
			organizer.SentBy = parsedURI
		default:
			if organizer.OtherParams == nil {
				organizer.OtherParams = make(map[string]string)
			}
			organizer.OtherParams[propName] = propValue
		}
	}

	parsedURI, err := url.Parse(value)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", icalerr.ErrInvalidOrganizer, err)
	}
	organizer.CalAddress = parsedURI

	return organizer, nil
}

// validateEvent ensures that all required values are present for an event.
// DTSTART is required only when the enclosing calendar has no METHOD property.
func validateEvent(event *model.Event, method string) error {
	if event.UID == "" {
		return icalerr.ErrMissingEventUIDProperty
	}
	if event.DTStamp.IsZero() {
		return icalerr.ErrMissingEventDTStampProperty
	}
	if method == "" && event.Start.IsZero() {
		return icalerr.ErrMissingEventDTStartProperty
	}
	if !event.Start.IsZero() && !event.End.IsZero() && event.Start.IsDate() != event.End.IsDate() {
		return icalerr.ErrMismatchedDateValueTypes
	}
	if event.Duration != 0 && event.Start.IsDate() && event.Duration%(24*time.Hour) != 0 {
		return icalerr.ErrDateDurationMustBeDayOrWeek
	}
	return nil
}
