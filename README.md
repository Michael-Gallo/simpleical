# Simple-ical

A very much not ready ICAL parser for Golang intended to follow the official [ICAL 2.0 spec](https://datatracker.ietf.org/doc/html/rfc5545) as closely as is reasonable.

Focused on ease of use and good documentation, with frequent links to the spec.

[![Go Reference](https://pkg.go.dev/badge/github.com/michael-gallo/simpleical.svg)](https://pkg.go.dev/github.com/michael-gallo/simpleical)

## Documentation

Full API documentation is available on [pkg.go.dev](https://pkg.go.dev/github.com/michael-gallo/simpleical).

## Deviations from spec

1. The VCALENDAR spec does not address whitespace at the end of lines. We assume in this parser it is to be ignored and right trim all whitespace.
2. The `DTSTAMP` property is [mandatory](https://datatracker.ietf.org/doc/html/rfc5545#section-3.6.1), however, I have seen real life examples where it is not filled out. Ergo I will not be enforcing it here. If I do enforce it in the future, it will be in an opt-in strict mode.

## License

This project is licensed under the Mozilla Public License 2.0. See the [LICENSE](LICENSE) file for details.


## Installation


```sh
go get github.com/michael-gallo/simpleical
```


## Performance
Performance tests are for simple-ical v 0.3.2 and were ran against [golang-ical v0.3.2](https://github.com/arran4/golang-ical/releases/tag/v0.3.2) and [gocal v0.9.1](https://github.com/apognu/gocal/releases/tag/v0.9.1)

### Specs
All tests were ran on a 5700X3D Processor with 32GB of RAM.

### Calendar File With Minimal Event

|         | SimpleIcal  | Gocal       | GolangIcal  |
|---------|-------------|-------------|-------------|
| sec/op  | 1.391µ ± 1% | 4.723µ ± 1% | 8.920µ ± 1% |
| B/op    | 5.023Ki ± 0%| 7.028Ki ± 0%| 7.938Ki ± 0%|
|allocs/op| 13.00 ± 0%  | 70.00 ± 0%  | 144.0 ± 0%  |

### Calendar File with single representative event

|         | SimpleIcal  | Gocal       | GolangIcal  |
|---------|-------------|-------------|-------------|
| sec/op  | 5.257µ ± 0% | 14.19µ ± 0% | 29.40µ ± 1% |
| B/op    | 7.039Ki ± 0%| 13.12Ki ± 0%| 19.28Ki ± 0%|
|allocs/op| 62.00 ± 0%  | 241.0 ± 0%  | 465.0 ± 0%  |


### Calendar File with Multiple Events

|         | SimpleIcal  | Gocal       | GolangIcal  |
|---------|-------------|-------------|-------------|
| sec/op  | 6.954µ ± 1% | 22.37µ ± 0% | 44.79µ ± 0% |
| B/op    | 9.062Ki ± 0%| 18.15Ki ± 0%| 27.79Ki ± 0%|
|allocs/op| 91.00 ± 0%  | 382.0 ± 0%  | 712.0 ± 0%  |


### Calendar File with Events and TODOs

|         | SimpleIcal  | Gocal       | GolangIcal  |
|---------|-------------|-------------|-------------|
| sec/op  | 8.592µ ± 1% | 21.46µ ± 1% | 59.87µ ± 1% |
| B/op    | 10.44Ki ± 0%| 19.01Ki ± 0%| 33.39Ki ± 0%|
|allocs/op| 113.0 ± 0%  | 418.0 ± 0%  | 970.0 ± 0%  |
