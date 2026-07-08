// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package rrule

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/michael-gallo/simpleical/icaldur"
)

// Frequency represents the RRULE FREQ property.
type Frequency string

const (
	FrequencySecondly Frequency = "SECONDLY"
	FrequencyMinutely Frequency = "MINUTELY"
	FrequencyHourly   Frequency = "HOURLY"
	FrequencyDaily    Frequency = "DAILY"
	FrequencyWeekly   Frequency = "WEEKLY"
	FrequencyMonthly  Frequency = "MONTHLY"
	FrequencyYearly   Frequency = "YEARLY"
)

// Weekday is a two-letter weekday abbreviation used by RRULE BYDAY property.
type Weekday string

const (
	WeekdayMonday    Weekday = "MO"
	WeekdayTuesday   Weekday = "TU"
	WeekdayWednesday Weekday = "WE"
	WeekdayThursday  Weekday = "TH"
	WeekdayFriday    Weekday = "FR"
	WeekdaySaturday  Weekday = "SA"
	WeekdaySunday    Weekday = "SU"
)

// ByDay represents a BYDAY property with an optional interval prefix.
type ByDay struct {
	// The day of the week that the event occurs on.
	Weekday Weekday
	// The interval between occurrences of the event.
	// eg: If Weekday is Tuesday, and Interval is 2, then the event will happen every other Tuesday.
	Interval int
}

// RRule represents an ical reccurence rule.
// https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.10.
type RRule struct {
	// The frequency of the event.
	// This MUST be specified.
	Frequency Frequency
	// The date and time until the rule ends, inclusive.
	// Can not occur with the Count property.
	Until *time.Time
	// The number of occurrences of the event.
	// Can not occur with the Until property.
	// DTStart always counts as the first occurrence.
	Count *int
	// The interval between occurrences of the event.
	// eg: an interval of 2 for a daily rule means the event will happen every other day.
	// Not mandatory, but treated as 1 if not present.
	Interval int
	// BYSECOND is a comma separated list of seconds of the minute that the event occurs on.
	BySecond []uint8

	// ByMinute is a comma separated list of minutes of the hour that the event occurs on.
	ByMinute []uint8

	// ByHour is a comma separated list of hours of the day that the event occurs on.
	ByHour []uint8
	// The day of the week that the event occurs on.
	// This is optional and repeatable.
	ByDay []ByDay

	// ByMonthDay is the day of the month that the event occurs on.
	// eg: 10th of the month, negative numbers are allowed to indicate the last day of the month.
	// for example, -3 is the third-to-last-day of the month.
	ByMonthDay []int

	// ByYearDay is the day of the year that the event occurs on.
	// eg: 100th day of the year, negative numbers are allowed to indicate the last day of the year.
	ByYearDay []int

	// ByWeekNo is the week number that the event occurs on.
	// eg: 20th week of the year, negative numbers are allowed to indicate the last week of the year.
	ByWeekNo []int8

	// The ByMonth(s) of the year that the event occurs on.
	ByMonth []int

	// BySetPos (by set position) is the position of the last BY- component to use.
	// eg: if FREQ=WEEKLY, BYDAY=TU,WE,TH and BySetPos=1, then the event will happen on the first Tuesday, Wednesday, or Thursday of the week
	BySetPos []int16
	// WKST (Week Start) is the first day of the work week. If not set, it defaults to Monday.
	WKST Weekday
}

