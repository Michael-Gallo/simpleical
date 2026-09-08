package benchmarks

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	"github.com/apognu/gocal"
	golangical "github.com/arran4/golang-ical"
	emersionical "github.com/emersion/go-ical"
	"github.com/michael-gallo/simpleical/ical"
)

const (
	// An extremely minimal ical file, with a single event with only required properties
	simpleFileName = "./test_simple.ical"
	// VEVENT-only comparative fixtures: no RRULE and no non-VEVENT components
	veventRichFileName      = "./test_vevent_rich.ical"
	veventsMultipleFileName = "./test_vevents_multiple.ical"
	singleFileName          = "./test_event.ical"
	multipleFileName        = "./test_multiple_events.ical"
	complexFileName         = "./test_complex.ical"
	// A stream of three sequential VCALENDAR objects, parsed with ical.Read
	multipleCalendarsFileName = "./test_multiple_calendars.ical"
)

func BenchmarkAllScenarios(b *testing.B) {
	testCases := []struct {
		fileName string
		testName string
	}{
		{simpleFileName, "Simple Event"},
		{singleFileName, "Single Event"},
		{multipleFileName, "Multiple Events"},
		{complexFileName, "Complex Calendar"},
	}
	for _, testCase := range testCases {
		b.Run(testCase.testName, func(b *testing.B) {
			benchmarkFile(b, testCase.fileName, testCase.testName)
		})
	}
}

func benchmarkFile(b *testing.B, fileName string, description string) {
	fileContent, err := os.ReadFile(fileName)
	if err != nil {
		b.Fatalf("Failed to read file %s: %v", fileName, err)
	}

	var reader bytes.Reader
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		reader.Reset(fileContent)
		cal, err := ical.ReadSingle(&reader)
		if err != nil {
			b.Fatalf("Failed to parse %s: %v", description, err)
		}

		// Basic validation to ensure parsing worked
		if cal == nil {
			b.Fatal("Calendar is nil")
		}

		// Prevent optimization
		_ = cal
	}
}

// BenchmarkMultipleCalendars measures parsing a stream containing multiple
// sequential VCALENDAR objects via ical.Read.
func BenchmarkMultipleCalendars(b *testing.B) {
	fileContent, err := os.ReadFile(multipleCalendarsFileName)
	if err != nil {
		b.Fatalf("Failed to read file %s: %v", multipleCalendarsFileName, err)
	}

	var reader bytes.Reader
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		reader.Reset(fileContent)
		cals, err := ical.Read(&reader)
		if err != nil {
			b.Fatalf("Failed to parse multiple calendars: %v", err)
		}
		if len(cals) != 3 {
			b.Fatalf("Expected 3 calendars, got %d", len(cals))
		}
	}
}

// BenchmarkComparativeVEVENT compares parse-into-structs on VEVENT-only
// calendars with no RRULE. Gocal is included here because it only materializes
// VEVENT components and would expand recurrences if an RRULE were present.
func BenchmarkComparativeVEVENT(b *testing.B) {
	testCases := []struct {
		fileName   string
		testName   string
		eventCount int
	}{
		{simpleFileName, "Simple Event", 1},
		{veventRichFileName, "Rich Event", 1},
		{veventsMultipleFileName, "Multiple Events", 2},
	}

	for _, testCase := range testCases {
		benchmarkVEVENTComparison(b, testCase.fileName, testCase.testName, testCase.eventCount)
	}
}

// BenchmarkComparativeCalendar compares parse-into-structs on full calendar
// objects (VTIMEZONE, RRULE, VTODO, VALARM, VJOURNAL). Gocal is omitted: it
// skips non-VEVENT components and expands RRULEs during Parse.
func BenchmarkComparativeCalendar(b *testing.B) {
	testCases := []struct {
		fileName string
		testName string
	}{
		{singleFileName, "Single Event"},
		{multipleFileName, "Multiple Events"},
		{complexFileName, "Complex Calendar"},
	}

	for _, testCase := range testCases {
		benchmarkCalendarComparison(b, testCase.fileName, testCase.testName)
	}
}

func benchmarkVEVENTComparison(b *testing.B, fileName string, testName string, eventCount int) {
	fileContent, err := os.ReadFile(fileName)
	if err != nil {
		b.Fatalf("Invalid File: %v", err)
	}
	var reader bytes.Reader

	b.Run(fmt.Sprintf("%s - SimpleIcal", testName), func(b *testing.B) {
		for b.Loop() {
			reader.Reset(fileContent)
			_, err := ical.ReadSingle(&reader)
			if err != nil {
				b.Fatalf("Failed to parse %s: %v", testName, err)
			}
		}
	})

	b.Run(fmt.Sprintf("%s - Gocal", testName), func(b *testing.B) {
		for b.Loop() {
			reader.Reset(fileContent)
			c := gocal.NewParser(&reader)
			c.SkipBounds = true // Parse all events regardless of date
			err := c.Parse()
			if err != nil {
				b.Fatalf("Failed to parse %s: %v", testName, err)
			}
			if got := len(c.Events); got != eventCount {
				b.Fatalf("gocal parsed %d events, want %d (RRULE expansion?)", got, eventCount)
			}
		}
	})

	b.Run(fmt.Sprintf("%s - GolangIcal", testName), func(b *testing.B) {
		for b.Loop() {
			reader.Reset(fileContent)
			_, err := golangical.ParseCalendar(&reader)
			if err != nil {
				b.Fatalf("Failed to parse %s: %v", testName, err)
			}
		}
	})

	b.Run(fmt.Sprintf("%s - Emersion", testName), func(b *testing.B) {
		for b.Loop() {
			reader.Reset(fileContent)
			_, err := emersionical.NewDecoder(&reader).Decode()
			if err != nil {
				b.Fatalf("Failed to parse %s: %v", testName, err)
			}
		}
	})
}

func benchmarkCalendarComparison(b *testing.B, fileName string, testName string) {
	fileContent, err := os.ReadFile(fileName)
	if err != nil {
		b.Fatalf("Invalid File: %v", err)
	}
	var reader bytes.Reader

	b.Run(fmt.Sprintf("%s - SimpleIcal", testName), func(b *testing.B) {
		for b.Loop() {
			reader.Reset(fileContent)
			_, err := ical.ReadSingle(&reader)
			if err != nil {
				b.Fatalf("Failed to parse %s: %v", testName, err)
			}
		}
	})

	b.Run(fmt.Sprintf("%s - GolangIcal", testName), func(b *testing.B) {
		for b.Loop() {
			reader.Reset(fileContent)
			_, err := golangical.ParseCalendar(&reader)
			if err != nil {
				b.Fatalf("Failed to parse %s: %v", testName, err)
			}
		}
	})

	b.Run(fmt.Sprintf("%s - Emersion", testName), func(b *testing.B) {
		for b.Loop() {
			reader.Reset(fileContent)
			_, err := emersionical.NewDecoder(&reader).Decode()
			if err != nil {
				b.Fatalf("Failed to parse %s: %v", testName, err)
			}
		}
	})
}
