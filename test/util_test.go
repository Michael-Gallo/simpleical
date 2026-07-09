package test

import (
	"os"
	"strings"
	"testing"

	"github.com/michael-gallo/simpleical/ical"
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
