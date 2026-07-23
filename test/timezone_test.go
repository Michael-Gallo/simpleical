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

	//go:embed test_data/timezones/test_timezone.ical
	testTimezoneInput string

	//go:embed test_data/timezones/test_timezone_missing_tzid.ical
	testTimezoneMissingTZIDInput string
	//go:embed test_data/timezones/test_timezone_duplicate_tzid.ical
	testTimezoneDuplicateTZIDInput string
	//go:embed test_data/timezones/test_timezone_duplicate_tzoffsetfrom.ical
	testTimezoneDuplicateTZOffsetFromInput string
	//go:embed test_data/timezones/test_timezone_duplicate_tzoffsetto.ical
	testTimezoneDuplicateTZOffsetToInput string
	//go:embed test_data/timezones/test_timezone_invalid_dtstart.ical
	testTimezoneInvalidDTStartInput string
	//go:embed test_data/timezones/test_timezone_utc_dtstart.ical
	testTimezoneUTCDTStartInput string
	//go:embed test_data/timezones/test_timezone_utc_rdate.ical
	testTimezoneUTCRDateInput string
	//go:embed test_data/timezones/test_timezone_duplicate_rrule.ical
	testTimezoneDuplicateRRuleInput string
)

func TestValidTimezone(t *testing.T) {
	testCases := []struct {
		name             string
		input            string
		expectedCalendar *model.Calendar
	}{
		{
			name:  "Valid VTIMEZONE",
			input: testTimezoneInput,
			expectedCalendar: &model.Calendar{
				ProdID:  "-//Test//Timezone Calendar//EN",
				Version: "2.0",
				TimeZones: []model.TimeZone{
					{
						TimeZoneID:  "America/New_York",
						LastMod:     utcDT(2024, 1, 1, 0, 0, 0),
						TimeZoneURL: &url.URL{Scheme: "http", Host: "tzurl.org", Path: "/zoneinfo-outlook/America/New_York"},
						Standard: []model.TimeZoneProperty{
							{
								TimeZoneOffsetFrom: model.UTCOffset("-0400"),
								TimeZoneOffsetTo:   model.UTCOffset("-0500"),
								DTStart:            floatDT(2024, 1, 2),
								TimeZoneName:       []string{"EST"},
								Comment:            texts("Eastern Standard Time"),
								Rdate:              []model.DateTime{floatDT(2024, 1, 2)},
							},
						},
						Daylight: []model.TimeZoneProperty{
							{
								TimeZoneOffsetFrom: model.UTCOffset("-0500"),
								TimeZoneOffsetTo:   model.UTCOffset("-0400"),
								DTStart:            floatDT(2024, 3, 2),
								TimeZoneName:       []string{"EDT"},
								Comment:            texts("Eastern Daylight Time"),
								Rdate:              []model.DateTime{floatDT(2024, 3, 2)},
							},
						},
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

func TestInvalidTimezone(t *testing.T) {
	testCases := []struct {
		name        string
		input       string
		expectedErr error
		errContains string
	}{
		{
			name:        "VTIMEZONE missing TZID",
			input:       testTimezoneMissingTZIDInput,
			expectedErr: icalerr.ErrMissingTimezoneTZIDProperty,
		},
		{
			name:        "VTIMEZONE invalid DTSTART",
			input:       testTimezoneInvalidDTStartInput,
			expectedErr: icalerr.ErrParseErrorInComponent,
		},
		{
			name:        "VTIMEZONE STANDARD DTSTART with UTC designator",
			input:       testTimezoneUTCDTStartInput,
			expectedErr: icalerr.ErrTimezoneLocalTimeRequired,
		},
		{
			name:        "VTIMEZONE STANDARD RDATE with UTC designator",
			input:       testTimezoneUTCRDateInput,
			expectedErr: icalerr.ErrTimezoneLocalTimeRequired,
		},
		{
			name:        "VTIMEZONE duplicate TZID",
			input:       testTimezoneDuplicateTZIDInput,
			expectedErr: icalerr.ErrDuplicatePropertyInComponent,
		},
		{
			name:        "VTIMEZONE duplicate TZOFFSETFROM",
			input:       testTimezoneDuplicateTZOffsetFromInput,
			expectedErr: icalerr.ErrDuplicatePropertyInComponent,
			errContains: "TZOFFSETFROM",
		},
		{
			name:        "VTIMEZONE duplicate TZOFFSETTO",
			input:       testTimezoneDuplicateTZOffsetToInput,
			expectedErr: icalerr.ErrDuplicatePropertyInComponent,
			errContains: "TZOFFSETTO",
		},
		{
			name:        "VTIMEZONE duplicate RRULE",
			input:       testTimezoneDuplicateRRuleInput,
			expectedErr: icalerr.ErrDuplicatePropertyInComponent,
			errContains: "RRULE",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			calendar, err := ical.ReadSingle(strings.NewReader(tc.input))
			require.ErrorIs(t, err, tc.expectedErr)
			if tc.errContains != "" {
				require.ErrorContains(t, err, tc.errContains)
			}
			assert.Nil(t, calendar)
		})
	}
}
