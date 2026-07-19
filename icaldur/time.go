package icaldur

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

var (
	ErrInvalidTimeFormat = errors.New("invalid iCal time format")
	ErrInvalidTimeValue  = errors.New("invalid time value")
	ErrLocalTimeRequired = errors.New("local time required; UTC designator Z is not allowed")
)

// parseDateParts parses YYYYMMDD from the first 8 characters of value.
// Callers must ensure value is long enough.
func parseDateParts(value string) (year, month, day int, err error) {
	year, err = strconv.Atoi(value[0:4])
	if err != nil {
		return 0, 0, 0, ErrInvalidTimeFormat
	}

	month, err = strconv.Atoi(value[4:6])
	if err != nil {
		return 0, 0, 0, ErrInvalidTimeFormat
	}
	if month < 1 || month > 12 {
		return 0, 0, 0, ErrInvalidTimeValue
	}

	day, err = strconv.Atoi(value[6:8])
	if err != nil {
		return 0, 0, 0, ErrInvalidTimeFormat
	}
	if day < 1 || day > 31 {
		return 0, 0, 0, ErrInvalidTimeValue
	}
	return year, month, day, nil
}

// parseIcalDate parses an iCal DATE value (YYYYMMDD) as midnight UTC.
// https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.4
func parseIcalDate(value string) (time.Time, error) {
	if len(value) != 8 {
		return time.Time{}, ErrInvalidTimeFormat
	}

	year, month, day, err := parseDateParts(value)
	if err != nil {
		return time.Time{}, err
	}

	t := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	if t.Day() != day {
		return time.Time{}, ErrInvalidTimeValue
	}
	return t, nil
}

// ParseIcalTime parses an iCal datetime string.
// Supports both UTC format (YYYYMMDDTHHMMSSZ) and floating time format (YYYYMMDDTHHMMSS).
// This manual implementation is faster than time.Parse for the fixed iCal format.
func ParseIcalTime(value string) (time.Time, error) {
	length := len(value)
	if length != 15 && length != 16 {
		return time.Time{}, ErrInvalidTimeFormat
	}

	if length == 16 {
		if value[15] != 'Z' {
			return time.Time{}, ErrInvalidTimeFormat
		}
	}

	year, month, day, err := parseDateParts(value)
	if err != nil {
		return time.Time{}, err
	}

	// Check for 'T' separator (position 8)
	if value[8] != 'T' {
		return time.Time{}, ErrInvalidTimeFormat
	}

	// Parse hour (positions 9-10)
	hour, err := strconv.Atoi(value[9:11])
	if err != nil {
		return time.Time{}, ErrInvalidTimeFormat
	}
	if hour < 0 || hour > 23 {
		return time.Time{}, ErrInvalidTimeValue
	}

	// Parse minute (positions 11-12)
	minute, err := strconv.Atoi(value[11:13])
	if err != nil {
		return time.Time{}, ErrInvalidTimeFormat
	}
	if minute < 0 || minute > 59 {
		return time.Time{}, ErrInvalidTimeValue
	}

	// Parse second (positions 13-14)
	second, err := strconv.Atoi(value[13:15])
	if err != nil {
		return time.Time{}, ErrInvalidTimeFormat
	}
	if second < 0 || second > 59 {
		return time.Time{}, ErrInvalidTimeValue
	}

	t := time.Date(year, time.Month(month), day, hour, minute, second, 0, time.UTC)
	// validate that day does not overflow (ie: february 31st would overflow to march 3rd)
	if t.Day() != day {
		return time.Time{}, ErrInvalidTimeValue
	}
	// All times are returned in UTC (floating times are treated as UTC per iCal spec)
	return t, nil
}

const valueTypeDate = "DATE"

// ParseIcalTimeOrDate parses a DATE-TIME value, or a DATE value when valueType is "DATE".
// valueType comparison is case-insensitive (RFC 5545 parameter values).
// When valueType is empty or "DATE-TIME", DATE-TIME parsing is used.
func ParseIcalTimeOrDate(value, valueType string) (time.Time, error) {
	if strings.EqualFold(valueType, valueTypeDate) {
		return parseIcalDate(value)
	}
	return ParseIcalTime(value)
}

// ParseIcalLocalTime parses an iCal datetime string that must be local wall time
// (FORM #1: DATE WITH LOCAL TIME). UTC values ending in "Z" are rejected.
func ParseIcalLocalTime(value string) (time.Time, error) {
	if len(value) > 0 && value[len(value)-1] == 'Z' {
		return time.Time{}, ErrLocalTimeRequired
	}
	return ParseIcalTime(value)
}
