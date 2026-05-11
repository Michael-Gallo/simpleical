package test

import (
	_ "embed"
	"net/url"
	"testing"
	"time"

	"github.com/michael-gallo/simpleical/ical"
	"github.com/michael-gallo/simpleical/internal/icalerr"
	"github.com/michael-gallo/simpleical/model"
	"github.com/stretchr/testify/assert"
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
						LastMod:     time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC),
						TimeZoneURL: &url.URL{Scheme: "http", Host: "tzurl.org", Path: "/zoneinfo-outlook/America/New_York"},
						Standard: []model.TimeZoneProperty{
							{
								TimeZoneOffsetFrom: "-0400",
								TimeZoneOffsetTo:   "-0500",
								DTStart:            time.Date(2024, time.January, 1, 2, 0, 0, 0, time.UTC),
								TimeZoneName:       []string{"EST"},
								Comment:            []string{"Eastern Standard Time"},
								Rdate:              []time.Time{time.Date(2024, time.January, 1, 2, 0, 0, 0, time.UTC)},
							},
						},
						Daylight: []model.TimeZoneProperty{
							{
								TimeZoneOffsetFrom: "-0500",
								TimeZoneOffsetTo:   "-0400",
								DTStart:            time.Date(2024, time.March, 1, 2, 0, 0, 0, time.UTC),
								TimeZoneName:       []string{"EDT"},
								Comment:            []string{"Eastern Daylight Time"},
								Rdate:              []time.Time{time.Date(2024, time.March, 1, 2, 0, 0, 0, time.UTC)},
							},
						},
					},
				},
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			calendar, err := ical.FromString(tc.input)
			assert.NoError(t, err)
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
			calendar, err := ical.FromString(tc.input)
			assert.ErrorIs(t, err, tc.expectedErr)
			if tc.errContains != "" {
				assert.ErrorContains(t, err, tc.errContains)
			}
			assert.Nil(t, calendar)
		})
	}
}
