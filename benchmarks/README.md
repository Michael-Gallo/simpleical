# iCalendar Parser Benchmarks

This directory contains comparative benchmarks against other Go iCalendar parsers.

## Setup

From this directory:

```bash
go mod tidy
```

Dependencies (gocal, golang-ical) are already listed in `go.mod`.

## Reproducing results

Run from the repository root (or `cd benchmarks` and drop the directory prefix).

### Quick local run

```bash
make bench
# equivalent:
cd benchmarks && go test -bench=BenchmarkAllScenarios -benchmem
```

### Stable multi-sample run (recommended for comparisons)

```bash
make bench-long
# equivalent:
mkdir -p benchmarks/artifacts
cd benchmarks && go test -bench=BenchmarkAllScenarios -benchmem -count=10 > artifacts/results.txt
```

Comparative parsers:

```bash
make bench-comparative
# equivalent:
mkdir -p benchmarks/artifacts
cd benchmarks && go test -bench=BenchmarkComparativeAll -benchmem -count=10 > artifacts/results_comparative.txt
```

Raw output is written under `benchmarks/artifacts/` and is **not** committed (see `.gitignore`).
Compare runs with [benchstat](https://pkg.go.dev/golang.org/x/perf/cmd/benchstat):

```bash
benchstat artifacts/results.txt
# or against a previous local capture:
benchstat old.txt artifacts/results.txt
```

### Environment notes

- **Go version:** use the module’s `go` version from the repo root `go.mod`.
- **CPU / GOMAXPROCS:** `ns/op` varies by host; allocation metrics (`B/op`, `allocs/op`) are the stable regression signal.
- **Committed baseline:** see [`BASELINE.md`](BASELINE.md) for expected `B/op` / `allocs/op` and the no-increase regression target.
- **CI:** `.github/workflows/ci.yml` runs a short smoke bench (`-benchtime=1x -count=1`) only; full raw results are local/CI artifacts, not source.

### Other useful invocations

```bash
# Specific comparative suite
go test -bench=BenchmarkComparativeAll -benchmem

# RRULE parsing
go test -bench=BenchmarkParseRRule -benchmem

# CPU / memory profiles
make bench-profile
# or:
go test -bench=BenchmarkAllScenarios -benchmem -cpuprofile=cpu.prof -memprofile=mem.prof
go tool pprof cpu.prof
go tool pprof mem.prof
```

## What is measured

- **Parsing speed:** time per operation (`ns/op`)
- **Memory usage:** bytes allocated per operation (`B/op`)
- **Allocation count:** allocations per operation (`allocs/op`)

Parsers compared in `BenchmarkComparativeAll`: simple-ical, [gocal](https://github.com/apognu/gocal), [golang-ical](https://github.com/arran4/golang-ical).

Summarized comparative numbers also live in the root [`README.md`](../README.md#performance).

## Adding more parsers

1. Add the dependency to `go.mod`
2. Create a new benchmark file (e.g., `other_parser_benchmark.go`)
3. Implement benchmark functions following the same pattern
4. Wire them into `BenchmarkComparativeAll`

## Test data

Fixture files live in this directory (`test_*.ical`). To add a scenario:

1. Add the `.ical` file here
2. Register it in the `testCases` slices in `benchmark_setup_test.go`
