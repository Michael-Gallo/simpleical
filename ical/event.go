package ical

import (
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/michael-gallo/simpleical/internal/icalerr"
	"github.com/michael-gallo/simpleical/model"
	"github.com/michael-gallo/simpleical/rrule"
)

const eventLocation = "Event"

// parseEventProperty parses a single property line and adds it to the provided vevent.
func parseEventProperty(propertyName string, value string, params map[string]string, event *model.Event) error {
	switch model.EventToken(propertyName) {
	case model.EventTokenDtstart:
		return setOnceTimePropertyWithParams(&event.Start, value, params, propertyName, eventLocation)
	case model.EventTokenDTStamp:
		return setOnceTimePropertyWithParams(&event.DTStamp, value, params, propertyName, eventLocation)

	// End and Duration are mutually exclusive
	case model.EventTokenDtend:
		if event.Duration != 0 {
			return icalerr.ErrInvalidDurationPropertyDtend
		}
		return setOnceTimePropertyWithParams(&event.End, value, params, propertyName, eventLocation)
	case model.EventTokenDuration:
		if event.End != (time.Time{}) {
			return icalerr.ErrInvalidDurationPropertyDtend
		}
		return setOnceDurationProperty(&event.Duration, value, propertyName, eventLocation)
	case model.EventTokenLastModified:
		return setOnceTimePropertyWithParams(&event.LastModified, value, params, propertyName, eventLocation)

	case model.EventTokenSummary:
		return setOnceProperty(&event.Summary, value, propertyName, eventLocation)
	case model.EventTokenDescription:
		return setOnceProperty(&event.Description, value, propertyName, eventLocation)
	case model.EventTokenLocation:
		return setOnceProperty(&event.Location, value, propertyName, eventLocation)
	case model.EventTokenUID:
		return setOnceProperty(&event.UID, value, propertyName, eventLocation)
	case model.EventTokenClass:
		return setOnceProperty(&event.Class, model.EventClass(value), propertyName, eventLocation)
	case model.EventTokenCreated:
		return setOnceTimePropertyWithParams(&event.Created, value, params, propertyName, eventLocation)
	case model.EventTokenPriority:
		return setOnceIntProperty(&event.Priority, value, propertyName, eventLocation)
	case model.EventTokenURL:
		return setOnceProperty(&event.URL, value, propertyName, eventLocation)
	case model.EventTokenRecurrenceID:
		return setOnceTimePropertyWithParams(&event.RecurrenceID, value, params, propertyName, eventLocation)
	case model.EventTokenContact:
		event.Contacts = append(event.Contacts, value)
		return nil

	case model.EventTokenStatus:
		return setOnceProperty(&event.Status, model.EventStatus(value), propertyName, eventLocation)
	case model.EventTokenTransp:
		return setOnceProperty(&event.Transp, model.EventTransp(value), propertyName, eventLocation)
	case model.EventTokenSequence:
		return setOnceIntProperty(&event.Sequence, value, propertyName, eventLocation)
	case model.EventTokenOrganizer:
		organizer, err := parseOrganizer(value, params)
		if err != nil {
			return err
		}
		return setOnceProperty(&event.Organizer, organizer, propertyName, eventLocation)
	case model.EventTokenComment:
		event.Comment = append(event.Comment, value)
	case model.EventTokenCategories:
		event.Categories = append(event.Categories, strings.Split(value, ",")...)
	case model.EventTokenGeo:
		if event.Geo != nil {
			return fmt.Errorf(icalerr.ErrDuplicatePropertyInComponentFormat, icalerr.ErrDuplicatePropertyInComponent, propertyName, eventLocation)
		}
		geo, err := parseGeo(value)
		if err != nil {
			return err
		}
		event.Geo = &geo
	case model.EventTokenRRule:
		rule, err := rrule.ParseRRule(value)
		if err != nil {
			return fmt.Errorf("%w: %w", icalerr.ErrInvalidRRule, err)
		}
		return setOnceProperty(&event.RRule, rule, propertyName, eventLocation)
	case model.EventTokenAttach:
		attachment, err := parseAttachment(value, params)
		if err != nil {
			return err
		}
		event.Attach = append(event.Attach, *attachment)
		return nil
	case model.EventTokenAttendee:
		attendee, err := parseAttendee(value, params)
		if err != nil {
			return err
		}
		event.Attendees = append(event.Attendees, *attendee)
		return nil
	case model.EventTokenExceptionDates:
		return appendCommaSeparatedTimePropertyWithParams(&event.ExceptionDates, value, params, propertyName, eventLocation)
	case model.EventTokenRequestStatus:
		event.RequestStatus = append(event.RequestStatus, value)
	case model.EventTokenRelated:
		event.Related = append(event.Related, value)
	case model.EventTokenResources:
		event.Resources = append(event.Resources, strings.Split(value, ",")...)
	case model.EventTokenRdate:
		return appendCommaSeparatedTimePropertyWithParams(&event.Rdate, value, params, propertyName, eventLocation)
	default:
		appendExtensionProperty(&event.XProp, &event.IANAProp, propertyName, value, params)
		return nil
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
	if method == "" && event.Start.IsZero() {
		return icalerr.ErrMissingEventDTStartProperty
	}
	return nil
}
