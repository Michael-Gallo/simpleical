// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package rrule

import "errors"

// Predefined errors for the rrule package.
var (
	// errInvalidRRuleString is returned when the rrule string format is invalid.
	errInvalidRRuleString = errors.New("invalid rrule string")

	// errDuplicateRRulePart is returned when the same recur-rule-part appears more than once (RFC 5545).
	errDuplicateRRulePart = errors.New("duplicate RRULE part")

	// errFrequencyRequired is returned when the frequency property is missing.
	errFrequencyRequired = errors.New("frequency is required")

	// errCountAndUntilBothSet is returned when both count and until properties are set.
	errCountAndUntilBothSet = errors.New("count and until cannot both be set")

	// errInvalidInterval is returned when the interval is not a positive integer.
	errInvalidInterval = errors.New("interval must be a positive integer")

	// errInvalidByDayString is returned when the BYDAY string format is invalid.
	errInvalidByDayString = errors.New("invalid BYDAY string")

	errInvalidFrequency = errors.New("invalid frequency")

	errInvalidWeekday = errors.New("invalid weekday")

	errInvalidWeekno = errors.New("invalid week number: must be between 1 and 53 or -1 and -53")

	errByWeekNoWithInvalidFrequency = errors.New("BYWEEKNO is only allowed for yearly frequency")

	errInvalidBySetPos = errors.New("BYSETPOS must be between -366 and 366, not 0")

	errInvalidByMinute = errors.New("BYMINUTE must be between 0 and 59")

	errInvalidByHour = errors.New("BYHOUR must be between 0 and 23")

	errInvalidBySecond = errors.New("BYSECOND must be between 0 and 59")

	errInvalidByMonth = errors.New("BYMONTH out of range: must be 1..12")

	errInvalidByMonthDay = errors.New("BYMONTHDAY out of range: must be -31..-1 or 1..31")

	errInvalidByYearDay = errors.New("BYYEARDAY out of range: must be -366..-1 or 1..366")
)
