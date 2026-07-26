// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package model

import "time"

// DateTimeForm identifies which RFC 5545 DATE / DATE-TIME form a value uses.
// https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.4
// https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.5
type DateTimeForm int

const (
	// DateTimeFormDate is a DATE value (YYYYMMDD).
	DateTimeFormDate DateTimeForm = iota + 1
	// DateTimeFormFloating is a DATE-TIME without UTC designator or TZID.
	DateTimeFormFloating
	// DateTimeFormUTC is a DATE-TIME with a trailing Z.
	DateTimeFormUTC
	// DateTimeFormLocalTZ is a DATE-TIME with a TZID parameter.
	DateTimeFormLocalTZ
)

// DateTime is an iCalendar DATE or DATE-TIME value that preserves form and TZID.
// Time holds the civil year/month/day(/hour/minute/second) in time.UTC location;
// Form and TZID convey the RFC semantics. Leap seconds are accepted on input and
// stored as second 59 per the RFC SHOULD guidance for implementations that do
// not support leap seconds.
type DateTime struct {
	Form DateTimeForm
	Time time.Time
	TZID string
}

// IsZero reports whether the DateTime is the zero value (unset).
func (d DateTime) IsZero() bool {
	return d.Form == 0 && d.Time.IsZero() && d.TZID == ""
}

// IsDate reports whether this value is a DATE (all-day) form.
func (d DateTime) IsDate() bool {
	return d.Form == DateTimeFormDate
}

// IsUTC reports whether this value is UTC DATE-TIME form.
func (d DateTime) IsUTC() bool {
	return d.Form == DateTimeFormUTC
}

// NewUTCDateTime builds a UTC DATE-TIME from a time.Time (civil fields used).
func NewUTCDateTime(t time.Time) DateTime {
	return DateTime{Form: DateTimeFormUTC, Time: t.UTC()}
}

// NewFloatingDateTime builds a floating DATE-TIME from a time.Time.
func NewFloatingDateTime(t time.Time) DateTime {
	return DateTime{Form: DateTimeFormFloating, Time: time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), time.UTC)}
}

// NewDate builds a DATE value from a time.Time (date portion only).
func NewDate(t time.Time) DateTime {
	return DateTime{Form: DateTimeFormDate, Time: time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC)}
}

// NewLocalTZDateTime builds a local-with-TZID DATE-TIME.
func NewLocalTZDateTime(t time.Time, tzid string) DateTime {
	return DateTime{
		Form: DateTimeFormLocalTZ,
		Time: time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), time.UTC),
		TZID: tzid,
	}
}
