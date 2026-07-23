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

	//go:embed test_data/freebusy/test_freebusy.ical
	testFreeBusyInput string

	//go:embed test_data/freebusy/test_freebusy_missing_uid.ical
	testFreeBusyMissingUIDInput string
	//go:embed test_data/freebusy/test_freebusy_missing_dtstamp.ical
	testFreeBusyMissingDTStampInput string
	//go:embed test_data/freebusy/test_freebusy_duplicate_uid.ical
	testFreeBusyDuplicateUIDInput string
	//go:embed test_data/freebusy/test_freebusy_invalid_freebusy.ical
	testFreeBusyInvalidFreeBusyInput string
	//go:embed test_data/freebusy/valid_test_freebusy_all_day_date.ical
	testFreeBusyAllDayDateInput string
)

func TestValidFreeBusy(t *testing.T) {
	testCases := []struct {
		name             string
		input            string
		expectedCalendar *model.Calendar
	}{
		{
			name:  "Valid free busy",
			input: testFreeBusyInput,
			expectedCalendar: &model.Calendar{
				ProdID:  "-//Test//FreeBusy Calendar//EN",
				Version: "2.0",
				FreeBusys: []model.FreeBusy{
					{
						UID:     "freebusy123@example.com",
						DTStamp: utcDT(2024, 1, 1, 0, 0, 0),
						Contact: "John Doe, Scheduling Assistant, +1-555-0123",
						DTStart: utcDT(2024, 1, 1, 0, 0, 0),
						DTEnd:   utcDT(2024, 1, 31, 23, 59, 59),
						Organizer: &model.Organizer{
							CommonName: "Calendar Owner",
							CalAddress: &url.URL{Scheme: "mailto", Opaque: "owner@example.com"},
						},
						Attendees: []model.Attendee{
							{CalAddress: &url.URL{Scheme: "mailto", Opaque: "user1@example.com"}},
							{CalAddress: &url.URL{Scheme: "mailto", Opaque: "user2@example.com"}},
						},
						Comment: texts("Available for meetings during business hours"),
						FreeBusy: []model.FreeBusyTime{
							{
								Start:  utcDT(2024, 1, 1, 9, 0, 0),
								End:    utcDT(2024, 1, 1, 12, 0, 0),
								FBType: model.FreeBusyStatusBusy,
							},
							{
								Start:  utcDT(2024, 1, 1, 13, 0, 0),
								End:    utcDT(2024, 1, 1, 17, 0, 0),
								FBType: model.FreeBusyStatusBusy,
							},
							{
								Start:  utcDT(2024, 1, 2, 10, 0, 0),
								End:    utcDT(2024, 1, 2, 11, 0, 0),
								FBType: model.FreeBusyStatusBusyTentative,
							},
						},
						URL: "https://calendar.example.com/freebusy/123",
					},
				},
			},
		},
		{
			name:  "Valid all-day VFREEBUSY with VALUE=DATE",
			input: testFreeBusyAllDayDateInput,
			expectedCalendar: &model.Calendar{
				ProdID:  "-//Test//FreeBusy Calendar//EN",
				Version: "2.0",
				FreeBusys: []model.FreeBusy{
					{
						UID:     "freebusy-date@example.com",
						DTStamp: utcDT(1997, 9, 1, 13, 0, 0),
						DTStart: dateDT(6, 28),
						DTEnd:   dateDT(7, 9),
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

func TestInvalidFreeBusy(t *testing.T) {
	testCases := []struct {
		name        string
		input       string
		expectedErr error
	}{
		{
			name:        "VFREEBUSY missing UID",
			input:       testFreeBusyMissingUIDInput,
			expectedErr: icalerr.ErrMissingFreeBusyUIDProperty,
		},
		{
			name:        "VFREEBUSY missing DTSTAMP",
			input:       testFreeBusyMissingDTStampInput,
			expectedErr: icalerr.ErrMissingFreeBusyDTStampProperty,
		},
		{
			name:        "VFREEBUSY invalid FREEBUSY format",
			input:       testFreeBusyInvalidFreeBusyInput,
			expectedErr: icalerr.ErrInvalidFreeBusyFormat,
		},
		{
			name:        "VFREEBUSY duplicate UID",
			input:       testFreeBusyDuplicateUIDInput,
			expectedErr: icalerr.ErrDuplicatePropertyInComponent,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			calendar, err := ical.ReadSingle(strings.NewReader(tc.input))
			assert.Nil(t, calendar)
			require.ErrorIs(t, err, tc.expectedErr)
		})
	}
}
