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

These numbers measure **parse into structs**: the time and allocations to turn ICS bytes into each library’s in-memory calendar. They do not measure recurrence expansion or serialization.

Comparisons use [golang-ical v0.3.5](https://github.com/arran4/golang-ical/releases/tag/v0.3.5), [gocal v0.9.1](https://github.com/apognu/gocal/releases/tag/v0.9.1), and [emersion/go-ical](https://pkg.go.dev/github.com/emersion/go-ical@v0.0.0-20250609112844-439c63cef608) (`v0.0.0-20250609112844-439c63cef608`). simple-ical is v0.6.1.

- **gocal** only appears on VEVENT-only files with no `RRULE`. It ignores non-`VEVENT` components and expands recurrences during `Parse()`.
- **golang-ical** and **emersion/go-ical** parse the full calendar into a generic property/component tree. They can serialize; that is not measured here. simple-ical types and validates during parse.

### Specs

All tests were ran on a 5700X3D Processor with 32GB of RAM.

### VEVENT parse

VEVENT-only calendars with no `RRULE`, `VTIMEZONE`, `VTODO`, `VALARM`, or `VJOURNAL`.

#### Simple Event

|         | SimpleIcal | Gocal | GolangIcal | Emersion |
|---------|------------|-------|------------|----------|
| sec/op  | TBD | TBD | TBD | TBD |
| B/op    | TBD | TBD | TBD | TBD |
|allocs/op| TBD | TBD | TBD | TBD |

#### Rich Event

|         | SimpleIcal | Gocal | GolangIcal | Emersion |
|---------|------------|-------|------------|----------|
| sec/op  | TBD | TBD | TBD | TBD |
| B/op    | TBD | TBD | TBD | TBD |
|allocs/op| TBD | TBD | TBD | TBD |

#### Multiple Events

|         | SimpleIcal | Gocal | GolangIcal | Emersion |
|---------|------------|-------|------------|----------|
| sec/op  | TBD | TBD | TBD | TBD |
| B/op    | TBD | TBD | TBD | TBD |
|allocs/op| TBD | TBD | TBD | TBD |

### Calendar parse

Full calendar objects (timezones, `RRULE` stored but not expanded, todos, alarms, journals). gocal is omitted.

#### Single Event

|         | SimpleIcal | GolangIcal | Emersion |
|---------|------------|------------|----------|
| sec/op  | TBD | TBD | TBD |
| B/op    | TBD | TBD | TBD |
|allocs/op| TBD | TBD | TBD |

#### Multiple Events

|         | SimpleIcal | GolangIcal | Emersion |
|---------|------------|------------|----------|
| sec/op  | TBD | TBD | TBD |
| B/op    | TBD | TBD | TBD |
|allocs/op| TBD | TBD | TBD |

#### Complex Calendar

|         | SimpleIcal | GolangIcal | Emersion |
|---------|------------|------------|----------|
| sec/op  | TBD | TBD | TBD |
| B/op    | TBD | TBD | TBD |
|allocs/op| TBD | TBD | TBD |
