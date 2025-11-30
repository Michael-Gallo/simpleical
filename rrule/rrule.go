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
	ByWeekNo int8

	// The ByMonth(s) of the year that the event occurs on.
	ByMonth []int

	// BySetPos (by set position) is the position of the last BY- component to use.
	// eg: if FREQ=WEEKLY, BYDAY=TU,WE,TH and BySetPos=1, then the event will happen on the first Tuesday, Wednesday, or Thursday of the week
	BySetPos int
	// WKST (Week Start) is the first day of the work week. If not set, it defaults to Monday.
	WKST Weekday
}

// ParseRRule takes an iCal reccurence rule string and parses it into a RRule struct.
// https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.10.
// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.5.3.
func ParseRRule(rruleString string) (*RRule, error) {
	rrule := &RRule{
		// Default to 1 if not present
		Interval: 1,
	}
	for part := range strings.SplitSeq(rruleString, ";") {
		tag, value, found := strings.Cut(part, "=")
		if !found {
			return nil, errInvalidRRuleString
		}
		switch tag {
		case "FREQ":
			// Validate the frequency is valid
			if !isValidFrequency(Frequency(value)) {
				return nil, fmt.Errorf("%w: %s", errInvalidFrequency, value)
			}
			rrule.Frequency = Frequency(value)
		case "INTERVAL":
			interval, err := strconv.Atoi(value)
			if err != nil {
				return nil, err
			}
			rrule.Interval = interval
		case "COUNT":
			count, err := strconv.Atoi(value)
			if err != nil {
				return nil, err
			}
			rrule.Count = &count
		case "UNTIL":
			until, err := icaldur.ParseIcalTime(value)
			if err != nil {
				return nil, err
			}
			rrule.Until = &until
		case "BYDAY":
			weekdays := strings.Split(value, ",")
			rrule.ByDay = make([]ByDay, 0, len(weekdays))
			for _, weekday := range weekdays {
				// if there is an interval other than 1, it can be expressed as the number at the start of the string
				interval, weekday, err := parseByDay(weekday)
				if err != nil {
					return nil, err
				}
				rrule.ByDay = append(rrule.ByDay, ByDay{Weekday: weekday, Interval: interval})
			}
		case "BYMONTH":
			months := strings.Split(value, ",")
			rrule.ByMonth = make([]int, 0, len(months))
			for _, month := range months {
				monthInt, err := strconv.Atoi(month)
				if err != nil {
					return nil, err
				}
				rrule.ByMonth = append(rrule.ByMonth, monthInt)
			}
		case "BYMONTHDAY":
			monthdays := strings.Split(value, ",")
			rrule.ByMonthDay = make([]int, 0, len(monthdays))
			for _, monthday := range monthdays {
				monthdayInt, err := strconv.Atoi(monthday)
				if err != nil {
					return nil, err
				}
				rrule.ByMonthDay = append(rrule.ByMonthDay, monthdayInt)
			}
		case "BYYEARDAY":
			yeardays := strings.Split(value, ",")
			rrule.ByYearDay = make([]int, 0, len(yeardays))
			for _, yearday := range yeardays {
				yeardayInt, err := strconv.Atoi(yearday)
				if err != nil {
					return nil, err
				}
				rrule.ByYearDay = append(rrule.ByYearDay, yeardayInt)
			}
		case "WKST":
			if !isValidWeekday(Weekday(value)) {
				return nil, errInvalidWeekday
			}
			rrule.WKST = Weekday(value)
		case "BYWEEKNO":
			weekno, err := strconv.ParseInt(value, 10, 8)
			if err != nil {
				return nil, err
			}
			if weekno < 1 || weekno > 53 || weekno == 0 {
				return nil, errInvalidWeekno
			}
			rrule.ByWeekNo = int8(weekno)
		case "BYSETPOS":
			bySetPos, err := strconv.Atoi(value)
			if err != nil {
				return nil, err
			}
			if bySetPos < -366 || bySetPos > 366 || bySetPos == 0 {
				return nil, errInvalidBySetPos
			}
			rrule.BySetPos = bySetPos
		case "BYMINUTE":
			minutes := strings.Split(value, ",")
			rrule.ByMinute = make([]uint8, 0, len(minutes))
			for _, minute := range minutes {
				minuteInt, err := strconv.ParseUint(minute, 10, 8)
				if err != nil {
					return nil, err
				}
				if minuteInt > 59 {
					return nil, errInvalidByMinute
				}
				rrule.ByMinute = append(rrule.ByMinute, uint8(minuteInt))
			}
		case "BYHOUR":
			hours := strings.Split(value, ",")
			rrule.ByHour = make([]uint8, 0, len(hours))
			for _, hour := range hours {
				hourInt, err := strconv.ParseUint(hour, 10, 8)
				if err != nil {
					return nil, err
				}
				if hourInt > 23 {
					return nil, errInvalidByHour
				}
				rrule.ByHour = append(rrule.ByHour, uint8(hourInt))
			}
		case "BYSECOND":
			seconds := strings.Split(value, ",")
			rrule.BySecond = make([]uint8, 0, len(seconds))
			for _, second := range seconds {
				secondInt, err := strconv.ParseUint(second, 10, 8)
				if err != nil {
					return nil, err
				}
				if secondInt > 59 {
					return nil, errInvalidBySecond
				}
				rrule.BySecond = append(rrule.BySecond, uint8(secondInt))
			}
		}

	}
	if err := validateRRule(rrule); err != nil {
		return nil, err
	}
	return rrule, nil
}

func validateRRule(rrule *RRule) error {
	if rrule.Frequency == "" {
		return errFrequencyRequired
	}
	if rrule.ByWeekNo != 0 && rrule.Frequency != FrequencyYearly {
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
