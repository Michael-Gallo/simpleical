# Benchmark baseline

Stable allocation baselines for `BenchmarkAllScenarios` (simple-ical only),
measured with:

```bash
cd benchmarks && go test -bench=BenchmarkAllScenarios -benchmem -count=5
```

Use these to catch regressions; do **not** commit host-specific raw `go test` output.

| Scenario          | B/op  | allocs/op |
|-------------------|------:|----------:|
| Simple Event      |  5096 |        14 |
| Single Event      |  7104 |        57 |
| Multiple Events   |  9168 |        88 |
| Complex Calendar  | 10576 |       106 |

**Regression target:** no increase in `B/op` or `allocs/op` for any scenario above.
`ns/op` is host- and load-dependent; compare it only on the same machine (or via
benchstat against a local previous run).

Raw `go test -bench` output belongs in local/CI artifacts under `artifacts/`
(see README), not in git.