// ParseRRule takes an iCal reccurence rule string and parses it into a RRule struct.
// https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.10.
// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.5.3.
func ParseRRule(rruleString string) (*RRule, error) {
	rrule := &RRule{}
	intervalSet := false
	for part := range strings.SplitSeq(rruleString, ";") {
		tag, value, found := strings.Cut(part, "=")
		if !found {
			return nil, errInvalidRRuleString
		}
		switch tag {
		case "FREQ":
			if err := setOnceValue(&rrule.Frequency, Frequency(value), tag); err != nil {
				return nil, err
			}
			// Validate the frequency is valid
			if !isValidFrequency(rrule.Frequency) {
				return nil, fmt.Errorf("%w: %s", errInvalidFrequency, value)
			}
		case "INTERVAL":
			interval, err := strconv.Atoi(value)
			if err != nil {
				return nil, err
			}
			if err := setOnceInterval(&rrule.Interval, &intervalSet, interval, tag); err != nil {
				return nil, err
			}
		case "COUNT":
			count, err := strconv.Atoi(value)
			if err != nil {
				return nil, err
			}
			if err := setOncePointer(&rrule.Count, count, tag); err != nil {
				return nil, err
			}
		case "UNTIL":
			until, err := icaldur.ParseIcalTime(value)
			if err != nil {
				return nil, err
			}
			if err := setOncePointer(&rrule.Until, until, tag); err != nil {
				return nil, err
			}
		case "BYDAY":
			weekdays := strings.Split(value, ",")
			byDay := make([]ByDay, 0, len(weekdays))
			for _, weekday := range weekdays {
				// if there is an interval other than 1, it can be expressed as the number at the start of the string
				interval, weekday, err := parseByDay(weekday)
				if err != nil {
					return nil, err
				}
				byDay = append(byDay, ByDay{Weekday: weekday, Interval: interval})
			}
			if err := setOnceSlice(&rrule.ByDay, byDay, tag); err != nil {
				return nil, err
			}
		case "BYMONTH":
			months := strings.Split(value, ",")
			byMonth := make([]int, 0, len(months))
			for _, month := range months {
				monthInt, err := strconv.Atoi(month)
				if err != nil {
					return nil, err
				}
				if monthInt < 1 || monthInt > 12 {
					return nil, errInvalidByMonth
				}
				byMonth = append(byMonth, monthInt)
			}
			if err := setOnceSlice(&rrule.ByMonth, byMonth, tag); err != nil {
				return nil, err
			}
		case "BYMONTHDAY":
			monthdays := strings.Split(value, ",")
			byMonthDay := make([]int, 0, len(monthdays))
			for _, monthday := range monthdays {
				monthdayInt, err := strconv.Atoi(monthday)
				if err != nil {
					return nil, err
				}
				if !validByMonthDay(monthdayInt) {
					return nil, errInvalidByMonthDay
				}
				byMonthDay = append(byMonthDay, monthdayInt)
			}
			if err := setOnceSlice(&rrule.ByMonthDay, byMonthDay, tag); err != nil {
				return nil, err
			}
		case "BYYEARDAY":
			yeardays := strings.Split(value, ",")
			byYearDay := make([]int, 0, len(yeardays))
			for _, yearday := range yeardays {
				yeardayInt, err := strconv.Atoi(yearday)
				if err != nil {
					return nil, err
				}
				if !validByYearDay(yeardayInt) {
					return nil, errInvalidByYearDay
				}
				byYearDay = append(byYearDay, yeardayInt)
			}
			if err := setOnceSlice(&rrule.ByYearDay, byYearDay, tag); err != nil {
				return nil, err
			}
		case "WKST":
			if !isValidWeekday(Weekday(value)) {
				return nil, errInvalidWeekday
			}
			if err := setOnceValue(&rrule.WKST, Weekday(value), tag); err != nil {
				return nil, err
			}
		case "BYWEEKNO":
			weekNumbers, err := parseSignedIntListBounded(value, int8(53), errInvalidWeekno)
			if err != nil {
				return nil, err
			}
			if err := setOnceSlice(&rrule.ByWeekNo, weekNumbers, tag); err != nil {
				return nil, err
			}
		case "BYSETPOS":
			bySetPos, err := parseSignedIntListBounded(value, int16(366), errInvalidBySetPos)
			if err != nil {
				return nil, err
			}
			if err := setOnceSlice(&rrule.BySetPos, bySetPos, tag); err != nil {
				return nil, err
			}
		case "BYMINUTE":
			minutes := strings.Split(value, ",")
			byMinute := make([]uint8, 0, len(minutes))
			for _, minute := range minutes {
				minuteInt, err := strconv.ParseUint(minute, 10, 8)
				if err != nil {
					return nil, err
				}
				if minuteInt > 59 {
					return nil, errInvalidByMinute
				}
				byMinute = append(byMinute, uint8(minuteInt))
			}
			if err := setOnceSlice(&rrule.ByMinute, byMinute, tag); err != nil {
				return nil, err
			}
		case "BYHOUR":
			hours := strings.Split(value, ",")
			byHour := make([]uint8, 0, len(hours))
			for _, hour := range hours {
				hourInt, err := strconv.ParseUint(hour, 10, 8)
				if err != nil {
					return nil, err
				}
				if hourInt > 23 {
					return nil, errInvalidByHour
				}
				byHour = append(byHour, uint8(hourInt))
			}
			if err := setOnceSlice(&rrule.ByHour, byHour, tag); err != nil {
				return nil, err
			}
		case "BYSECOND":
			seconds := strings.Split(value, ",")
			bySecond := make([]uint8, 0, len(seconds))
			for _, second := range seconds {
				secondInt, err := strconv.ParseUint(second, 10, 8)
				if err != nil {
					return nil, err
				}
				if secondInt > 59 {
					return nil, errInvalidBySecond
				}
				bySecond = append(bySecond, uint8(secondInt))
			}
			if err := setOnceSlice(&rrule.BySecond, bySecond, tag); err != nil {
				return nil, err
			}
		default:
			return nil, fmt.Errorf("%w: %s", errInvalidRRuleString, tag)
		}
	}
	if !intervalSet {
		rrule.Interval = 1
	}
	if err := validateRRule(rrule); err != nil {
		return nil, err
	}
	return rrule, nil
}

func setOnceValue[T comparable](field *T, value T, tag string) error {
	var zero T
	if *field != zero {
		return fmt.Errorf("%w: %s", errDuplicateRRulePart, tag)
	}
	*field = value
	return nil
}

func setOnceInterval(field *int, isSet *bool, value int, tag string) error {
	if *isSet {
		return fmt.Errorf("%w: %s", errDuplicateRRulePart, tag)
	}
	*field = value
	*isSet = true
	return nil
}

