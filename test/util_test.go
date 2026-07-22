package test

import (
	"os"
	"strings"
	"testing"
	"time"

	"github.com/michael-gallo/simpleical/ical"
	"github.com/michael-gallo/simpleical/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// test to ensure that we get the same result from a file reader as we do from a string reader
func TestReadSingleFromFile(t *testing.T) {
	file, err := os.Open("test_data/calendar/valid_calendar.ical")
	require.NoError(t, err)
	defer file.Close()

	calendarFromFile, err := ical.ReadSingle(file)
	require.NoError(t, err)
	calendarFromString, err := ical.ReadSingle(strings.NewReader(testValidCalendarInput))
	require.NoError(t, err)
	assert.Equal(t, *calendarFromFile, *calendarFromString)
}

// test to ensure that Read gets the same result from a file reader as it does from a string reader
func TestReadFromFile(t *testing.T) {
	file, err := os.Open("test_data/calendar/valid_multiple_calendars.ical")
	require.NoError(t, err)
	defer file.Close()

	calendarsFromFile, err := ical.Read(file)
	require.NoError(t, err)
	calendarsFromString, err := ical.Read(strings.NewReader(testMultipleCalendarsInput))
	require.NoError(t, err)
	assert.Equal(t, calendarsFromFile, calendarsFromString)
}

func text(s string) model.TextValue {
	return model.TextValue{Value: s}
}

func texts(ss ...string) []model.TextValue {
	out := make([]model.TextValue, len(ss))
	for i, s := range ss {
		out[i] = model.TextValue{Value: s}
	}
	return out
}

func related(ss ...string) []model.RelatedToValue {
	out := make([]model.RelatedToValue, len(ss))
	for i, s := range ss {
		out[i] = model.RelatedToValue{Value: s}
	}
	return out
}

func utcDT(year int, month time.Month, day, hour, min, sec int) model.DateTime {
	return model.NewUTCDateTime(time.Date(year, month, day, hour, min, sec, 0, time.UTC))
}

func floatDT(year int, month time.Month, day, hour, min, sec int) model.DateTime {
	return model.NewFloatingDateTime(time.Date(year, month, day, hour, min, sec, 0, time.UTC))
}

func dateDT(year int, month time.Month, day int) model.DateTime {
	return model.NewDate(time.Date(year, month, day, 0, 0, 0, 0, time.UTC))
}

func rdateUTC(year int, month time.Month, day, hour, min, sec int) model.RecurrenceDate {
	dt := utcDT(year, month, day, hour, min, sec)
	return model.RecurrenceDate{DateTime: &dt}
}

func triggerStart(d time.Duration) model.Trigger {
	return model.Trigger{Duration: &d, Related: model.TriggerRelatedStart}
}
