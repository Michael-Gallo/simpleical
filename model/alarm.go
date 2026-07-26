// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package model

import (
	"time"
)

// AlarmAction represents the possible values for a VALARM's ACTION field.
// actionvalue also admits iana-token and x-name. Recognized values are stored
// in canonical uppercase form. RFC 5545 requires applications to ignore alarms
// whose action they do not recognize, so the parser discards those alarms and
// only the three constants below reach a parsed Alarm.
// See: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.6.1
type AlarmAction string

const (
	AlarmActionAudio   AlarmAction = "AUDIO"
	AlarmActionDisplay AlarmAction = "DISPLAY"
	AlarmActionEmail   AlarmAction = "EMAIL"
)

// TriggerRelated is the RELATED parameter on TRIGGER (START or END).
// https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.14
type TriggerRelated string

const (
	TriggerRelatedStart TriggerRelated = "START"
	TriggerRelatedEnd   TriggerRelated = "END"
)

// Trigger specifies when an alarm will trigger.
// Exactly one of Duration or Absolute is set.
// Related defaults to START when unset.
// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.6.3
type Trigger struct {
	Duration *time.Duration // set for DURATION form
	Absolute *DateTime      // set for DATE-TIME form
	Related  TriggerRelated // RELATED param, default START
}

// Alarm represents a VALARM component in the iCalendar format.
// A VALARM is a grouping of component properties that defines an alarm.
// VALARM components are sub-components of VEVENT and VTODOs
// https://datatracker.ietf.org/doc/html/rfc5545#section-3.6.6
type Alarm struct {
	// REQUIRED, MUST NOT occur more than once
	// Defines the action to be invoked when an alarm is triggered.
	// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.6.1
	Action AlarmAction

	// REQUIRED, MUST NOT occur more than once
	// Specifies when an alarm will trigger.
	// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.6.3
	Trigger Trigger

	// OPTIONAL, MAY occur more than once for EMAIL; at most one for AUDIO.
	// Provides the capability to associate a document object with an alarm.
	// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.1.1
	Attach []Attachment

	// OPTIONAL, MUST NOT occur more than once.
	// Specifies a positive duration of time for repeating alarms.
	// Must co-occur with Repeat (both present or both absent).
	// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.6.5
	Duration time.Duration

	// OPTIONAL, MUST NOT occur more than once (for DISPLAY and EMAIL actions)
	// Provides a more complete description of the alarm than that provided by the SUMMARY property.
	// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.1.5
	Description string

	// OPTIONAL, MUST NOT occur more than once.
	// Defines the number of times the alarm should be repeated, including zero.
	// Nil means the property was absent; must co-occur with Duration.
	// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.6.4
	Repeat *int

	// OPTIONAL, MUST NOT occur more than once (for EMAIL action)
	// Defines a short summary or subject for the alarm.
	// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.1.12
	Summary string

	// OPTIONAL, MAY occur more than once (for EMAIL action, at least one required)
	// Specifies the participants that are invited to the alarm.
	// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.4.1
	Attendees []Attendee

	// OPTIONAL, MAY occur more than once
	// A Non-Standard Property. Can be represented by any name with a X-prefix.
	// Parameters and repeats are preserved.
	// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.8.2
	XProp []ExtensionProperty

	// OPTIONAL, MAY occur more than once
	// An unrecognized IANA-style property name (non X- prefix).
	// Parameters and repeats are preserved. No IANA-registration validation.
	// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.8.1
	IANAProp []ExtensionProperty
}
