package test

import (
	_ "embed"
	"net/url"
	"testing"
	"time"

	"github.com/michael-gallo/simpleical/ical"
	"github.com/michael-gallo/simpleical/internal/icalerr"
	"github.com/michael-gallo/simpleical/model"
	"github.com/michael-gallo/simpleical/rrule"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (

	//go:embed test_data/events/test_event_invalid_organizer.ical
	testIcalInvalidOrganizerInput string
	//go:embed test_data/events/test_event_full_organizer.ical
	testIcalFullOrganizerInput string
	//go:embed test_data/events/test_event_invalid_start.ical
	testIcalInvalidStartInput string
	//go:embed test_data/events/test_event_invalid_end.ical
	testIcalInvalidEndInput string
	//go:embed test_data/events/test_event_content_after_end_block.ical
	testIcalContentAfterEndBlockInput string
	//go:embed test_data/events/test_event_duplicate_uid.ical
	testIcalDuplicateUIDInput string
	//go:embed test_data/events/test_event_duplicate_sequence.ical
	testIcalDuplicateSequenceInput string
	//go:embed test_data/events/test_event_duplicate_status.ical
	testIcalDuplicateStatusInput string
	//go:embed test_data/events/test_event_duplicate_organizer.ical
	testIcalDuplicateOrganizerInput string
	//go:embed test_data/events/test_event_both_duration_and_end.ical
	testIcalBothDurationAndEndInput string
	//go:embed test_data/events/test_event_both_duration_and_end_duration_first.ical
	testIcalBothDurationAndEndDurationFirstInput string
	//go:embed test_data/events/invalid_test_event_with_bad_rrule.ical
	testIcalInvalidRRuleInput string
	//go:embed test_data/events/test_event_missing_colon.ical
	testIcalMissingColonInput string
	//go:embed test_data/events/test_event_missing_uid.ical
	testIcalMissingUIDInput string
	//go:embed test_data/events/test_event_missing_dtstart.ical
	testIcalMissingDTStartInput string
	//go:embed test_data/events/test_event_with_alarm.ical
	testEventWithAlarmInput string
	//go:embed test_data/events/test_event_alarm_missing_action.ical
	testEventAlarmMissingActionInput string
	//go:embed test_data/events/test_event_alarm_missing_description_display.ical
	testEventAlarmMissingDescriptionDisplayInput string
	//go:embed test_data/events/test_event_alarm_missing_attendee_email.ical
	testEventAlarmMissingAttendeeEmailInput string
	//go:embed test_data/events/test_event_alarm_missing_attach_audio.ical
	testEventAlarmMissingAttachAudioInput string
	//go:embed test_data/events/test_event_invalid_alarm_multiple_description.ical
	testEventAlarmInvalidMultipleDescriptionInput string
	//go:embed test_data/events/valid_test_event_with_rrule.ical
	testEventWithRRuleInput string
	//go:embed test_data/events/valid_test_event_with_attachment.ical
	testEventWithAttachmentInput string
)

func TestValidEvent(t *testing.T) {
	testCases := []struct {
		name             string
		input            string
		expectedCalendar *model.Calendar
	}{
		{
			name:  "Valid organizer with all parameters set",
			input: testIcalFullOrganizerInput,
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
						Description: "Event Description",
						Location:    "555 Fake Street",
						Organizer: &model.Organizer{
							CommonName: "JohnSmith",
							Directory:  &url.URL{Scheme: "ldap", Host: "example.com:6666", Path: "/o=DC Associates,c=US", RawQuery: "??(cn=John%20Smith)"},
							CalAddress: &url.URL{Scheme: "mailto", Opaque: "jsmith@example.com"},
							Language:   "en-us",
							SentBy:     &url.URL{Scheme: "mailto", Opaque: "mailtojsmith@example.com"},
							OtherParams: map[string]string{
								"MISCFIELD":  "TEST",
								"MISCFIELD2": "TEST2",
							},
						},
						Status:       model.EventStatusConfirmed,
						Sequence:     1,
						Comment:      []string{"I Am", "A Comment"},
						Categories:   []string{"first", "second", "third"},
						Geo:          []float64{37.386013, -122.082932},
						Transp:       model.EventTranspOpaque,
						Contacts:     []string{"Jim Dolittle, ABC Industries, +1-919-555-1234"},
						LastModified: time.Date(2021, time.January, 1, 0, 0, 0, 0, time.UTC),
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
			},
		},
		{
			name:  "Valid VEVENT with ATTACH",
			input: testEventWithAttachmentInput,
			expectedCalendar: &model.Calendar{
				ProdID:  "-//Event//Event Calendar//EN",
				Version: "2.0",
				Events: []model.Event{
					{
						UID:         "13235@example.com",
						DTStamp:     time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC),
						Start:       time.Date(2025, time.September, 28, 18, 30, 0, 0, time.UTC),
						End:         time.Date(2025, time.September, 28, 20, 30, 0, 0, time.UTC),
						Summary:     "Event with attachment",
						Description: "Event Description",
						Attach: []model.Attachment{
							{
								FormatType: "application/pdf",
								Value:      "URI",
								URI:        &url.URL{Scheme: "https", Host: "example.com", Path: "/files/report.pdf"},
							},
						},
					},
				},
			},
		},
		{
			name:  "Valid VEVENT with VALARM",
			input: testEventWithAlarmInput,
			expectedCalendar: &model.Calendar{
				ProdID:  "-//Event//Event Calendar//EN",
				Version: "2.0",
				Events: []model.Event{
					{
						UID:         "13235@example.com",
						DTStamp:     time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC),
						Start:       time.Date(2025, time.September, 28, 18, 30, 0, 0, time.UTC),
						End:         time.Date(2025, time.September, 28, 20, 30, 0, 0, time.UTC),
						Summary:     "Event with Alarm",
						Description: "Event Description",
						Alarms: []model.Alarm{
							{
								Action:      model.AlarmActionDisplay,
								Trigger:     "-PT15M",
								Description: "Reminder: Event starting in 15 minutes",
								Repeat:      2,
								Duration:    5 * time.Minute,
							},
							{
								Action:      model.AlarmActionEmail,
								Trigger:     "-PT1H",
								Description: "Email reminder for upcoming event",
								Summary:     "Event Reminder",
								Attendees:   []url.URL{{Scheme: "mailto", Opaque: "user@example.com"}},
							},
						},
					},
				},
			},
		},
		{
			name:  "Valid VEVENT with RRULE",
			input: testEventWithRRuleInput,
			expectedCalendar: &model.Calendar{
				ProdID:  "-//Event//Event Calendar//EN",
				Version: "2.0",
				Events: []model.Event{
					{
						UID:     "13235@example.com",
						DTStamp: time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC),
						Start:   time.Date(2025, time.September, 28, 18, 30, 0, 0, time.UTC),
						End:     time.Date(2025, time.September, 28, 20, 30, 0, 0, time.UTC),
						RRule: &rrule.RRule{
							Frequency: rrule.FrequencyDaily,
							Interval:  1,
							Count:     new(10),
						},
						Summary:     "Event with recurrence rule",
						Description: "Event Description",
					},
				},
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			calendar, err := ical.FromString(tc.input)
			require.NoError(t, err)
			assert.Equal(t, *tc.expectedCalendar, *calendar)
		})
	}
}

