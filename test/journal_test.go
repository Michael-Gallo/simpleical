package test

import (
	_ "embed"
	"net/url"
	"strings"
	"testing"

	"github.com/michael-gallo/simpleical/ical"
	"github.com/michael-gallo/simpleical/internal/icalerr"
	"github.com/michael-gallo/simpleical/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	//go:embed test_data/journals/test_journal.ical
	testJournalInput string
	//go:embed test_data/journals/test_journal_missing_uid.ical
	testJournalMissingUIDInput string
	//go:embed test_data/journals/test_journal_duplicate_uid.ical
	testJournalDuplicateUIDInput string
	//go:embed test_data/journals/test_journal_multiple_exdates.ical
	testJournalMultipleExdatesInput string
	//go:embed test_data/journals/test_journal_missing_dtstamp.ical
	testJournalMissingDTStampInput string
	//go:embed test_data/journals/test_journal_duplicate_status.ical
	testJournalDuplicateStatusInput string
	//go:embed test_data/journals/test_journal_duplicate_organizer.ical
	testJournalDuplicateOrganizerInput string
	//go:embed test_data/journals/valid_test_journal_all_day_date.ical
	testJournalAllDayDateInput string
	//go:embed test_data/journals/valid_test_journal_attach_rdate_related.ical
	testJournalAttachRdateRelatedInput string
)

func TestValidJournal(t *testing.T) {
	testCases := []struct {
		name             string
		input            string
		expectedCalendar *model.Calendar
	}{
		{
			name:  "Valid VJOURNAL",
			input: testJournalInput,
			expectedCalendar: &model.Calendar{
				ProdID:  "-//Test//Journal Calendar//EN",
				Version: "2.0",
				Journals: []model.Journal{
					{
						UID:          "journal123@example.com",
						DTStamp:      utcDT(2024, 1, 1, 0, 0, 0),
						Summary:      text("Project status update"),
						Description:  texts("Completed the initial research phase", "Identified key stakeholders and requirements"),
						Class:        model.ClassConfidential,
						Status:       model.JournalStatusFinal,
						Created:      utcDT(2024, 1, 1, 9, 0, 0),
						LastModified: utcDT(2024, 1, 15, 12, 0, 0),
						DTStart:      utcDT(2024, 1, 1, 9, 0, 0),
						Organizer: &model.Organizer{
							CommonName: "Project Lead",
							CalAddress: &url.URL{Scheme: "mailto", Opaque: "lead@example.com"},
						},
						Attendees: []model.Attendee{
							{CalAddress: &url.URL{Scheme: "mailto", Opaque: "stakeholder1@example.com"}},
							{CalAddress: &url.URL{Scheme: "mailto", Opaque: "stakeholder2@example.com"}},
						},
						Contacts:   []string{"Jane Doe, Project Manager, +1-555-0456"},
						Categories: []string{"work", "project", "status"},
						Comment:    texts("This journal entry documents the completion of Phase 1"),
						URL:        "https://project.example.com/journal/123",
					},
				},
			},
		},
		{
			name:  "Valid VJOURNAL with Multiple Exception Dates",
			input: testJournalMultipleExdatesInput,
			expectedCalendar: &model.Calendar{
				ProdID:  "-//Test//Journal Calendar//EN",
				Version: "2.0",
				Journals: []model.Journal{
					{
						UID:         "journal123@example.com",
						DTStamp:     utcDT(2024, 1, 1, 0, 0, 0),
						DTStart:     utcDT(2024, 1, 1, 9, 0, 0),
						Summary:     text("Journal with Multiple Exception Dates"),
						Description: texts("This journal has multiple exception dates to test the append functionality"),
						Class:       model.ClassConfidential,
						Status:      model.JournalStatusFinal,
						ExceptionDates: []model.DateTime{
							utcDT(2024, 1, 15, 9, 0, 0),
							utcDT(2024, 1, 22, 9, 0, 0),
							utcDT(2024, 1, 29, 9, 0, 0),
						},
					},
				},
			},
		},
		{
			name:  "Valid all-day VJOURNAL with VALUE=DATE",
			input: testJournalAllDayDateInput,
			expectedCalendar: &model.Calendar{
				ProdID:  "-//Test//Journal Calendar//EN",
				Version: "2.0",
				Journals: []model.Journal{
					{
						UID:     "journal-date@example.com",
						DTStamp: utcDT(1997, 9, 1, 13, 0, 0),
						Summary: text("All-day journal"),
						DTStart: dateDT(6, 28),
					},
				},
			},
		},
		{
			name:  "Valid VJOURNAL with ATTACH, SEQUENCE, RECURRENCE-ID, RDATE, RELATED-TO, REQUEST-STATUS",
			input: testJournalAttachRdateRelatedInput,
			expectedCalendar: &model.Calendar{
				ProdID:  "-//Test//Journal Calendar//EN",
				Version: "2.0",
				Journals: []model.Journal{
					{
						UID:          "journal-attach-rdate-related@example.com",
						DTStamp:      utcDT(2024, 1, 1, 0, 0, 0),
						DTStart:      utcDT(2024, 1, 1, 9, 0, 0),
						Summary:      text("Journal with ATTACH RDATE RELATED-TO"),
						Sequence:     2,
						RecurrenceID: utcDT(2024, 1, 1, 9, 0, 0),
						Attach: []model.Attachment{
							{
								FormatType: "application/pdf",
								Value:      model.AttachValueURI,
								URI:        &url.URL{Scheme: "https", Host: "example.com", Path: "/files/journal.pdf"},
							},
						},
						Rdate: []model.RecurrenceDate{
							rdateUTC(2024, 1, 8, 9, 0),
							rdateUTC(2024, 1, 15, 9, 0),
						},
						RequestStatus: []string{"2.0;Success"},
						Related:       related("parent-journal@example.com"),
					},
				},
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			calendar, err := ical.ReadSingle(strings.NewReader(tc.input))
			require.NoError(t, err)
			assert.Equal(t, *tc.expectedCalendar, *calendar)
		})
	}
}

func TestInvalidJournal(t *testing.T) {
	testCases := []struct {
		name        string
		input       string
		expectedErr error
	}{
		{
			name:        "VJOURNAL missing UID",
			input:       testJournalMissingUIDInput,
			expectedErr: icalerr.ErrMissingJournalUIDProperty,
		},
		{
			name:        "VJOURNAL duplicate UID",
			input:       testJournalDuplicateUIDInput,
			expectedErr: icalerr.ErrDuplicatePropertyInComponent,
		},
		{
			name:        "VJOURNAL missing DTSTAMP",
			input:       testJournalMissingDTStampInput,
			expectedErr: icalerr.ErrMissingJournalDTStampProperty,
		},
		{
			name:        "VJOURNAL duplicate STATUS",
			input:       testJournalDuplicateStatusInput,
			expectedErr: icalerr.ErrDuplicatePropertyInComponent,
		},
		{
			name:        "VJOURNAL duplicate ORGANIZER",
			input:       testJournalDuplicateOrganizerInput,
			expectedErr: icalerr.ErrDuplicatePropertyInComponent,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			calendar, err := ical.ReadSingle(strings.NewReader(tc.input))
			require.ErrorIs(t, err, tc.expectedErr)
			assert.Nil(t, calendar)
		})
	}
}
