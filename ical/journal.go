package ical

import (
	"fmt"
	"strings"

	"github.com/michael-gallo/simpleical/internal/icalerr"
	"github.com/michael-gallo/simpleical/model"
)

const journalLocation = "Journal"

// parseJournalProperty parses a single property line and adds it to the provided journal.
func parseJournalProperty(propertyName string, value string, params map[string]string, journal *model.Journal) error {
	switch propertyName {
	case model.PropDTStamp:
		return setOnceTimePropertyWithParams(&journal.DTStamp, value, params, propertyName, journalLocation)
	case model.PropUID:
		return setOnceProperty(&journal.UID, value, propertyName, journalLocation)
	case model.PropClass:
		return setOnceProperty(&journal.Class, model.Class(value), propertyName, journalLocation)
	case model.PropCreated:
		return setOnceTimePropertyWithParams(&journal.Created, value, params, propertyName, journalLocation)
	case model.PropDTStart:
		return setOnceTimePropertyWithParams(&journal.DTStart, value, params, propertyName, journalLocation)
	case model.PropLastModified:
		return setOnceTimePropertyWithParams(&journal.LastModified, value, params, propertyName, journalLocation)
	case model.PropOrganizer:
		organizer, err := parseOrganizer(value, params)
		if err != nil {
			return err
		}
		return setOnceProperty(&journal.Organizer, organizer, propertyName, journalLocation)
	case model.PropRecurrenceID:
		return setOnceTimePropertyWithParams(&journal.RecurrenceID, value, params, propertyName, journalLocation)
	case model.PropSequence:
		return setOnceIntProperty(&journal.Sequence, value, propertyName, journalLocation)
	case model.PropStatus:
		return setOnceProperty(&journal.Status, model.JournalStatus(value), propertyName, journalLocation)
	case model.PropSummary:
		return setOnceProperty(&journal.Summary, value, propertyName, journalLocation)
	case model.PropURL:
		return setOnceProperty(&journal.URL, value, propertyName, journalLocation)

	// Repeatable properties
	case model.PropAttach:
		attachment, err := parseAttachment(value, params)
		if err != nil {
			return err
		}
		journal.Attach = append(journal.Attach, *attachment)
		return nil
	case model.PropAttendee:
		attendee, err := parseAttendee(value, params)
		if err != nil {
			return err
		}
		journal.Attendees = append(journal.Attendees, *attendee)
	case model.PropCategories:
		journal.Categories = append(journal.Categories, strings.Split(value, ",")...)
	case model.PropComment:
		journal.Comment = append(journal.Comment, value)
	case model.PropContact:
		journal.Contacts = append(journal.Contacts, value)
	case model.PropDescription:
		journal.Description = append(journal.Description, value)
	case model.PropExDate:
		return appendCommaSeparatedTimePropertyWithParams(&journal.ExceptionDates, value, params, propertyName, journalLocation)
	case model.PropRelatedTo:
		journal.Related = append(journal.Related, value)
	case model.PropRDate:
		return appendCommaSeparatedTimePropertyWithParams(&journal.Rdate, value, params, propertyName, journalLocation)
	case model.PropRequestStatus:
		journal.RequestStatus = append(journal.RequestStatus, value)
	default:
		return fmt.Errorf("%w: %s", icalerr.ErrInvalidJournalProperty, propertyName)
	}
	return nil
}

// validateJournal ensures that all required values are present for a journal.
func validateJournal(journal *model.Journal) error {
	if journal.UID == "" {
		return icalerr.ErrMissingJournalUIDProperty
	}
	if journal.DTStamp.IsZero() {
		return icalerr.ErrMissingJournalDTStampProperty
	}
	return nil
}
