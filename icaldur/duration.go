// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package icaldur

import (
	"errors"
	"strconv"
	"strings"
	"time"
	"unicode"
)

var (
	errEmpty            = errors.New("empty duration")
	errBadPrefix        = errors.New("duration must start with P (optionally preceded by + or -)")
	errUnexpectedChar   = errors.New("unexpected character")
	errMissingUnit      = errors.New("missing unit after number")
	errMixedWeeks       = errors.New("weeks form (PnW) cannot be mixed with other components")
	errTimeWithoutT     = errors.New("time components require a preceding 'T'")
	errDuplicateUnit    = errors.New("duplicate time unit")
	errDuplicateT       = errors.New("duplicate 'T' marker")
	errNoComponents     = errors.New("duration must have at least one component")
	errTWithoutTimePart = errors.New("'T' marker must be followed by time components")
)

// ParseICalDuration parses an iCalendar (RFC 5545 section 3.3.6) duration string into a time.Duration.
// An optional leading '+' or '-' sign indicates a positive or negative duration.
// Accepted forms are a weeks-only `PnW` (must be the only component) or `P[nD][T[nH][nM][nS]]`
// with date units D (days) and time units H (hours), M (minutes), and S (seconds) after a `T` marker.
// Invalid inputs are rejected, including mixed weeks with other components, duplicate units,
// multiple `T` markers, and time components without a preceding `T`.
func ParseICalDuration(s string) (time.Duration, error) {
	if len(s) == 0 {
		return 0, errEmpty
	}

	// Trim spaces (optional)
	start, end := 0, len(s)
	for start < end && unicode.IsSpace(rune(s[start])) {
		start++
	}
	for end > start && unicode.IsSpace(rune(s[end-1])) {
		end--
	}
	if start == end {
		return 0, errEmpty
	}
	s = s[start:end]

	sign := int64(1)
	i := 0

	// Optional sign
	switch s[i] {
	case '+':
		i++
	case '-':
		sign = -1
		i++
	}

	// Must start with 'P'
	if i >= len(s) || s[i] != 'P' {
		return 0, errBadPrefix
	}
	i++

	var (
		inTime                     bool
		dur                        int64 // nanoseconds
		usedD, usedH, usedM, usedS bool
	)

	// Helper to read a positive integer
	readInt := func() (int64, bool) {
		if i >= len(s) || !unicode.IsDigit(rune(s[i])) {
			return 0, false
		}
		start := i
		for i < len(s) && unicode.IsDigit(rune(s[i])) {
			i++
		}
		v, err := strconv.ParseInt(s[start:i], 10, 64)
		if err != nil {
			return 0, false
		}
		return v, true
	}

	// Special-case weeks: PnW and nothing else
	// Detect if there's a 'W' anywhere; if present, it must be the only unit
	if rel := strings.IndexByte(s[i:], 'W'); rel != -1 {
		wpos := i + rel
		// Ensure there are only digits between i and wpos, and nothing after W
		numStart := i
		if numStart >= wpos {
			return 0, errMissingUnit
		}
		for j := numStart; j < wpos; j++ {
			if !unicode.IsDigit(rune(s[j])) {
				return 0, errUnexpectedChar
			}
		}
		if wpos != len(s)-1 {
			return 0, errMixedWeeks
		}
		v, err := strconv.ParseInt(s[numStart:wpos], 10, 64)
		if err != nil {
			return 0, err
		}
		dur = v * 7 * 24 * int64(time.Hour)
		return time.Duration(sign * dur), nil
	}

	// Otherwise parse date/time components: P[nD][T[nH][nM][nS]]
	var (
		seenT             bool
		seenComponent     bool
		seenTimeComponent bool
	)

	for i < len(s) {
		if s[i] == 'T' {
			if seenT {
				return 0, errDuplicateT
			}
			seenT = true
			inTime = true
			i++
			continue
		}

		v, ok := readInt()
		if !ok {
			return 0, errMissingUnit
		}
		if i >= len(s) {
			return 0, errMissingUnit
		}
		unit := s[i]
		i++

		switch unit {
		case 'D':
			if inTime {
				return 0, errUnexpectedChar
			}
			if usedD {
				return 0, errDuplicateUnit
			}
			usedD = true
			seenComponent = true
			dur += v * 24 * int64(time.Hour)
		case 'H':
			if !inTime {
				return 0, errTimeWithoutT
			}
			if usedH {
				return 0, errDuplicateUnit
			}
			usedH = true
			seenComponent = true
			seenTimeComponent = true
			dur += v * int64(time.Hour)
		case 'M':
			if !inTime {
				return 0, errTimeWithoutT
			}
			if usedM {
				return 0, errDuplicateUnit
			}
			usedM = true
			seenComponent = true
			seenTimeComponent = true
			dur += v * int64(time.Minute)
		case 'S':
			if !inTime {
				return 0, errTimeWithoutT
			}
			if usedS {
				return 0, errDuplicateUnit
			}
			usedS = true
			seenComponent = true
			seenTimeComponent = true
			dur += v * int64(time.Second)
		default:
			return 0, errUnexpectedChar
		}
	}

	// Reject empty durations like "P".
	if !seenComponent {
		return 0, errNoComponents
	}

	// Reject trailing 'T' after date components without time values (e.g., "P1DT").
	if seenT && !seenTimeComponent {
		return 0, errTWithoutTimePart
	}

	return time.Duration(sign * dur), nil
}
