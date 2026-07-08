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
