# AGENTS.md

This is an icalendar parser focused on performance and compliance with the RFC5545 spec

# Setting Properties

- We have `setOnce` functions in property_setters.go, this handles errors having to do with setting duplicate properties, please use this when appropriate

# Tests

- Please ensure that, when dealing with icalendar properties, we have integration test coverage in the `tests/` folder
