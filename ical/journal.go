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
	switch model.JournalToken(propertyName) {
	case model.JournalTokenDTStamp:
		return setOnceTimeProperty(&journal.DTStamp, value, propertyName, journalLocation)
	case model.JournalTokenUID:
		return setOnceProperty(&journal.UID, value, propertyName, journalLocation)
	case model.JournalTokenClass:
		return setOnceProperty(&journal.Class, model.JournalClass(value), propertyName, journalLocation)
	case model.JournalTokenCreated:
		return setOnceTimeProperty(&journal.Created, value, propertyName, journalLocation)
	case model.JournalTokenDTStart:
		return setOnceTimePropertyWithParams(&journal.DTStart, value, params, propertyName, journalLocation)
	case model.JournalTokenLastModified:
		return setOnceTimeProperty(&journal.LastModified, value, propertyName, journalLocation)
	case model.JournalTokenOrganizer:
		organizer, err := parseOrganizer(value, params)
		if err != nil {
			return err
		}
		return setOnceProperty(&journal.Organizer, organizer, propertyName, journalLocation)
	case model.JournalTokenRecurrenceID:
		return setOnceTimePropertyWithParams(&journal.RecurrenceID, value, params, propertyName, journalLocation)
	case model.JournalTokenSequence:
		return setOnceIntProperty(&journal.Sequence, value, propertyName, journalLocation)
	case model.JournalTokenStatus:
		return setOnceProperty(&journal.Status, model.JournalStatus(value), propertyName, journalLocation)
	case model.JournalTokenSummary:
		return setOnceProperty(&journal.Summary, value, propertyName, journalLocation)
	case model.JournalTokenURL:
		return setOnceProperty(&journal.URL, value, propertyName, journalLocation)

	// Repeatable properties
	case model.JournalTokenAttach:
		attachment, err := parseAttachment(value, params)
		if err != nil {
			return err
		}
		journal.Attach = append(journal.Attach, *attachment)
		return nil
	case model.JournalTokenAttendee:
		attendee, err := parseAttendee(value, params)
		if err != nil {
			return err
		}
		journal.Attendees = append(journal.Attendees, *attendee)
	case model.JournalTokenCategories:
		journal.Categories = append(journal.Categories, strings.Split(value, ",")...)
	case model.JournalTokenComment:
		journal.Comment = append(journal.Comment, value)
	case model.JournalTokenContact:
		journal.Contacts = append(journal.Contacts, value)
	case model.JournalTokenDescription:
		journal.Description = append(journal.Description, value)
	case model.JournalTokenExceptionDates:
		return appendTimePropertyWithParams(&journal.ExceptionDates, value, params, propertyName, journalLocation)
	case model.JournalTokenRelated:
		journal.Related = append(journal.Related, value)
	case model.JournalTokenRdate:
		return appendTimePropertyWithParams(&journal.Rdate, value, params, propertyName, journalLocation)
	case model.JournalTokenRequestStatus:
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
