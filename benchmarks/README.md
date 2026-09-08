# iCalendar Parser Benchmarks

Comparative parse-into-structs benchmarks against other Go iCalendar libraries.

These benches measure turning ICS bytes into each library’s in-memory structs. They do not expand recurrences or serialize.

## Libraries

| Library | When it is compared | Parse API used |
|---------|---------------------|----------------|
| simpleical | Every table | `ical.ReadSingle` |
| [gocal](https://github.com/apognu/gocal) | VEVENT-only files with no `RRULE` | `NewParser` + `SkipBounds` + `Parse` |
| [golang-ical](https://github.com/arran4/golang-ical) | VEVENT and full-calendar tables | `ParseCalendar` |
| [emersion/go-ical](https://github.com/emersion/go-ical) | VEVENT and full-calendar tables | `NewDecoder.Decode` |

gocal is omitted from the full-calendar benches: it ignores non-`VEVENT` components and expands `RRULE`s during `Parse()`.

## Setup

```bash
go mod tidy
```

## Running Benchmarks

From this directory, or via the repo `Makefile`.

```bash
# simpleical-only scenarios
go test -bench=BenchmarkAllScenarios -benchmem

# comparative: VEVENT-only (includes gocal) and full calendar (no gocal)
go test -bench=BenchmarkComparative -benchmem

# Makefile: ten runs plus benchstat
make bench-comparative
```

### CPU / memory profiling

```bash
go test -bench=BenchmarkAllScenarios -benchmem -cpuprofile=cpu.prof -memprofile=mem.prof
go tool pprof cpu.prof
go tool pprof mem.prof
```

## What is measured

- **Parsing speed**: ns/op (reported as sec/op in README tables)
- **Memory**: B/op
- **Allocations**: allocs/op

## Test data

Fixtures live in this directory:

- `test_simple.ical`, `test_vevent_rich.ical`, `test_vevents_multiple.ical` — VEVENT-only, no `RRULE` (`BenchmarkComparativeVEVENT`)
- `test_event.ical`, `test_multiple_events.ical`, `test_complex.ical` — full calendars (`BenchmarkComparativeCalendar` and simpleical-only `BenchmarkAllScenarios`)
- `test_multiple_calendars.ical` — sequential `VCALENDAR` stream (`BenchmarkMultipleCalendars`, simpleical only)

To add a VEVENT comparative case, add a `.ical` file with only `VEVENT` components and no `RRULE`, then extend the `testCases` slice in `BenchmarkComparativeVEVENT`. The gocal sub-bench asserts `len(Events)` equals the number of `VEVENT`s so accidental recurrence expansion fails the run.
