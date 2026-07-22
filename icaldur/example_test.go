package icaldur_test

import (
	"fmt"
	"time"

	"github.com/michael-gallo/simpleical/icaldur"
)

func ExampleParseICalDuration() {
	duration, err := icaldur.ParseICalDuration("P15DT5H0M20S")
	if err != nil {
		panic(err)
	}
	fmt.Println(duration.String())
	// Output: 365h0m20s
}

func ExampleParseIcalTime() {
	t, err := icaldur.ParseIcalTime("20250928T183000Z")
	if err != nil {
		panic(err)
	}
	fmt.Println(t.Year())
	fmt.Println(t.Month())
	fmt.Println(t.Day())
	fmt.Println(t.Hour())
	fmt.Println(t.Minute())
	fmt.Println(t.Second())
	// Output: 2025
	// September
	// 28
	// 18
	// 30
	// 0
}

func ExampleParseIcalTimeOrDate() {
	date, err := icaldur.ParseIcalTimeOrDate("20070628", "DATE")
	if err != nil {
		panic(err)
	}
	fmt.Println(date.Format("2006-01-02"))
	// Output: 2007-06-28
}

func ExampleParseTemporal() {
	temporal, err := icaldur.ParseTemporal("19971102", "DATE")
	if err != nil {
		panic(err)
	}
	fmt.Println(temporal.Form == icaldur.FormDate)
	fmt.Println(temporal.Time.Format("2006-01-02"))
	// Output:
	// true
	// 1997-11-02
}

func ExampleParseTemporalDateTime() {
	temporal, err := icaldur.ParseTemporalDateTime("19980118T230000")
	if err != nil {
		panic(err)
	}
	fmt.Println(temporal.Form == icaldur.FormFloating)
	// Output: true
}

func ExampleParseIcalUTCTime() {
	t, err := icaldur.ParseIcalUTCTime("19970714T170000Z")
	if err != nil {
		panic(err)
	}
	fmt.Println(t.Format(time.RFC3339))
	// Output: 1997-07-14T17:00:00Z
}

func ExampleParseIcalLocalTime() {
	t, err := icaldur.ParseIcalLocalTime("19980119T020000")
	if err != nil {
		panic(err)
	}
	fmt.Println(t.Hour())
	// Output: 2
}

func ExampleParseUTCOffset() {
	offset, err := icaldur.ParseUTCOffset("-0500")
	if err != nil {
		panic(err)
	}
	fmt.Println(offset)
	// Output: -0500
}
