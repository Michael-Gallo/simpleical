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

	//go:embed test_data/todos/test_todo.ical
	testTodoInput string
	//go:embed test_data/todos/test_todo_missing_uid.ical
	testTodoMissingUIDInput string
	//go:embed test_data/todos/test_todo_both_due_and_duration.ical
	testTodoBothDueAndDurationInput string
	//go:embed test_data/todos/test_todo_duplicate_uid.ical
	testTodoDuplicateUIDInput string
	//go:embed test_data/todos/test_todo_invalid_geo.ical
	testTodoInvalidGeoInput string
	//go:embed test_data/todos/test_todo_invalid_geo_latitude.ical
	testTodoInvalidGeoLatitudeInput string
	//go:embed test_data/todos/test_todo_invalid_geo_longitude.ical
	testTodoInvalidGeoLongitudeInput string
	//go:embed test_data/todos/test_todo_with_rrule.ical
	testTodoWithRRuleInput string
	//go:embed test_data/todos/test_todo_duplicate_organizer.ical
	testTodoDuplicateOrganizerInput string
	//go:embed test_data/todos/test_todo_duplicate_status.ical
	testTodoDuplicateStatusInput string
	//go:embed test_data/todos/test_todo_duration_without_dtstart.ical
	testTodoDurationWithoutDTStartInput string
	//go:embed test_data/todos/test_todo_missing_dtstamp.ical
	testTodoMissingDTStampInput string
	//go:embed test_data/todos/test_todo_transp_not_allowed.ical
	testTodoTranspNotAllowedInput string
	//go:embed test_data/todos/valid_test_todo_all_day_date.ical
	testTodoAllDayDateInput string
	//go:embed test_data/todos/valid_test_todo_attach_exdate_rdate_alarm.ical
	testTodoAttachExdateRdateAlarmInput string
)

