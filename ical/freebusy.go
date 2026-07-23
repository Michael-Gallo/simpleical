package ical

import (
	"fmt"
	"strings"

	"github.com/michael-gallo/simpleical/internal/icalerr"
	"github.com/michael-gallo/simpleical/model"
)

const freeBusyLocation = "FreeBusy"

// parseFreeBusyProperty parses a single property line and adds it to the provided freebusy.
func parseFreeBusyProperty(propertyName string, value string, params map[string]string, freeBusy *model.FreeBusy) error {
	switch propertyName {
	case model.PropDTStamp:
		return setOnceUTCTimePropertyWithParams(&freeBusy.DTStamp, value, params, propertyName, freeBusyLocation)
	case model.PropUID:
		return setOnceProperty(&freeBusy.UID, value, propertyName, freeBusyLocation)
	case model.PropContact:
		return setOnceProperty(&freeBusy.Contact, value, propertyName, freeBusyLocation)
	case model.PropDTStart:
		return setOnceTimePropertyWithParams(&freeBusy.DTStart, value, params, propertyName, freeBusyLocation)
	case model.PropDTEnd:
		return setOnceTimePropertyWithParams(&freeBusy.DTEnd, value, params, propertyName, freeBusyLocation)
	case model.PropOrganizer:
		organizer, err := parseOrganizer(value, params)
		if err != nil {
			return err
		}
		return setOnceProperty(&freeBusy.Organizer, organizer, propertyName, freeBusyLocation)
	case model.PropURL:
		return setOnceProperty(&freeBusy.URL, value, propertyName, freeBusyLocation)

	// Repeatable properties
	case model.PropAttendee:
		attendee, err := parseAttendee(value, params)
		if err != nil {
			return err
		}
		freeBusy.Attendees = append(freeBusy.Attendees, *attendee)
	case model.PropComment:
		return appendTextProperty(&freeBusy.Comment, value, params)
	case model.PropFreeBusy:
		fbType := model.FreeBusyStatusBusy
		if raw := params[model.ParamFBType]; raw != "" {
			parsed, err := parseFreeBusyStatus(raw)
			if err != nil {
				return err
			}
			fbType = parsed
		}
		for part := range strings.SplitSeq(value, ",") {
			period, err := parsePeriod(part)
			if err != nil {
				return fmt.Errorf("%w: %s property %s in iCal: %w", icalerr.ErrParseErrorInComponent, freeBusyLocation, propertyName, err)
			}
			freeBusy.FreeBusy = append(freeBusy.FreeBusy, model.FreeBusyTime{
				Start:    period.Start,
				End:      period.End,
				Duration: period.Duration,
				FBType:   fbType,
			})
		}
	case model.PropRequestStatus:
		freeBusy.RequestStatus = append(freeBusy.RequestStatus, value)
	default:
		appendExtensionProperty(&freeBusy.XProp, &freeBusy.IANAProp, propertyName, value, params)
	}
	return nil
}

// validateFreeBusy ensures that all required values are present for a freebusy.
func validateFreeBusy(freeBusy *model.FreeBusy) error {
	if freeBusy.UID == "" {
		return icalerr.ErrMissingFreeBusyUIDProperty
	}
	if freeBusy.DTStamp.IsZero() {
		return icalerr.ErrMissingFreeBusyDTStampProperty
	}
	// VFREEBUSY DTSTART/DTEND must be UTC DATE-TIME, not DATE.
	if (!freeBusy.DTStart.IsZero() && freeBusy.DTStart.IsDate()) ||
		(!freeBusy.DTEnd.IsZero() && freeBusy.DTEnd.IsDate()) {
		return icalerr.ErrUTCValueRequired
	}
	return nil
}