func setOncePointer[T any](field **T, value T, tag string) error {
	if *field != nil {
		return fmt.Errorf("%w: %s", errDuplicateRRulePart, tag)
	}
	*field = &value
	return nil
}

func setOnceSlice[T any](field *[]T, value []T, tag string) error {
	if *field != nil {
		return fmt.Errorf("%w: %s", errDuplicateRRulePart, tag)
	}
	*field = value
	return nil
}

func validateRRule(rrule *RRule) error {
	if rrule.Frequency == "" {
		return errFrequencyRequired
	}
	if len(rrule.ByWeekNo) > 0 && rrule.Frequency != FrequencyYearly {
		return errByWeekNoWithInvalidFrequency
	}
	if rrule.Count != nil && rrule.Until != nil {
		return errCountAndUntilBothSet
	}
	if rrule.Interval <= 0 {
		return errInvalidInterval
	}
	return nil
}

func validByMonthDay(v int) bool {
	return (v >= 1 && v <= 31) || (v <= -1 && v >= -31)
}

func validByYearDay(v int) bool {
	return (v >= 1 && v <= 366) || (v <= -1 && v >= -366)
}

// parseSignedIntListBounded parses a comma-separated list of signed integers into a
// compact slice type. This is a perf micro-optimization: a single-pass parser avoids
// strings.Split and strconv, validates RFC ranges during parsing, and stores values as
// int8/int16. Used for BYWEEKNO and BYSETPOS.
func parseSignedIntListBounded[T ~int8 | ~int16](value string, maxVal T, outOfRange error) ([]T, error) {
	parsed := make([]T, 0, 4)
	maxInt := int(maxVal)
	i := 0
	for i < len(value) {
		neg := false
		switch value[i] {
		case '-':
			neg = true
			i++
			if i == len(value) {
				return nil, strconv.ErrSyntax
			}
		case '+':
			i++
			if i == len(value) {
				return nil, strconv.ErrSyntax
			}
		}

		start := i
		n := 0
		for i < len(value) && value[i] >= '0' && value[i] <= '9' {
			n = n*10 + int(value[i]-'0')
			if n > maxInt {
				return nil, outOfRange
			}
			i++
		}
		if i == start {
			return nil, strconv.ErrSyntax
		}
		if neg {
			if n > maxInt {
				return nil, outOfRange
			}
			n = -n
		}
		if n == 0 {
			return nil, outOfRange
		}
		parsed = append(parsed, T(n))

		if i == len(value) {
			break
		}
		if value[i] != ',' {
			return nil, strconv.ErrSyntax
		}
		i++
	}
	if len(parsed) == 0 {
		return nil, strconv.ErrSyntax
	}
	return parsed, nil
}

// parseByDay parses a BYDAY value string and returns the interval and weekday.
// The string can be in the format "20MO" (interval + weekday) or just "MO" (weekday only).
// If no interval is specified, the interval defaults to 1.
// Valid weekdays are: MO, TU, WE, TH, FR, SA, SU.
// Returns (interval, weekday, error) where interval is an integer and weekday is a string.
func parseByDay(byDayString string) (int, Weekday, error) {
	if byDayString == "" {
		return 0, "", errInvalidByDayString
	}

	// Check if string starts with a digit or minus sign
	if len(byDayString) > 0 && (byDayString[0] >= '0' && byDayString[0] <= '9' || byDayString[0] == '-') {
		// Find where the digits end (including negative sign)
		digitEnd := 0
		for i, char := range byDayString {
			if char < '0' || char > '9' {
				// Allow minus sign at the beginning
				if char == '-' && i == 0 {
					continue
				}
				digitEnd = i
				break
			}
			digitEnd = i + 1
		}

		// Extract interval and weekday
		intervalStr := byDayString[:digitEnd]
		weekday := Weekday(byDayString[digitEnd:])

		// Validate weekday
		if !isValidWeekday(weekday) {
			return 0, "", errInvalidByDayString
		}

		// Parse interval (can be negative)
		interval, err := strconv.Atoi(intervalStr)
		if err != nil {
			return 0, "", errInvalidByDayString
		}

		return interval, weekday, nil
	}

	// No interval prefix, check if it's a valid weekday
	if !isValidWeekday(Weekday(byDayString)) {
		return 0, "", errInvalidByDayString
	}

	return 1, Weekday(byDayString), nil
}

// isValidWeekday checks if the string is a valid weekday abbreviation.
func isValidWeekday(weekday Weekday) bool {
	switch weekday {
	case WeekdayMonday, WeekdayTuesday, WeekdayWednesday, WeekdayThursday, WeekdayFriday, WeekdaySaturday, WeekdaySunday:
		return true
	default:
		return false
	}
}

func isValidFrequency(frequency Frequency) bool {
	switch frequency {
	case FrequencySecondly, FrequencyMinutely, FrequencyHourly, FrequencyDaily, FrequencyWeekly, FrequencyMonthly, FrequencyYearly:
		return true
	default:
		return false
	}
}
