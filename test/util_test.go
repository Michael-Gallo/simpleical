package test

import (
	"testing"

	"github.com/michael-gallo/simpleical/ical"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// test to ensure that we get the same result from IcalFromFile as we do from IcalReader
func TestIcalFromFile(t *testing.T) {
	calendarFromFile, err := ical.FromFileName("test_data/calendar/valid_calendar.ical")
	require.NoError(t, err)
	calendarFromString, err := ical.FromString(testValidCalendarInput)
	require.NoError(t, err)
	assert.Equal(t, *calendarFromFile, *calendarFromString)
}
