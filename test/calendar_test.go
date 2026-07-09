package test

import (
	_ "embed"
	"errors"
	"io"
	"net/url"
	"testing"
	"time"

	"github.com/michael-gallo/simpleical/ical"
	"github.com/michael-gallo/simpleical/internal/icalerr"
	"github.com/michael-gallo/simpleical/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var errBoom = errors.New("boom")

// errReader always fails on Read, used to verify I/O errors are preserved.
type errReader struct {
	err error
}

func (r errReader) Read([]byte) (int, error) {
	return 0, r.err
}

var (
	//go:embed test_data/calendar/valid_calendar_with_event_and_timezone.ical
	testIcalWithEventAndTimezoneInput string
	//go:embed test_data/calendar/valid_calendar.ical
	testValidCalendarInput string
	//go:embed test_data/calendar/valid_empty_calendar.ical
	testEmptyCalendarInput string
	//go:embed test_data/calendar/invalid_calendar_double_begin.ical
	testCalendarDoubleBeginInput string
	//go:embed test_data/calendar/valid_calendar_trailing_whitespace.ical
	testTrailingWithSpaceInput string
	//go:embed test_data/calendar/no_begin_calendar.ical
	testInvalidBeginCalendarInput string
	//go:embed test_data/calendar/no_end_calendar.ical
	testInvalidEndCalendarInput string
	//go:embed test_data/calendar/empty_line_calendar.ical
	testInvalidEmptyLineCalendarInput string
	//go:embed test_data/calendar/calendar_missing_version.ical
	testCalendarMissingVersionInput string
	//go:embed test_data/calendar/calendar_missing_prodid.ical
	testCalendarMissingProdIDInput string
	//go:embed test_data/calendar/valid_calendar_with_carriage_returns.ical
	testValidCalendarWithCarriageReturnsInput string
	//go:embed test_data/calendar/valid_calendar_folded_lines.ical
	testValidCalendarFoldedLinesInput string
)

func TestParseCalendarSuccess(t *testing.T) {
	fullExpectedCalendar := &model.Calendar{
		ProdID:   "-//Event//Event Calendar//EN",
		Version:  "2.0",
		Method:   "REQUEST",
		CalScale: "GREGORIAN",
		Events: []model.Event{
			{
				DTStamp:     time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC),
				UID:         "13235@example.com",
				Comment:     []string{"I Am", "A Comment"},
				Start:       time.Date(2025, time.September, 28, 18, 30, 0, 0, time.UTC),
				End:         time.Date(2025, time.September, 28, 20, 30, 0, 0, time.UTC),
				Summary:     "Event Summary",
				Description: "Event Description",
				Location:    "555 Fake Street",
				Organizer: &model.Organizer{
					CommonName: "Org",
					CalAddress: &url.URL{Scheme: "mailto", Opaque: "hello@world"},
				},
				Status:       model.EventStatusConfirmed,
				Sequence:     1,
				Transp:       model.EventTranspOpaque,
				Contacts:     []string{"Jim Dolittle, ABC Industries, +1-919-555-1234"},
				LastModified: time.Date(2021, time.January, 1, 0, 0, 0, 0, time.UTC),
				Categories:   []string{"first", "second", "third"},
				Geo:          &[2]float64{37.386013, -122.082932},
			},
		},
		TimeZones: []model.TimeZone{
			{
				TimeZoneID: "America/Detroit",
				Standard: []model.TimeZoneProperty{
					{
						TimeZoneOffsetFrom: "+0000",
						TimeZoneOffsetTo:   "+0000",
						DTStart:            time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC),
					},
				},
			},
		},
	}
	testCases := []struct {
		name             string
		input            string
		expectedCalendar *model.Calendar
	}{
		{
			name:             "Valid iCal event",
			input:            testIcalWithEventAndTimezoneInput,
			expectedCalendar: fullExpectedCalendar,
		},
		{
			name:  "Valid calendar",
			input: testValidCalendarInput,
			expectedCalendar: &model.Calendar{
				ProdID:   "-//Event//Event Calendar//EN",
				Version:  "2.0",
				Method:   "REQUEST",
				CalScale: "GREGORIAN",
			},
		},
		{
			name:  "No VEVENT block",
			input: testEmptyCalendarInput,
			expectedCalendar: &model.Calendar{
				Version: "2.0",
				ProdID:  "Id",
				Events:  nil,
			},
		},
		{
			name:  "Calendar with trailing space",
			input: testTrailingWithSpaceInput,
			expectedCalendar: &model.Calendar{
				ProdID:   "-//Event//Event Calendar//EN",
				Version:  "2.0",
				Method:   "REQUEST",
				CalScale: "GREGORIAN",
			},
		},
		{
			name:             "Calendar with carriage returns",
			input:            testValidCalendarWithCarriageReturnsInput,
			expectedCalendar: fullExpectedCalendar,
		},
		{
			name:  "Calendar with folded content lines",
			input: testValidCalendarFoldedLinesInput,
			expectedCalendar: &model.Calendar{
				ProdID:   "-//Event//Event Calendar//EN",
				Version:  "2.0",
				Method:   "REQUEST",
				CalScale: "GREGORIAN",
				Events: []model.Event{
					{
						DTStamp:     time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC),
						UID:         "13235@example.com",
						Start:       time.Date(2025, time.September, 28, 18, 30, 0, 0, time.UTC),
						End:         time.Date(2025, time.September, 28, 20, 30, 0, 0, time.UTC),
						Summary:     "Event Summary",
						Description: "This is a long description that exists on a long line.",
						Location:    "555 Fake Street",
					},
				},
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.name == "Calendar with carriage returns" {
				assert.Contains(t, tc.input, "\r\n", "fixture must retain CRLF line endings")
			}
			calendar, err := ical.FromString(tc.input)
			require.NoError(t, err)
			assert.Equal(t, *tc.expectedCalendar, *calendar)
		})
	}
}

