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
Performance tests are for simple-ical v5.1 and were ran against [golang-ical v0.3.5](https://github.com/arran4/golang-ical/releases/tag/v0.3.5) and [gocal v0.9.1](https://github.com/apognu/gocal/releases/tag/v0.9.1)

### Specs
All tests were ran on a 5700X3D Processor with 32GB of RAM.

### Calendar File With Minimal Event
|         | SimpleIcal  | Gocal       | GolangIcal  |
|---------|-------------|-------------|--------------|
| sec/op  | 1.454µ ± 0% | 4.736µ ± 0% | 9.506µ ± 0% |
| B/op    | 5.023Ki ± 0% | 7.028Ki ± 0% | 7.988Ki ± 0% |
|allocs/op| 13.00 ± 0% | 70.00 ± 0% | 144.00 ± 0% |

### Calendar File with single representative event
|         | SimpleIcal  | Gocal       | GolangIcal  |
|---------|-------------|-------------|--------------|
| sec/op  | 5.670µ ± 3% | 14.407µ ± 1% | 31.828µ ± 1% |
| B/op    | 7.086Ki ± 0% | 13.119Ki ± 0% | 19.375Ki ± 0% |
|allocs/op| 62.00 ± 0% | 241.00 ± 0% | 465.00 ± 0% |


### Calendar File with Multiple Events
|         | SimpleIcal  | Gocal       | GolangIcal  |
|---------|-------------|-------------|--------------|
| sec/op  | 7.398µ ± 1% | 23.020µ ± 2% | 48.202µ ± 1% |
| B/op    | 9.156Ki ± 0% | 18.147Ki ± 0% | 27.941Ki ± 0% |
|allocs/op| 91.00 ± 0% | 382.00 ± 0% | 712.00 ± 0% |


### Calendar File with Events and TODOs
|         | SimpleIcal  | Gocal       | GolangIcal  |
|---------|-------------|-------------|--------------|
| sec/op  | 8.944µ ± 1% | 21.573µ ± 0% | 64.367µ ± 0% |
| B/op    | 10.266Ki ± 0% | 19.015Ki ± 0% | 33.549Ki ± 0% |
|allocs/op| 109.00 ± 0% | 418.00 ± 0% | 970.00 ± 0% |
