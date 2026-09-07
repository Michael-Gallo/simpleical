// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package ical

import (
	"strings"

	"github.com/michael-gallo/simpleical/internal/icalerr"
)

// unescapeText decodes RFC 5545 TEXT escapes: \\ \; \, \n \N.
// https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.11
func unescapeText(s string) (string, error) {
	if !strings.ContainsRune(s, '\\') {
		return s, nil
	}
	var b strings.Builder
	b.Grow(len(s))
	for i := 0; i < len(s); i++ {
		if s[i] != '\\' {
			b.WriteByte(s[i])
			continue
		}
		i++
		if i >= len(s) {
			return "", icalerr.ErrInvalidTextEscape
		}
		switch s[i] {
		case '\\', ';', ',':
			b.WriteByte(s[i])
		case 'n', 'N':
			b.WriteByte('\n')
		default:
			return "", icalerr.ErrInvalidTextEscape
		}
	}
	return b.String(), nil
}

// splitUnescapedComma splits a TEXT list on unescaped commas and decodes
// RFC 5545 TEXT escapes in the same pass. An escaped comma (\,) is data.
func splitUnescapedComma(s string) ([]string, error) {
	if s == "" {
		return nil, nil
	}
	var parts []string
	var b strings.Builder
	b.Grow(len(s))
	for i := 0; i < len(s); i++ {
		if s[i] == '\\' {
			i++
			if i >= len(s) {
				return nil, icalerr.ErrInvalidTextEscape
			}
			switch s[i] {
			case '\\', ';', ',':
				b.WriteByte(s[i])
			case 'n', 'N':
				b.WriteByte('\n')
			default:
				return nil, icalerr.ErrInvalidTextEscape
			}
			continue
		}
		if s[i] == ',' {
			parts = append(parts, b.String())
			b.Reset()
			continue
		}
		b.WriteByte(s[i])
	}
	parts = append(parts, b.String())
	return parts, nil
}