func TestParseCalendarError(t *testing.T) {
	testCases := []struct {
		name        string
		input       string
		expectedErr error
	}{
		{
			name:        "Calendar with no BEGIN:VCALENDAR",
			input:       testInvalidBeginCalendarInput,
			expectedErr: icalerr.ErrInvalidCalendarFormatMissingBegin,
		},
		{
			name:        "Calendar with no END:VCALENDAR",
			input:       testInvalidEndCalendarInput,
			expectedErr: icalerr.ErrInvalidCalendarFormatMissingEnd,
		},
		{
			name:        "Empty line in calendar",
			input:       testInvalidEmptyLineCalendarInput,
			expectedErr: icalerr.ErrInvalidCalendarEmptyLine,
		},
		{
			name:        "Calendar missing VERSION property",
			input:       testCalendarMissingVersionInput,
			expectedErr: icalerr.ErrMissingCalendarVersionProperty,
		},
		{
			name:        "Calendar missing PRODID property",
			input:       testCalendarMissingProdIDInput,
			expectedErr: icalerr.ErrMissingCalendarProdIDProperty,
		},
		{
			name:        "Empty input",
			input:       "",
			expectedErr: icalerr.ErrNoCalendarFound,
		},
		{
			name:        "Two BEGIN:VCALENDAR at top",
			input:       testCalendarDoubleBeginInput,
			expectedErr: icalerr.ErrNestedBeginVCalendar,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			calendar, err := ical.FromString(tc.input)
			require.ErrorIs(t, err, tc.expectedErr)
			assert.Nil(t, calendar)
		})
	}
}

func TestReadPreservesInitialScanIOError(t *testing.T) {
	calendar, err := ical.Read(errReader{err: errBoom})
	require.Error(t, err)
	require.ErrorIs(t, err, errBoom)
	require.NotErrorIs(t, err, icalerr.ErrNoCalendarFound)
	assert.Nil(t, calendar)
}

func TestReadEmptyReaderReturnsNoCalendarFound(t *testing.T) {
	calendar, err := ical.Read(io.MultiReader())
	require.ErrorIs(t, err, icalerr.ErrNoCalendarFound)
	assert.Nil(t, calendar)
}
