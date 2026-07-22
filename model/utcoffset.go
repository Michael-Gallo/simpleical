// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package model

// UTCOffset is a validated UTC-OFFSET value (e.g. "-0500", "+0530", "+013015").
// The raw string is retained; parsing validates the format.
// https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.14
type UTCOffset string
