package ical

import (
	"fmt"
	"strings"

	"github.com/michael-gallo/simpleical/icaldur"
	"github.com/michael-gallo/simpleical/internal/icalerr"
	"github.com/michael-gallo/simpleical/model"
)

const freeBusyLocation = "FreeBusy"

// parseFreeBusyProperty parses a single property line and adds it to the provided freebusy.
func parseFreeBusyProperty(propertyName string, value string, params map[string]string, freeBusy *model.FreeBusy) error {
	switch propertyName {
	case model.PropDTStamp:
		return setOnceTimePropertyWithParams(&freeBusy.DTStamp, value, params, propertyName, freeBusyLocation)
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
		freeBusy.Organizer = organizer
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
		freeBusy.Comment = append(freeBusy.Comment, value)
	case model.PropFreeBusy:
		fbTime, err := parseFreeBusyTime(value)
		if err != nil {
			return err
		}
		freeBusy.FreeBusy = append(freeBusy.FreeBusy, fbTime)
	case model.PropRequestStatus:
		freeBusy.RequestStatus = append(freeBusy.RequestStatus, value)
	default:
		return fmt.Errorf("%w: %s", icalerr.ErrInvalidFreeBusyProperty, propertyName)
	}
	return nil
}

// parseFreeBusyTime parses a FREEBUSY property value into a FreeBusyTime struct.
// Format: "/" separated start/end datetime pair, optionally followed by "/" and status.
// Example: "19970101T180000Z/19970102T070000Z" or "19970101T180000Z/19970102T070000Z/BUSY"
func parseFreeBusyTime(value string) (model.FreeBusyTime, error) {
	// Extract start time (everything before first '/')
	startStr, remaining, found := strings.Cut(value, "/")
	if !found {
		return model.FreeBusyTime{}, fmt.Errorf("%w: %s", icalerr.ErrInvalidFreeBusyFormat, value)
	}

	startTime, err := icaldur.ParseIcalTime(startStr)
	if err != nil {
		return model.FreeBusyTime{}, fmt.Errorf("invalid start time in FREEBUSY property: %w", err)
	}

	// Extract end time and optional status (everything after first '/')
	endStr, statusStr, hasStatus := strings.Cut(remaining, "/")
	endTime, err := icaldur.ParseIcalTime(endStr)
	if err != nil {
		return model.FreeBusyTime{}, fmt.Errorf("invalid end time in FREEBUSY property: %w", err)
	}

	fbTime := model.FreeBusyTime{
		Start: startTime,
		End:   endTime,
	}

	// Optional status parameter
	if hasStatus {
		fbTime.Status = model.FreeBusyStatus(statusStr)
	} else {
		// Default to BUSY if no status specified
		fbTime.Status = model.FreeBusyStatusBusy
	}

	return fbTime, nil
}

// validateFreeBusy ensures that all required values are present for a freebusy.
func validateFreeBusy(freeBusy *model.FreeBusy) error {
	if freeBusy.UID == "" {
		return icalerr.ErrMissingFreeBusyUIDProperty
	}
	if freeBusy.DTStart.IsZero() {
		return icalerr.ErrMissingFreeBusyDTStartProperty
	}
	return nil
}
