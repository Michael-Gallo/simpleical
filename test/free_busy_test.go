package test

import (
	_ "embed"
	"net/url"
	"strings"
	"testing"
	"time"

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
						DTStamp: time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC),
						Contact: "John Doe, Scheduling Assistant, +1-555-0123",
						DTStart: time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC),
						DTEnd:   time.Date(2024, time.January, 31, 23, 59, 59, 0, time.UTC),
						Organizer: &model.Organizer{
							CommonName: "Calendar Owner",
							CalAddress: &url.URL{Scheme: "mailto", Opaque: "owner@example.com"},
						},
						Attendees: []model.Attendee{
							{CalAddress: &url.URL{Scheme: "mailto", Opaque: "user1@example.com"}},
							{CalAddress: &url.URL{Scheme: "mailto", Opaque: "user2@example.com"}},
						},
						Comment: []string{"Available for meetings during business hours"},
						FreeBusy: []model.FreeBusyTime{
							{
								Start:  time.Date(2024, time.January, 1, 9, 0, 0, 0, time.UTC),
								End:    time.Date(2024, time.January, 1, 12, 0, 0, 0, time.UTC),
								Status: model.FreeBusyStatusBusy,
							},
							{
								Start:  time.Date(2024, time.January, 1, 13, 0, 0, 0, time.UTC),
								End:    time.Date(2024, time.January, 1, 17, 0, 0, 0, time.UTC),
								Status: model.FreeBusyStatusBusy,
							},
							{
								Start:  time.Date(2024, time.January, 2, 10, 0, 0, 0, time.UTC),
								End:    time.Date(2024, time.January, 2, 11, 0, 0, 0, time.UTC),
								Status: model.FreeBusyStatusBusyTentative,
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
						DTStamp: time.Date(1997, time.September, 1, 13, 0, 0, 0, time.UTC),
						DTStart: time.Date(2007, time.June, 28, 0, 0, 0, 0, time.UTC),
						DTEnd:   time.Date(2007, time.July, 9, 0, 0, 0, 0, time.UTC),
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
