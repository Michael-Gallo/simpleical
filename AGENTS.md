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

- Please ensure that, when dealing with icalendar properties, we have integration test coverage in the `tests/` folder
- `testify` is an intentional test dependency: `assert.Equal` on large parsed calendars/components gives readable diffs. Do not rip it out just to eliminate a test-only dep.

# Public API

- Be conservative about expanding the public API. Prefer unexported helpers unless a symbol is needed by other packages or is intentionally part of the documented surface. Do not export low-level helpers just for symmetry with related exported functions.
- Parser sentinel errors live in `internal/icalerr` on purpose: they stay off the public API while integration tests in `test/` can still `errors.Is` against them. Do not fold that package into `ical` (or export the sentinels) just to remove an internal package.

# Godoc Examples

- Every exported function must have a corresponding `Example...` in an `_test.go` file so it appears in godoc
