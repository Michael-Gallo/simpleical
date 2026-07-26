// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package model

// RecurrenceDate is a single RDATE value: either a DATE/DATE-TIME or a PERIOD.
// Exactly one of DateTime or Period is set.
// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.5.2
type RecurrenceDate struct {
	DateTime *DateTime // set for DATE/DATE-TIME values
	Period   *Period   // set for PERIOD values
}
