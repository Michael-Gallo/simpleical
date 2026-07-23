// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package model

import "time"

// TextValue is a TEXT property value with common optional parameters.
// https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.11
type TextValue struct {
	Value    string
	Language string
	Altrep   string
}

// RelatedToValue is a RELATED-TO property value with optional RELTYPE.
// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.4.5
type RelatedToValue struct {
	Value   string
	RelType string
}

// Period is an iCalendar PERIOD value: either start/end or start/duration.
// Exactly one of End or Duration is set (Duration != 0 means start/duration form).
// https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.9
type Period struct {
	Start    DateTime
	End      DateTime
	Duration time.Duration
}
