package test

import (
	"testing"

	"github.com/michael-gallo/simpleical/ical"
	"github.com/stretchr/testify/assert"
)

// test to ensure that we get the same result from IcalFromFile as we do from IcalReader
func TestIcalFromFile(t *testing.T) {
	calendarFromFile, err := ical.FromFileName("test_data/calendar/valid_calendar.ical")
	assert.NoError(t, err)
	calendarFromString, err := ical.FromString(testValidCalendarInput)
	assert.NoError(t, err)
	assert.Equal(t, *calendarFromFile, *calendarFromString)
}
