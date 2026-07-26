# AGENTS.md

This is an icalendar parser focused on performance and compliance with the RFC5545 spec

# RFC 5545 Reference

- The local RFC 5545 spec lives in `docs/specs/rfc5545/` (split into Markdown for targeted lookup). The full canonical text is `docs/specs/rfc5545-icalendar.txt`.
- Prefer reading the local spec over web search when verifying iCalendar behavior, property syntax, or value-type rules.
- Useful entry points:
  - `docs/specs/rfc5545/value-types/` — data type definitions (e.g. `recur.md`, `integer.md`)
  - `docs/specs/rfc5545/properties/` — property definitions (e.g. `rrule.md`)
  - `docs/specs/rfc5545/components/` — component definitions (e.g. `vevent.md`)

# Setting Properties

- We have `setOnce` functions in property_setters.go, which handle errors related to setting duplicate properties; please use these when appropriate

# Tests

- Please ensure that, when dealing with icalendar properties, we have integration test coverage in the `test/` folder
- `testify` is an intentional test dependency: `assert.Equal` on large parsed calendars/components gives readable diffs. Do not rip it out just to eliminate a test-only dep.

## iCalendar test fixtures

- iCalendar input belongs in a `.ical` file under `test/test_data/<area>/`, pulled in with `//go:embed` in the test file's `var` block. Do not inline calendar bodies as Go string literals; a real file can be diffed, opened in an editor, and fed to another parser to cross-check our reading of the spec.
- Name fixtures for what they assert: `valid_*.ical` for input that must parse, `invalid_*.ical` for input that must be rejected.
- Fixtures are byte-sensitive. Trailing whitespace and line endings are part of the test. Any fixture that needs CRLF must be listed in `.gitattributes` with `text eol=crlf`, or git will normalize it away.
- LF is fine for fixtures that only ever feed our own parser, since the scanner strips a trailing `\r`. Fixtures meant to be read by other tools must be CRLF: RFC 5545 section 3.1 delimits content lines with CRLF, and external validators flag LF-only input.

## Known spec gaps

- `test/spec_gap_test.go` holds cases where the parser disagrees with RFC 5545. Each test asserts the behavior the spec requires and is skipped with a `gap:` message citing the section that justifies it.
- Its fixtures in `test/test_data/spec_gaps/` exist to be pasted into third-party validators, so keep each one a complete, conformant calendar object that differs from a clean file only in the single thing under test. Note that third-party validators are not authoritative and do produce false positives; the local spec in `docs/specs/rfc5545/` wins.
- Before skipping a new gap test, run it once un-skipped and confirm it fails for the reason you expect. A skipped test that would have passed anyway documents nothing.
- When you fix a gap, un-skip its test in the same change. That is how the fix is verified.

# Public API

- Be conservative about expanding the public API. Prefer unexported helpers unless a symbol is needed by other packages or is intentionally part of the documented surface. Do not export low-level helpers just for symmetry with related exported functions.
- Parser sentinel errors live in `internal/icalerr` on purpose: they stay off the public API while integration tests in `test/` can still `errors.Is` against them. Do not fold that package into `ical` (or export the sentinels) just to remove an internal package.

# Documentation

- Every exported function must have a corresponding `Example...` in an `_test.go` file so it appears in godoc
- Unexported functions and methods should also have a short leading comment describing purpose and non-obvious constraints
