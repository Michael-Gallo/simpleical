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
	"github.com/michael-gallo/simpleical/rrule"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (

	//go:embed test_data/events/test_event_invalid_organizer.ical
	testIcalInvalidOrganizerInput string
	//go:embed test_data/events/test_event_invalid_geo_latitude.ical
	testIcalInvalidGeoLatitudeInput string
	//go:embed test_data/events/test_event_invalid_geo_longitude.ical
	testIcalInvalidGeoLongitudeInput string
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
	//go:embed test_data/events/test_event_invalid_attach_value.ical
	testEventInvalidAttachValueInput string
	//go:embed test_data/events/valid_test_event_with_rrule_byweekno_bysetpos.ical
	testEventWithRRuleByWeekNoBySetPosInput string
	//go:embed test_data/events/test_event_attendee_params.ical
	testEventAttendeeParamsInput string
	//go:embed test_data/events/valid_test_event_with_new_props.ical
	testEventWithNewPropsInput string
	//go:embed test_data/events/valid_test_event_method_no_dtstart.ical
	testEventMethodNoDTStartInput string
	//go:embed test_data/events/valid_test_event_all_day_date.ical
	testEventAllDayDateInput string
	//go:embed test_data/events/test_event_duplicate_class.ical
	testEventDuplicateClassInput string
	//go:embed test_data/events/test_event_duplicate_created.ical
	testEventDuplicateCreatedInput string
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
						Geo:          &[2]float64{37.386013, -122.082932},
						Transp:       model.TranspOpaque,
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
								Value:      model.AttachValueURI,
								URI:        &url.URL{Scheme: "https", Host: "example.com", Path: "/files/report.pdf"},
							},
						},
					},
				},
			},
		},
		{
			name:  "Valid VEVENT with ATTENDEE parameters",
			input: testEventAttendeeParamsInput,
			expectedCalendar: &model.Calendar{
				ProdID:  "-//Event//Event Calendar//EN",
				Version: "2.0",
				Events: []model.Event{
					{
						UID:     "attendee-params@example.com",
						DTStamp: time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC),
						Start:   time.Date(2025, time.September, 28, 18, 30, 0, 0, time.UTC),
						End:     time.Date(2025, time.September, 28, 20, 30, 0, 0, time.UTC),
						Summary: "Event with Attendee Params",
						Attendees: []model.Attendee{
							{
								CommonName: "Jane Doe",
								CalAddress: &url.URL{Scheme: "mailto", Opaque: "jdoe@example.com"},
								Role:       "REQ-PARTICIPANT",
								PartStat:   "ACCEPTED",
								DelegatedFrom: []*url.URL{
									{Scheme: "mailto", Opaque: "bob@example.com"},
								},
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
								Attendees: []model.Attendee{
									{CalAddress: &url.URL{Scheme: "mailto", Opaque: "user@example.com"}},
								},
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
		{
			name:  "Valid VEVENT with RRULE BYWEEKNO and BYSETPOS lists",
			input: testEventWithRRuleByWeekNoBySetPosInput,
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
							Frequency: rrule.FrequencyYearly,
							Interval:  1,
							ByWeekNo:  []int8{1, -2, 53},
							ByDay:     []rrule.ByDay{{Weekday: rrule.WeekdayMonday, Interval: 1}},
						},
						Summary:     "Event with multi-value BYWEEKNO",
						Description: "Event Description",
					},
					{
						UID:     "13236@example.com",
						DTStamp: time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC),
						Start:   time.Date(2025, time.September, 28, 18, 30, 0, 0, time.UTC),
						End:     time.Date(2025, time.September, 28, 20, 30, 0, 0, time.UTC),
						RRule: &rrule.RRule{
							Frequency: rrule.FrequencyMonthly,
							Interval:  1,
							ByDay: []rrule.ByDay{
								{Weekday: rrule.WeekdayTuesday, Interval: 1},
								{Weekday: rrule.WeekdayWednesday, Interval: 1},
								{Weekday: rrule.WeekdayThursday, Interval: 1},
							},
							BySetPos: []int16{1, -1, 366},
						},
						Summary:     "Event with multi-value BYSETPOS",
						Description: "Event Description",
					},
				},
			},
		},
		{
			name:  "Valid VEVENT with previously missing properties",
			input: testEventWithNewPropsInput,
			expectedCalendar: &model.Calendar{
				ProdID:  "-//Event//Event Calendar//EN",
				Version: "2.0",
				Events: []model.Event{
					{
						UID:          "13235@example.com",
						DTStamp:      time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC),
						Start:        time.Date(2025, time.September, 28, 18, 30, 0, 0, time.UTC),
						End:          time.Date(2025, time.September, 28, 20, 30, 0, 0, time.UTC),
						Summary:      "Event with new properties",
						Class:        model.ClassPrivate,
						Created:      time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC),
						Priority:     5,
						URL:          "https://example.com/event/13235",
						RecurrenceID: time.Date(2025, time.September, 28, 18, 30, 0, 0, time.UTC),
						ExceptionDates: []time.Time{
							time.Date(2025, time.October, 5, 18, 30, 0, 0, time.UTC),
							time.Date(2025, time.October, 6, 18, 30, 0, 0, time.UTC),
						},
						Rdate: []time.Time{
							time.Date(2025, time.October, 12, 18, 30, 0, 0, time.UTC),
							time.Date(2025, time.October, 19, 18, 30, 0, 0, time.UTC),
						},
						RequestStatus: []string{"2.0;Success"},
						Related:       []string{"parent-event@example.com"},
						Resources:     []string{"projector", "whiteboard"},
					},
				},
			},
		},
		{
			name:  "Valid VEVENT without DTSTART when METHOD is present",
			input: testEventMethodNoDTStartInput,
			expectedCalendar: &model.Calendar{
				ProdID:  "-//Event//Event Calendar//EN",
				Version: "2.0",
				Method:  "REQUEST",
				Events: []model.Event{
					{
						UID:         "13235@example.com",
						DTStamp:     time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC),
						Summary:     "Event without DTSTART",
						Description: "METHOD is present so DTSTART is optional",
					},
				},
			},
		},
		{
			name:  "Valid all-day VEVENT with VALUE=DATE",
			input: testEventAllDayDateInput,
			expectedCalendar: &model.Calendar{
				ProdID:  "-//Event//Event Calendar//EN",
				Version: "2.0",
				Events: []model.Event{
					{
						UID:     "19970901T130000Z-123403@example.com",
						DTStamp: time.Date(1997, time.September, 1, 13, 0, 0, 0, time.UTC),
						Start:   time.Date(2007, time.June, 28, 0, 0, 0, 0, time.UTC),
						End:     time.Date(2007, time.July, 9, 0, 0, 0, 0, time.UTC),
						Summary: "Festival International de Jazz de Montreal",
						Transp:  model.TranspTransparent,
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
			name:        "VEVENT invalid GEO latitude",
			input:       testIcalInvalidGeoLatitudeInput,
			expectedErr: icalerr.ErrInvalidGeoPropertyLatitude,
		},
		{
			name:        "VEVENT invalid GEO longitude",
			input:       testIcalInvalidGeoLongitudeInput,
			expectedErr: icalerr.ErrInvalidGeoPropertyLongitude,
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
		{
			name:        "VEVENT with invalid ATTACH VALUE parameter",
			input:       testEventInvalidAttachValueInput,
			expectedErr: icalerr.ErrParseErrorInComponent,
		},
		{
			name:        "Duplicate CLASS",
			input:       testEventDuplicateClassInput,
			expectedErr: icalerr.ErrDuplicatePropertyInComponent,
		},
		{
			name:        "Duplicate CREATED",
			input:       testEventDuplicateCreatedInput,
			expectedErr: icalerr.ErrDuplicatePropertyInComponent,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			calendar, err := ical.ReadSingle(strings.NewReader(tc.input))
			require.Nil(t, calendar)
			require.ErrorIs(t, err, tc.expectedErr)
		})
	}
}
