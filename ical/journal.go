package ical

import (
	"fmt"

	"github.com/michael-gallo/simpleical/internal/icalerr"
	"github.com/michael-gallo/simpleical/model"
	"github.com/michael-gallo/simpleical/rrule"
)

const journalLocation = "Journal"

// parseJournalProperty parses a single property line and adds it to the provided journal.
func parseJournalProperty(propertyName string, value string, params map[string]string, journal *model.Journal) error {
	switch propertyName {
	case model.PropDTStamp:
		return setOnceUTCTimePropertyWithParams(&journal.DTStamp, value, params, propertyName, journalLocation)
	case model.PropUID:
		return setOnceProperty(&journal.UID, value, propertyName, journalLocation)
	case model.PropClass:
		class, err := parseClass(value)
		if err != nil {
			return err
		}
		return setOnceProperty(&journal.Class, class, propertyName, journalLocation)
	case model.PropCreated:
		return setOnceUTCTimePropertyWithParams(&journal.Created, value, params, propertyName, journalLocation)
	case model.PropDTStart:
		return setOnceTimePropertyWithParams(&journal.DTStart, value, params, propertyName, journalLocation)
	case model.PropLastModified:
		return setOnceUTCTimePropertyWithParams(&journal.LastModified, value, params, propertyName, journalLocation)
	case model.PropOrganizer:
		organizer, err := parseOrganizer(value, params)
		if err != nil {
			return err
		}
		return setOnceProperty(&journal.Organizer, organizer, propertyName, journalLocation)
	case model.PropRecurrenceID:
		if err := setOnceTimePropertyWithParams(&journal.RecurrenceID, value, params, propertyName, journalLocation); err != nil {
			return err
		}
		journal.RecurrenceIDRange = params[model.ParamRange]
		return nil
	case model.PropSequence:
		return setOnceIntProperty(&journal.Sequence, value, propertyName, journalLocation)
	case model.PropStatus:
		status, err := parseJournalStatus(value)
		if err != nil {
			return err
		}
		return setOnceProperty(&journal.Status, status, propertyName, journalLocation)
	case model.PropSummary:
		return setOnceTextProperty(&journal.Summary, value, params, propertyName, journalLocation)
	case model.PropURL:
		return setOnceProperty(&journal.URL, value, propertyName, journalLocation)
	case model.PropRRule:
		rule, err := rrule.ParseRRule(value)
		if err != nil {
			return fmt.Errorf("%w: %w", icalerr.ErrInvalidRRule, err)
		}
		return setOnceProperty(&journal.RRule, rule, propertyName, journalLocation)

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
		return appendTextListProperty(&journal.Categories, value)
	case model.PropComment:
		return appendTextProperty(&journal.Comment, value, params)
	case model.PropContact:
		journal.Contacts = append(journal.Contacts, value)
	case model.PropDescription:
		return appendTextProperty(&journal.Description, value, params)
	case model.PropExDate:
		return appendCommaSeparatedTimePropertyWithParams(&journal.ExceptionDates, value, params, propertyName, journalLocation)
	case model.PropRelatedTo:
		return appendRelatedToProperty(&journal.Related, value, params)
	case model.PropRDate:
		return appendRDateProperty(&journal.Rdate, value, params, propertyName, journalLocation)
	case model.PropRequestStatus:
		journal.RequestStatus = append(journal.RequestStatus, value)
	default:
		appendExtensionProperty(&journal.XProp, &journal.IANAProp, propertyName, value, params)
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
