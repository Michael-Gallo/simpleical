package ical_test

import (
	"fmt"
	"os"
	"strings"

	"github.com/michael-gallo/simpleical/ical"
)

const testIcalString string = `BEGIN:VCALENDAR
VERSION:2.0
PRODID:-//Event//Event Calendar//EN
CALSCALE:GREGORIAN
METHOD:REQUEST
BEGIN:VTIMEZONE
TZID:America/Detroit
BEGIN:STANDARD
DTSTART:19700101T000000
TZOFFSETFROM:+0000
TZOFFSETTO:+0000
END:STANDARD
END:VTIMEZONE
BEGIN:VEVENT
UID:13235@example.com
DTSTAMP:19700101T000000Z
DTSTART:20250928T183000Z
DTEND:20250928T203000Z
SUMMARY:Event Summary
DESCRIPTION:Event Description
LOCATION:555 Fake Street
ORGANIZER;CN=Org:MAILTO:hello@world
STATUS:CONFIRMED
SEQUENCE:0
TRANSP:OPAQUE
END:VEVENT
END:VCALENDAR
`

func ExampleRead() {
	// A single iCalendar stream may contain multiple sequential VCALENDAR objects
	multiCalendarStream := `BEGIN:VCALENDAR
VERSION:2.0
PRODID:-//First//Calendar//EN
BEGIN:VEVENT
UID:first@example.com
DTSTAMP:19700101T000000Z
DTSTART:19700101T000000Z
SUMMARY:First Event
END:VEVENT
END:VCALENDAR
BEGIN:VCALENDAR
VERSION:2.0
PRODID:-//Second//Calendar//EN
BEGIN:VEVENT
UID:13235@example.com
DTSTAMP:19700101T000000Z
DTSTART:20250928T183000Z
SUMMARY:Event Summary
END:VEVENT
END:VCALENDAR
`
	reader := strings.NewReader(multiCalendarStream)
	calendars, err := ical.Read(reader)
	if err != nil {
		panic(err)
	}

	fmt.Println(len(calendars))
	fmt.Println(calendars[0].ProdID)
	fmt.Println(calendars[1].Events[0].Summary.Value)
	// Output:
	// 2
	// -//First//Calendar//EN
	// Event Summary
}

func ExampleReadSingle() {
	reader := strings.NewReader(testIcalString)
	calendar, err := ical.ReadSingle(reader)
	if err != nil {
		panic(err)
	}

	fmt.Println(calendar.ProdID)
	fmt.Println(calendar.TimeZones[0].TimeZoneID)
	fmt.Println(calendar.Events[0].Summary.Value)
	// Output:
	// -//Event//Event Calendar//EN
	// America/Detroit
	// Event Summary
}

// ExampleReadSingle_file demonstrates parsing a calendar from a file.
func ExampleReadSingle_file() {
	file, err := os.Open("../test/test_data/calendar/valid_calendar.ical")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	calendar, err := ical.ReadSingle(file)
	if err != nil {
		panic(err)
	}

	fmt.Println(calendar.ProdID)
	fmt.Println(calendar.CalScale)
	// Output:
	// -//Event//Event Calendar//EN
	// GREGORIAN
}
