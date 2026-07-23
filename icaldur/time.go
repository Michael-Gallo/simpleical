// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

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
	ErrInvalidUTCOffset  = errors.New("invalid UTC-OFFSET format")
	ErrTZIDWithUTC       = errors.New("TZID must not be applied to UTC DATE-TIME values")
)

// Form identifies the RFC 5545 DATE / DATE-TIME form of a parsed value.
type Form int

const (
	FormDate Form = iota + 1
	FormFloating
	FormUTC
)

// Temporal is a parsed DATE or DATE-TIME without TZID binding.
type Temporal struct {
	Form Form
	Time time.Time
}

const valueTypeDate = "DATE"

// parseDateParts parses YYYYMMDD from the first 8 characters of value.
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

func parseIcalDate(value string) (Temporal, error) {
	if len(value) != 8 {
		return Temporal{}, ErrInvalidTimeFormat
	}
	year, month, day, err := parseDateParts(value)
	if err != nil {
		return Temporal{}, err
	}
	t := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	if t.Day() != day {
		return Temporal{}, ErrInvalidTimeValue
	}
	return Temporal{Form: FormDate, Time: t}, nil
}

// ParseIcalTime parses an iCal DATE-TIME string into a Temporal.
// Supports UTC (…Z) and floating (no Z) forms. Leap second 60 is accepted and
// stored as 59 per RFC 5545 SHOULD guidance.
func ParseIcalTime(value string) (time.Time, error) {
	t, err := ParseTemporalDateTime(value)
	if err != nil {
		return time.Time{}, err
	}
	return t.Time, nil
}

// ParseTemporalDateTime parses a DATE-TIME value, distinguishing UTC vs floating.
func ParseTemporalDateTime(value string) (Temporal, error) {
	length := len(value)
	if length != 15 && length != 16 {
		return Temporal{}, ErrInvalidTimeFormat
	}
	isUTC := false
	if length == 16 {
		if value[15] != 'Z' {
			return Temporal{}, ErrInvalidTimeFormat
		}
		isUTC = true
	}

	year, month, day, err := parseDateParts(value)
	if err != nil {
		return Temporal{}, err
	}
	if value[8] != 'T' {
		return Temporal{}, ErrInvalidTimeFormat
	}

	hour, err := strconv.Atoi(value[9:11])
	if err != nil {
		return Temporal{}, ErrInvalidTimeFormat
	}
	if hour < 0 || hour > 23 {
		return Temporal{}, ErrInvalidTimeValue
	}
	minute, err := strconv.Atoi(value[11:13])
	if err != nil {
		return Temporal{}, ErrInvalidTimeFormat
	}
	if minute < 0 || minute > 59 {
		return Temporal{}, ErrInvalidTimeValue
	}
	second, err := strconv.Atoi(value[13:15])
	if err != nil {
		return Temporal{}, ErrInvalidTimeFormat
	}
	// RFC 5545 allows second 60 for a positive leap second; map to 59.
	if second < 0 || second > 60 {
		return Temporal{}, ErrInvalidTimeValue
	}
	if second == 60 {
		second = 59
	}

	t := time.Date(year, time.Month(month), day, hour, minute, second, 0, time.UTC)
	if t.Day() != day {
		return Temporal{}, ErrInvalidTimeValue
	}
	form := FormFloating
	if isUTC {
		form = FormUTC
	}
	return Temporal{Form: form, Time: t}, nil
}

// ParseIcalTimeOrDate parses a DATE-TIME value, or a DATE value when valueType is "DATE".
func ParseIcalTimeOrDate(value, valueType string) (time.Time, error) {
	t, err := ParseTemporal(value, valueType)
	if err != nil {
		return time.Time{}, err
	}
	return t.Time, nil
}

// ParseTemporal parses a DATE or DATE-TIME value.
func ParseTemporal(value, valueType string) (Temporal, error) {
	if strings.EqualFold(valueType, valueTypeDate) {
		return parseIcalDate(value)
	}
	return ParseTemporalDateTime(value)
}

// ParseIcalLocalTime parses an iCal datetime that must be local wall time (no Z).
func ParseIcalLocalTime(value string) (time.Time, error) {
	t, err := ParseTemporalDateTime(value)
	if err != nil {
		return time.Time{}, err
	}
	if t.Form == FormUTC {
		return time.Time{}, ErrLocalTimeRequired
	}
	return t.Time, nil
}

// ParseUTCOffset parses a UTC-OFFSET value: [+/-]HHMM[SS].
// https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.14
func ParseUTCOffset(value string) (string, error) {
	if len(value) != 5 && len(value) != 7 {
		return "", ErrInvalidUTCOffset
	}
	if value[0] != '+' && value[0] != '-' {
		return "", ErrInvalidUTCOffset
	}
	for i := 1; i < len(value); i++ {
		if value[i] < '0' || value[i] > '9' {
			return "", ErrInvalidUTCOffset
		}
	}
	// RFC 5545 forbids negative-zero offsets.
	if value == "-0000" || value == "-000000" {
		return "", ErrInvalidUTCOffset
	}
	hour, err := strconv.Atoi(value[1:3])
	if err != nil || hour > 23 {
		return "", ErrInvalidUTCOffset
	}
	minute, err := strconv.Atoi(value[3:5])
	if err != nil || minute > 59 {
		return "", ErrInvalidUTCOffset
	}
	if len(value) == 7 {
		sec, err := strconv.Atoi(value[5:7])
		if err != nil || sec > 59 {
			return "", ErrInvalidUTCOffset
		}
	}
	return value, nil
}