func TestValidTodo(t *testing.T) {
	testCases := []struct {
		name             string
		input            string
		expectedCalendar *model.Calendar
	}{
		{
			name:  "Valid VTODO",
			input: testTodoInput,
			expectedCalendar: &model.Calendar{
				ProdID:  "-//Test//Todo Calendar//EN",
				Version: "2.0",
				Todos: []model.Todo{
					{
						UID:             "todo123@example.com",
						DTStamp:         utcDT(2024, 1, 1, 0, 0, 0),
						Summary:         text("Complete project documentation"),
						Description:     text("Write comprehensive documentation for the new API"),
						Location:        text("Office"),
						Class:           model.ClassConfidential,
						Status:          model.TodoStatusInProcess,
						PercentComplete: 75,
						Created:         utcDT(2024, 1, 1, 0, 0, 0),
						LastModified:    utcDT(2024, 1, 15, 12, 0, 0),
						DTStart:         utcDT(2024, 1, 1, 9, 0, 0),
						Due:             utcDT(2024, 1, 30, 17, 0, 0),
						Organizer: &model.Organizer{
							CommonName: "Project Manager",
							CalAddress: &url.URL{Scheme: "mailto", Opaque: "pm@example.com"},
						},
						Attendees: []model.Attendee{
							{CalAddress: &url.URL{Scheme: "mailto", Opaque: "dev1@example.com"}},
							{CalAddress: &url.URL{Scheme: "mailto", Opaque: "dev2@example.com"}},
						},
						Contacts:   []string{"John Doe, Engineering Team, +1-555-0123"},
						Categories: []string{"work", "urgent", "project"},
						Comment:    texts("This is a critical task for the Q1 release"),
						Resources:  []string{"laptop", "meeting-room"},
						Geo:        &[2]float64{37.7749, -122.4194},
						URL:        "https://project.example.com/todo/123",
						IANAProp:   []model.ExtensionProperty{{Name: "PRIORITY", Value: "1"}},
					},
				},
			},
		},
		{
			name:  "Valid VTODO with RRULE",
			input: testTodoWithRRuleInput,
			expectedCalendar: &model.Calendar{
				ProdID:  "-//Test//Todo Calendar//EN",
				Version: "2.0",
				Todos: []model.Todo{
					{
						UID:             "todo123@example.com",
						DTStamp:         utcDT(2024, 1, 1, 0, 0, 0),
						Summary:         text("Complete project documentation"),
						Description:     text("Write comprehensive documentation for the new API"),
						Location:        text("Office"),
						Class:           model.ClassConfidential,
						Status:          model.TodoStatusInProcess,
						PercentComplete: 75,
						Created:         utcDT(2024, 1, 1, 0, 0, 0),
						LastModified:    utcDT(2024, 1, 15, 12, 0, 0),
						DTStart:         utcDT(2024, 1, 1, 9, 0, 0),
						Due:             utcDT(2024, 1, 30, 17, 0, 0),
						Organizer: &model.Organizer{
							CommonName: "Project Manager",
							CalAddress: &url.URL{Scheme: "mailto", Opaque: "pm@example.com"},
						},
						Attendees: []model.Attendee{
							{CalAddress: &url.URL{Scheme: "mailto", Opaque: "dev1@example.com"}},
							{CalAddress: &url.URL{Scheme: "mailto", Opaque: "dev2@example.com"}},
						},
						Contacts:   []string{"John Doe, Engineering Team, +1-555-0123"},
						Categories: []string{"work", "urgent", "project"},
						Comment:    texts("This is a critical task for the Q1 release"),
						Resources:  []string{"laptop", "meeting-room"},
						Geo:        &[2]float64{37.7749, -122.4194},
						URL:        "https://project.example.com/todo/123",
						IANAProp:   []model.ExtensionProperty{{Name: "PRIORITY", Value: "1"}},
						RRule: &rrule.RRule{
							Frequency: rrule.FrequencyDaily,
							Interval:  1,
							Count:     new(10),
						},
					},
				},
			},
		},
		{
			name:  "Valid all-day VTODO with VALUE=DATE",
			input: testTodoAllDayDateInput,
			expectedCalendar: &model.Calendar{
				ProdID:  "-//Test//Todo Calendar//EN",
				Version: "2.0",
				Todos: []model.Todo{
					{
						UID:     "todo-date@example.com",
						DTStamp: utcDT(1997, 9, 1, 13, 0, 0),
						Summary: text("All-day todo"),
						DTStart: dateDT(6, 28),
						Due:     dateDT(7, 9),
					},
				},
			},
		},
		{
			name:  "Valid VTODO with ATTACH, COMPLETED, SEQUENCE, RECURRENCE-ID, EXDATE, RDATE, RELATED-TO, REQUEST-STATUS, VALARM",
			input: testTodoAttachExdateRdateAlarmInput,
			expectedCalendar: &model.Calendar{
				ProdID:  "-//Test//Todo Calendar//EN",
				Version: "2.0",
				Todos: []model.Todo{
					{
						UID:          "todo-attach-exdate-rdate-alarm@example.com",
						DTStamp:      utcDT(2024, 1, 1, 0, 0, 0),
						DTStart:      utcDT(2024, 1, 1, 9, 0, 0),
						Due:          utcDT(2024, 1, 30, 17, 0, 0),
						Summary:      text("Todo with ATTACH EXDATE RDATE VALARM"),
						Completed:    utcDT(2024, 1, 20, 15, 0, 0),
						Sequence:     3,
						RecurrenceID: utcDT(2024, 1, 1, 9, 0, 0),
						Attach: []model.Attachment{
							{
								FormatType: "application/pdf",
								Value:      model.AttachValueURI,
								URI:        &url.URL{Scheme: "https", Host: "example.com", Path: "/files/todo.pdf"},
							},
						},
						ExceptionDates: []model.DateTime{
							utcDT(2024, 1, 8, 9, 0, 0),
							utcDT(2024, 1, 15, 9, 0, 0),
						},
						Rdate: []model.RecurrenceDate{
							rdateUTC(2024, 1, 22, 9, 0),
							rdateUTC(2024, 1, 29, 9, 0),
						},
						RequestStatus: []string{"2.0;Success"},
						Related:       related("parent-todo@example.com"),
						Alarms: []model.Alarm{
							{
								Action:      model.AlarmActionDisplay,
								Trigger:     triggerStart(-15 * time.Minute),
								Description: "Reminder: todo due soon",
								Repeat:      2,
								Duration:    5 * time.Minute,
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

func TestInvalidTodo(t *testing.T) {
	testCases := []struct {
		name        string
		input       string
		expectedErr error
	}{
		{
			name:        "VTODO missing UID",
			input:       testTodoMissingUIDInput,
			expectedErr: icalerr.ErrMissingTodoUIDProperty,
		},
		{
			name:        "VTODO both DUE and DURATION",
			input:       testTodoBothDueAndDurationInput,
			expectedErr: icalerr.ErrInvalidDurationPropertyDue,
		},
		{
			name:        "VTODO invalid GEO",
			input:       testTodoInvalidGeoInput,
			expectedErr: icalerr.ErrInvalidGeoProperty,
		},
		{
			name:        "VTODO invalid GEO latitude",
			input:       testTodoInvalidGeoLatitudeInput,
			expectedErr: icalerr.ErrInvalidGeoPropertyLatitude,
		},
		{
			name:        "VTODO invalid GEO longitude",
			input:       testTodoInvalidGeoLongitudeInput,
			expectedErr: icalerr.ErrInvalidGeoPropertyLongitude,
		},
		{
			name:        "VTODO duplicate UID",
			input:       testTodoDuplicateUIDInput,
			expectedErr: icalerr.ErrDuplicatePropertyInComponent,
		},
		{
			name:        "VTODO Organizer set twice",
			input:       testTodoDuplicateOrganizerInput,
			expectedErr: icalerr.ErrDuplicatePropertyInComponent,
		},
		{
			name:        "VTODO Status set twice",
			input:       testTodoDuplicateStatusInput,
			expectedErr: icalerr.ErrDuplicatePropertyInComponent,
		},
		{
			name:        "VTODO DURATION without DTSTART",
			input:       testTodoDurationWithoutDTStartInput,
			expectedErr: icalerr.ErrDurationRequiresDTStart,
		},
		{
			name:        "VTODO missing DTSTAMP",
			input:       testTodoMissingDTStampInput,
			expectedErr: icalerr.ErrMissingTodoDTStampProperty,
		},
		{
			name:        "VTODO TRANSP not allowed",
			input:       testTodoTranspNotAllowedInput,
			expectedErr: icalerr.ErrTodoTranspNotAllowed,
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
