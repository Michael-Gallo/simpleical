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
- **golang-ical** and **emersion/go-ical** parse the full calendar into a generic property/component tree. They can serialize; that is not measured here. simple-ical parses into typed structures and validates during parse.

### Specs

All benchmarks were run on an Intel Xeon processor with 16 GB of RAM.

### VEVENT parse

VEVENT-only calendars with no `RRULE`, `VTIMEZONE`, `VTODO`, `VALARM`, or `VJOURNAL`.

#### Simple Event

|         | SimpleIcal   | Gocal        | GolangIcal   | Emersion     |
|---------|--------------|--------------|--------------|--------------|
| sec/op  | 2.536µ ± 1%  | 4.698µ ± 1%  | 8.203µ ± 1%  | 3.229µ ± 0%  |
| B/op    | 5.391Ki ± 0% | 7.024Ki ± 0% | 7.918Ki ± 0% | 6.477Ki ± 0% |
|allocs/op| 16 ± 0%      | 70 ± 0%      | 144 ± 0%     | 49 ± 0%      |

#### Rich Event

|         | SimpleIcal   | Gocal        | GolangIcal    | Emersion      |
|---------|--------------|--------------|---------------|---------------|
| sec/op  | 4.743µ ± 1%  | 8.546µ ± 0%  | 19.43µ ± 0%   | 7.271µ ± 1%   |
| B/op    | 7.008Ki ± 0% | 10.89Ki ± 0% | 16.29Ki ± 0%  | 12.31Ki ± 0%  |
|allocs/op| 50 ± 0%      | 165 ± 0%     | 352 ± 0%      | 140 ± 0%      |

#### Multiple Events

|         | SimpleIcal   | Gocal        | GolangIcal    | Emersion      |
|---------|--------------|--------------|---------------|---------------|
| sec/op  | 6.961µ ± 1%  | 14.94µ ± 8%  | 31.70µ ± 1%   | 11.10µ ± 1%   |
| B/op    | 9.438Ki ± 0% | 15.47Ki ± 0% | 24.38Ki ± 0%  | 17.82Ki ± 0%  |
|allocs/op| 81 ± 0%      | 272 ± 0%     | 589 ± 0%      | 211 ± 0%      |

### Calendar parse

Full calendar objects (timezones, `RRULE` stored but not expanded, todos, alarms, journals). gocal is omitted.

#### Single Event

|         | SimpleIcal   | GolangIcal    | Emersion      |
|---------|--------------|---------------|---------------|
| sec/op  | 5.787µ ± 6%  | 26.49µ ± 2%   | 10.03µ ± 3%   |
| B/op    | 7.867Ki ± 0% | 19.22Ki ± 0%  | 14.48Ki ± 0%  |
|allocs/op| 63 ± 0%      | 465 ± 0%      | 180 ± 0%      |

#### Multiple Events

|         | SimpleIcal    | GolangIcal    | Emersion      |
|---------|---------------|---------------|---------------|
| sec/op  | 9.369µ ± 1%   | 40.34µ ± 1%   | 14.33µ ± 1%   |
| B/op    | 10.63Ki ± 0%  | 27.72Ki ± 0%  | 20.19Ki ± 0%  |
|allocs/op| 97 ± 0%       | 712 ± 0%      | 255 ± 0%      |

#### Complex Calendar

|         | SimpleIcal    | GolangIcal    | Emersion      |
|---------|---------------|---------------|---------------|
| sec/op  | 10.84µ ± 1%   | 53.80µ ± 2%   | 17.85µ ± 1%   |
| B/op    | 11.93Ki ± 0%  | 32.94Ki ± 0%  | 23.30Ki ± 0%  |
|allocs/op| 118 ± 0%      | 965 ± 0%      | 338 ± 0%      |
