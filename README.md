# Simple-ical

A very much not ready ICAL parser for Golang intended to follow the official [ICAL 2.0 spec](https://datatracker.ietf.org/doc/html/rfc5545) as closely as is reasonable.

Focused on ease of use and good documentation, with frequent links to the spec.

[![Go Reference](https://pkg.go.dev/badge/github.com/michael-gallo/simpleical.svg)](https://pkg.go.dev/github.com/michael-gallo/simpleical)

## Documentation

Full API documentation is available on [pkg.go.dev](https://pkg.go.dev/github.com/michael-gallo/simpleical).

## License

This project is licensed under the Mozilla Public License 2.0. See the [LICENSE](LICENSE) file for details.


## Installation


```sh
go get github.com/michael-gallo/simpleical
```

## Usage

`ical.Read` parses an iCalendar stream from any `io.Reader` into a `[]*model.Calendar`. Per [RFC 5545 section 3.4](https://datatracker.ietf.org/doc/html/rfc5545#section-3.4), a stream may contain multiple sequential `VCALENDAR` objects, and `Read` handles any number of them.

```go
file, err := os.Open("calendars.ics")
if err != nil {
	return err
}
defer file.Close()

calendars, err := ical.Read(file)
```

If you expect exactly one `VCALENDAR`, use `ical.ReadSingle`, which returns a single `*model.Calendar` and fails with `ErrContentAfterEndBlock` if anything (including a second calendar) follows `END:VCALENDAR`:

```go
calendar, err := ical.ReadSingle(strings.NewReader(icalData))
```


## Performance

These numbers measure each library’s public parse entrypoint on the same ICS bytes. They do not measure recurrence expansion or serialization.

**This is not the same amount of work.** simple-ical parses into typed fields (`Event.Summary`, `DateTime`, `rrule.RRule`, and so on) and checks RFC 5545 rules (required properties, duplicates, DTEND vs DURATION, TZID references, enums). [golang-ical](https://github.com/arran4/golang-ical/releases/tag/v0.3.5) and [emersion/go-ical](https://pkg.go.dev/github.com/emersion/go-ical@v0.0.0-20250609112844-439c63cef608) stop at a generic property/component tree: names, raw string values, and nested `BEGIN`/`END`. They do not type dates or RRULEs, and they do not validate the calendar, at parse time. A faster simple-ical number here means it did *more* than those two columns, not less.

[gocal v0.9.1](https://github.com/apognu/gocal/releases/tag/v0.9.1) is a different product again: VEVENT-only, and it expands recurrences during `Parse()`. It only appears on VEVENT-only files with no `RRULE`. golang-ical and emersion can serialize; that is not measured here. simple-ical is v0.6.1.

### Specs

All benchmarks were run on a 5700X3D processor with 32 GB of RAM.

Numbers will be filled from a local `make bench-comparative` run.

### VEVENT parse

VEVENT-only calendars with no `RRULE`, `VTIMEZONE`, `VTODO`, `VALARM`, or `VJOURNAL`.

#### Simple Event

|         | SimpleIcal | Gocal | GolangIcal | Emersion |
|---------|------------|-------|------------|----------|
| sec/op  |            |       |            |          |
| B/op    |            |       |            |          |
|allocs/op|            |       |            |          |

#### Rich Event

|         | SimpleIcal | Gocal | GolangIcal | Emersion |
|---------|------------|-------|------------|----------|
| sec/op  |            |       |            |          |
| B/op    |            |       |            |          |
|allocs/op|            |       |            |          |

#### Multiple Events

|         | SimpleIcal | Gocal | GolangIcal | Emersion |
|---------|------------|-------|------------|----------|
| sec/op  |            |       |            |          |
| B/op    |            |       |            |          |
|allocs/op|            |       |            |          |

### Calendar parse

Full calendar objects (timezones, `RRULE` stored but not expanded, todos, alarms, journals). gocal is omitted.

#### Single Event

|         | SimpleIcal | GolangIcal | Emersion |
|---------|------------|------------|----------|
| sec/op  |            |            |          |
| B/op    |            |            |          |
|allocs/op|            |            |          |

#### Multiple Events

|         | SimpleIcal | GolangIcal | Emersion |
|---------|------------|------------|----------|
| sec/op  |            |            |          |
| B/op    |            |            |          |
|allocs/op|            |            |          |

#### Complex Calendar

|         | SimpleIcal | GolangIcal | Emersion |
|---------|------------|------------|----------|
| sec/op  |            |            |          |
| B/op    |            |            |          |
|allocs/op|            |            |          |
