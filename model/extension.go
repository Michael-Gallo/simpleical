// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package model

// ExtensionProperty holds an X- property or unrecognized IANA property.
// Names are stored uppercased. Parameters are preserved when present.
// Multiple occurrences of the same name are retained in parse order.
// There is no validation that an IANA property name is actually registered.
// See: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.8.1
// See: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.8.2
type ExtensionProperty struct {
	Name   string
	Value  string
	Params map[string]string
}
