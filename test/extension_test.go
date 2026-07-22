package test

import (
	_ "embed"
	"strings"
	"testing"
	"time"

	"github.com/michael-gallo/simpleical/ical"
	"github.com/michael-gallo/simpleical/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	//go:embed test_data/calendar/valid_calendar_with_extensions.ical
	testCalendarWithExtensionsInput string
	//go:embed test_data/timezones/valid_timezone_with_extensions.ical
	testTimezoneWithExtensionsInput string
)

func TestExtensionPropertiesAndUnknownComponents(t *testing.T) {
	calendar, err := ical.ReadSingle(strings.NewReader(testCalendarWithExtensionsInput))
	require.NoError(t, err)

	assert.Equal(t, []model.ExtensionProperty{
		{Name: "X-WR-CALNAME", Value: "Extension Calendar"},
		{Name: "X-WR-CALNAME", Value: "Duplicate Name"},
	}, calendar.XProp)
	assert.Equal(t, []model.ExtensionProperty{
		{Name: "DRESSCODE", Value: "CASUAL"},
	}, calendar.IANAProp)

	require.Len(t, calendar.Events, 1)
	event := calendar.Events[0]
	assert.Equal(t, "Lowercase Summary", event.Summary)
	assert.Equal(t, []model.ExtensionProperty{
		{
			Name:  "X-ABC-MMSUBJ",
			Value: "http://example.org/mysubj.au",
			Params: map[string]string{
				"VALUE":   "URI",
				"FMTTYPE": "audio/basic",
			},
		},
		{Name: "X-FOO", Value: "one"},
		{Name: "X-FOO", Value: "two"},
	}, event.XProp)
	assert.Equal(t, []model.ExtensionProperty{
		{Name: "DRESSCODE", Value: "FORMAL"},
	}, event.IANAProp)

	require.Len(t, event.Alarms, 1)
	assert.Equal(t, []model.ExtensionProperty{
		{Name: "X-ALARM-ID", Value: "42"},
	}, event.Alarms[0].XProp)

	require.Len(t, calendar.Todos, 1)
	assert.Equal(t, []model.ExtensionProperty{
		{Name: "X-TODO-FLAG", Value: "YES"},
	}, calendar.Todos[0].XProp)
	assert.Equal(t, []model.ExtensionProperty{
		{
			Name:   "NON-SMOKING",
			Value:  "TRUE",
			Params: map[string]string{"VALUE": "BOOLEAN"},
		},
	}, calendar.Todos[0].IANAProp)

	require.Len(t, calendar.Journals, 1)
	assert.Equal(t, []model.ExtensionProperty{
		{Name: "X-JOURNAL-NOTE", Value: "note"},
	}, calendar.Journals[0].XProp)
	assert.Equal(t, []model.ExtensionProperty{
		{Name: "COLOR", Value: "BLUE"},
	}, calendar.Journals[0].IANAProp)

	require.Len(t, calendar.FreeBusys, 1)
	assert.Equal(t, []model.ExtensionProperty{
		{Name: "X-FB-SOURCE", Value: "cache"},
	}, calendar.FreeBusys[0].XProp)
	assert.Equal(t, []model.ExtensionProperty{
		{Name: "BUSYTYPE", Value: "BUSY"},
	}, calendar.FreeBusys[0].IANAProp)

	// Unknown components are ignored, not stored.
	assert.Empty(t, calendar.TimeZones)
	assert.Equal(t, time.Date(2024, time.January, 1, 9, 0, 0, 0, time.UTC), event.Start)
}

func TestTimezoneExtensionProperties(t *testing.T) {
	calendar, err := ical.ReadSingle(strings.NewReader(testTimezoneWithExtensionsInput))
	require.NoError(t, err)

	require.Len(t, calendar.TimeZones, 1)
	tz := calendar.TimeZones[0]
	assert.Equal(t, "America/New_York", tz.TimeZoneID)
	assert.Equal(t, []model.ExtensionProperty{
		{Name: "X-LIC-LOCATION", Value: "America/New_York"},
	}, tz.XProp)
	assert.Equal(t, []model.ExtensionProperty{
		{Name: "TZUNTIL", Value: "20301231T000000Z"},
	}, tz.IANAProp)

	require.Len(t, tz.Standard, 1)
	assert.Equal(t, []model.ExtensionProperty{
		{Name: "X-STANDARD-NOTE", Value: "winter"},
	}, tz.Standard[0].XProp)
	assert.Equal(t, []model.ExtensionProperty{
		{Name: "OBSERVANCE", Value: "STANDARD"},
	}, tz.Standard[0].IANAProp)
}

func TestCaseInsensitiveBeginVCalendar(t *testing.T) {
	input := "begin:vcalendar\r\nVERSION:2.0\r\nPRODID:-//Test//EN\r\nEND:VCALENDAR\r\n"
	calendar, err := ical.ReadSingle(strings.NewReader(input))
	require.NoError(t, err)
	assert.Equal(t, "2.0", calendar.Version)
	assert.Equal(t, "-//Test//EN", calendar.ProdID)
}