func TestInvalidEvent(t *testing.T) {
	testCases := []struct {
		name        string
		input       string
		expectedErr error
	}{
		{
			name:        "Invalid organizer",
			input:       testIcalInvalidOrganizerInput,
			expectedErr: icalerr.ErrInvalidOrganizer,
		},
		{
			name:        "Invalid start date",
			input:       testIcalInvalidStartInput,
			expectedErr: icalerr.ErrParseErrorInComponent,
		},
		{
			name:        "Invalid end date",
			input:       testIcalInvalidEndInput,
			expectedErr: icalerr.ErrParseErrorInComponent,
		},
		{
			name:        "Content after END:VCALENDAR",
			input:       testIcalContentAfterEndBlockInput,
			expectedErr: icalerr.ErrContentAfterEndBlock,
		},
		{
			name:        "Duplicate UID",
			input:       testIcalDuplicateUIDInput,
			expectedErr: icalerr.ErrDuplicatePropertyInComponent,
		},
		{
			name:        "Duplicate sequence",
			input:       testIcalDuplicateSequenceInput,
			expectedErr: icalerr.ErrDuplicatePropertyInComponent,
		},
		{
			name:        "Duplicate STATUS",
			input:       testIcalDuplicateStatusInput,
			expectedErr: icalerr.ErrDuplicatePropertyInComponent,
		},
		{
			name:        "Duplicate ORGANIZER",
			input:       testIcalDuplicateOrganizerInput,
			expectedErr: icalerr.ErrDuplicatePropertyInComponent,
		},
		{
			name:        "Both duration and end date are specified, DTEND first",
			input:       testIcalBothDurationAndEndInput,
			expectedErr: icalerr.ErrInvalidDurationPropertyDtend,
		},
		{
			name:        "Both duration and end date are specified, DURATION first",
			input:       testIcalBothDurationAndEndDurationFirstInput,
			expectedErr: icalerr.ErrInvalidDurationPropertyDtend,
		},
		{
			name:        "Missing colon in event property line",
			input:       testIcalMissingColonInput,
			expectedErr: icalerr.ErrInvalidPropertyLine,
		},
		{
			name:        "Missing UID",
			input:       testIcalMissingUIDInput,
			expectedErr: icalerr.ErrMissingEventUIDProperty,
		},
		{
			name:        "Missing DTSTART",
			input:       testIcalMissingDTStartInput,
			expectedErr: icalerr.ErrMissingEventDTStartProperty,
		},
		{
			name:        "VALARM missing ACTION",
			input:       testEventAlarmMissingActionInput,
			expectedErr: icalerr.ErrMissingAlarmActionProperty,
		},
		{
			name:        "VALARM DISPLAY missing DESCRIPTION",
			input:       testEventAlarmMissingDescriptionDisplayInput,
			expectedErr: icalerr.ErrMissingAlarmDescriptionForDisplay,
		},
		{
			name:        "VALARM EMAIL missing ATTENDEE",
			input:       testEventAlarmMissingAttendeeEmailInput,
			expectedErr: icalerr.ErrMissingAlarmAttendeesForEmail,
		},
		{
			name:        "VALARM AUDIO missing ATTACH",
			input:       testEventAlarmMissingAttachAudioInput,
			expectedErr: icalerr.ErrMissingAlarmAttachForAudio,
		},
		{
			name:        "VALARM multiple DESCRIPTION",
			input:       testEventAlarmInvalidMultipleDescriptionInput,
			expectedErr: icalerr.ErrDuplicatePropertyInComponent,
		},
		{
			name:        "Invalid RRULE",
			input:       testIcalInvalidRRuleInput,
			expectedErr: icalerr.ErrInvalidRRule,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			calendar, err := ical.FromString(tc.input)
			require.Nil(t, calendar)
			require.ErrorIs(t, err, tc.expectedErr)
		})
	}
}
