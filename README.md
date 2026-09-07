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
Performance tests are for simple-ical v0.6.1 and were ran against [golang-ical v0.3.5](https://github.com/arran4/golang-ical/releases/tag/v0.3.5) and [gocal v0.9.1](https://github.com/apognu/gocal/releases/tag/v0.9.1)

### Specs
All tests were ran on a 5700X3D Processor with 32GB of RAM.

### Simple Event
|         | SimpleIcal  | Gocal       | GolangIcal  |
|---------|-------------|-------------|--------------|
| sec/op  | 2.173µ ± 0% | 4.484µ ± 1% | 8.611µ ± 0% |
| B/op    | 5.391Ki ± 0% | 7.028Ki ± 0% | 7.983Ki ± 0% |
|allocs/op| 16 ± 0% | 70 ± 0% | 144 ± 0% |

### Single Event
|         | SimpleIcal  | Gocal       | GolangIcal  |
|---------|-------------|-------------|--------------|
| sec/op  | 7.160µ ± 1% | 13.25µ ± 0% | 28.72µ ± 1% |
| B/op    | 7.867Ki ± 0% | 13.12Ki ± 0% | 19.36Ki ± 0% |
|allocs/op| 63 ± 0% | 241 ± 0% | 465 ± 0% |

### Multiple Events
|         | SimpleIcal  | Gocal       | GolangIcal  |
|---------|-------------|-------------|--------------|
| sec/op  | 10.25µ ± 0% | 20.86µ ± 0% | 43.59µ ± 1% |
| B/op    | 10.63Ki ± 0% | 18.15Ki ± 0% | 27.91Ki ± 0% |
|allocs/op| 97 ± 0% | 382 ± 0% | 712 ± 0% |

### Complex Calendar
|         | SimpleIcal  | Gocal       | GolangIcal  |
|---------|-------------|-------------|--------------|
| sec/op  | 12.81µ ± 1% | 19.62µ ± 0% | 57.99µ ± 0% |
| B/op    | 11.93Ki ± 0% | 18.91Ki ± 0% | 33.22Ki ± 0% |
|allocs/op| 118 ± 0% | 417 ± 0% | 965 ± 0% |
