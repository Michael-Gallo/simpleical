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

## Comparison

| | simple-ical | gocal | golang-ical | emersion/go-ical |
|---|---|---|---|---|
| Output | Typed `Calendar` / `Event` / `Todo` / … | `[]Event` with `time.Time` | Generic property tree | Generic component tree |
| RFC 5545 checks on parse | Required props, duplicates, enums, TZID refs | UID / DTSTART / DTSTAMP (optional strictness) | Line/structure syntax | `BEGIN`/`END` matching |
| `VEVENT` | Yes | Yes | Yes | Yes |
| `VTODO`, `VJOURNAL`, `VALARM`, `VTIMEZONE`, `VFREEBUSY` | Yes | Ignored | Stored as components | Stored as components |
| `RRULE` | Parsed into a struct, not expanded | Parsed and expanded during `Parse()` | Left as a string until getters | Left as a string until `RecurrenceSet()` |
| Serialize | No | No | Yes | Yes |
| Multiple sequential `VCALENDAR`s | `ical.Read` | No | One calendar per `ParseCalendar` | One calendar per `Decode` (loop for more) |

## Performance

These numbers measure each library’s public parse entrypoint on the same ICS bytes.

**This is not the same amount of work.** simple-ical parses into typed fields (`Event.Summary`, `DateTime`, `rrule.RRule`, and so on) and checks RFC 5545 rules (required properties, duplicates, DTEND vs DURATION, TZID references, enums). [golang-ical](https://github.com/arran4/golang-ical/releases/tag/v0.3.6) and [emersion/go-ical](https://pkg.go.dev/github.com/emersion/go-ical@v0.0.0-20250609112844-439c63cef608) do not map properties into typed fields and they do not validate the calendar.

[gocal v0.9.1](https://github.com/apognu/gocal/releases/tag/v0.9.1) is in many ways a different product: VEVENT-only, and it expands recurrences during `Parse()`. Because of this we only compare it to VEVENT-only calendars without recurrence rules. golang-ical and emersion can serialize; that is not measured here. simple-ical is v0.6.1.

### Specs

All benchmarks were run on a 5700X3D processor with 32 GB of RAM.

### VEVENT parse

VEVENT-only calendars with no `RRULE`, `VTIMEZONE`, `VTODO`, `VALARM`, or `VJOURNAL`.

#### Simple Event

|           | SimpleIcal | Gocal | GolangIcal | Emersion |
| --------- | ---------- | ----- | ---------- | -------- |
| sec/op    | 2.172µ     | 4.611µ | 9.154µ    | 2.538µ   |
| B/op      | 5.391Ki    | 7.028Ki | 8.062Ki  | 6.477Ki  |
| allocs/op | 16         | 70    | 146        | 49       |

#### Rich Event

|           | SimpleIcal | Gocal | GolangIcal | Emersion |
| --------- | ---------- | ----- | ---------- | -------- |
| sec/op    | 5.656µ     | 9.715µ | 23.02µ    | 7.744µ   |
| B/op      | 7.008Ki    | 10.90Ki | 16.55Ki  | 12.31Ki  |
| allocs/op | 50         | 165   | 354        | 140      |

#### Multiple Events

|           | SimpleIcal | Gocal  | GolangIcal | Emersion |
| --------- | ---------- | ------ | ---------- | -------- |
| sec/op    | 8.352µ     | 15.61µ | 37.76µ     | 11.99µ   |
| B/op      | 9.438Ki    | 15.48Ki | 24.71Ki   | 17.82Ki  |
| allocs/op | 81         | 272    | 591        | 211      |

### Calendar parse

Full calendar objects (timezones, `RRULE` stored but not expanded, todos, alarms, journals). gocal is omitted.

#### Single Event

|           | SimpleIcal | GolangIcal | Emersion |
| --------- | ---------- | ---------- | -------- |
| sec/op    | 7.155µ     | 29.65µ     | 9.500µ   |
| B/op      | 7.867Ki    | 19.51Ki    | 14.48Ki  |
| allocs/op | 63         | 467        | 180      |

#### Multiple Events

|           | SimpleIcal | GolangIcal | Emersion |
| --------- | ---------- | ---------- | -------- |
| sec/op    | 10.17µ     | 44.87µ     | 13.99µ   |
| B/op      | 10.63Ki    | 28.09Ki    | 20.19Ki  |
| allocs/op | 97         | 714        | 255      |

#### Complex Calendar

|           | SimpleIcal | GolangIcal | Emersion |
| --------- | ---------- | ---------- | -------- |
| sec/op    | 12.67µ     | 60.04µ     | 17.14µ   |
| B/op      | 11.93Ki    | 33.41Ki    | 23.30Ki  |
| allocs/op | 118        | 967        | 338      |
