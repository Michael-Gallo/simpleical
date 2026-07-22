package test

import (
	"strings"
	"testing"
	"time"

	"github.com/michael-gallo/simpleical/ical"
	"github.com/michael-gallo/simpleical/internal/icalerr"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTextUnescapeAndEscapedListSeparators(t *testing.T) {
	input := `BEGIN:VCALENDAR
VERSION:2.0
PRODID:-//Test//EN
BEGIN:VEVENT
UID:text@example.com
DTSTAMP:19970101T000000Z
DTSTART:19970101T000000Z
SUMMARY:Line1\nLine2
CATEGORIES:A\,B,C
END:VEVENT
END:VCALENDAR
`
	cal, err := ical.ReadSingle(strings.NewReader(input))
	require.NoError(t, err)
	assert.Equal(t, "Line1\nLine2", cal.Events[0].Summary.Value)
	assert.Equal(t, []string{"A,B", "C"}, cal.Events[0].Categories)
}

func TestNestedTopLevelComponentRejected(t *testing.T) {
	input := `BEGIN:VCALENDAR
VERSION:2.0
PRODID:-//Test//EN
BEGIN:VEVENT
UID:nested@example.com
DTSTAMP:19970101T000000Z
DTSTART:19970101T000000Z
BEGIN:VTODO
UID:todo@example.com
DTSTAMP:19970101T000000Z
END:VTODO
END:VEVENT
END:VCALENDAR
`
	_, err := ical.ReadSingle(strings.NewReader(input))
	require.Error(t, err)
	assert.ErrorIs(t, err, icalerr.ErrComponentNotAllowedHere)
}

func TestLeapSecondAccepted(t *testing.T) {
	input := `BEGIN:VCALENDAR
VERSION:2.0
PRODID:-//Test//EN
BEGIN:VEVENT
UID:leap@example.com
DTSTAMP:19970630T235960Z
DTSTART:19970630T235960Z
SUMMARY:Leap second
END:VEVENT
END:VCALENDAR
`
	cal, err := ical.ReadSingle(strings.NewReader(input))
	require.NoError(t, err)
	assert.Equal(t, 59, cal.Events[0].Start.Time.Second())
	assert.True(t, cal.Events[0].Start.IsUTC())
}

func TestParamNameCaseInsensitive(t *testing.T) {
	input := `BEGIN:VCALENDAR
VERSION:2.0
PRODID:-//Test//EN
BEGIN:VEVENT
UID:param@example.com
DTSTAMP:19970101T000000Z
DTSTART:19970101T000000Z
ATTENDEE;cn=Bob:mailto:bob@example.com
END:VEVENT
END:VCALENDAR
`
	cal, err := ical.ReadSingle(strings.NewReader(input))
	require.NoError(t, err)
	require.Len(t, cal.Events[0].Attendees, 1)
	assert.Equal(t, "Bob", cal.Events[0].Attendees[0].CommonName)
}

func TestTrailingValueWhitespacePreserved(t *testing.T) {
	input := "BEGIN:VCALENDAR\r\n" +
		"VERSION:2.0\r\n" +
		"PRODID:-//Test//EN\r\n" +
		"BEGIN:VEVENT\r\n" +
		"UID:space@example.com\r\n" +
		"DTSTAMP:19970101T000000Z\r\n" +
		"DTSTART:19970101T000000Z\r\n" +
		"SUMMARY:Hello   \r\n" +
		"END:VEVENT\r\n" +
		"END:VCALENDAR\r\n"
	cal, err := ical.ReadSingle(strings.NewReader(input))
	require.NoError(t, err)
	assert.Equal(t, "Hello   ", cal.Events[0].Summary.Value)
}

func TestFreeBusyDurationPeriodAndFBType(t *testing.T) {
	input := `BEGIN:VCALENDAR
VERSION:2.0
PRODID:-//Test//EN
BEGIN:VFREEBUSY
UID:fb@example.com
DTSTAMP:19970101T000000Z
FREEBUSY;FBTYPE=BUSY-UNAVAILABLE:19970308T160000Z/PT8H30M,19970309T160000Z/19970309T170000Z
END:VFREEBUSY
END:VCALENDAR
`
	cal, err := ical.ReadSingle(strings.NewReader(input))
	require.NoError(t, err)
	require.Len(t, cal.FreeBusys[0].FreeBusy, 2)
	assert.Equal(t, 8*time.Hour+30*time.Minute, cal.FreeBusys[0].FreeBusy[0].Duration)
	assert.True(t, cal.FreeBusys[0].FreeBusy[1].End.IsUTC())
}
