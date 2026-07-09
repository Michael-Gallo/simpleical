// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package model

import (
	"time"

	"github.com/michael-gallo/simpleical/rrule"
)

// TodoStatus represents the possible values for a VTODO's STATUS field.
// See: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.1.11
type TodoStatus string

const (
	TodoStatusNeedsAction TodoStatus = "NEEDS-ACTION"
	TodoStatusCompleted   TodoStatus = "COMPLETED"
	TodoStatusInProcess   TodoStatus = "IN-PROCESS"
	TodoStatusCancelled   TodoStatus = "CANCELLED" //nolint:misspell // iCalendar property name, not a typo
)

// TodoClass represents the possible values for a VTODO's CLASS field.
// See: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.1.3
type TodoClass string

const (
	TodoClassPublic       TodoClass = "PUBLIC"
	TodoClassPrivate      TodoClass = "PRIVATE"
	TodoClassConfidential TodoClass = "CONFIDENTIAL"
)

// TodoTransp represents the possible values for a VTODO's TRANSP field.
// See: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.2.7
type TodoTransp string

const (
	TodoTranspTransparent TodoTransp = "TRANSPARENT"
	TodoTranspOpaque      TodoTransp = "OPAQUE"
)

// Todo represents a VTODO component in the iCalendar format.
// A VTODO is a grouping of component properties that describe a to-do.
// https://datatracker.ietf.org/doc/html/rfc5545#section-3.6.2
type Todo struct {
	// REQUIRED, MUST NOT occur more than once
	// a DTSTAMP property defines the date and time that the instance of the calendar component was created.
	// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.7.2
	DTStamp time.Time

	// REQUIRED, MUST NOT occur more than once
	// The unique identifier for the todo.
	// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.4.7
	UID string

	// OPTIONAL, MUST NOT occur more than once
	// Access Classification for the calendar component.
	// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.1.3
	Class TodoClass

	// OPTIONAL, MUST NOT occur more than once
	// Specifies the date and time that a to-do was actually completed.
	// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.2.1
	Completed time.Time

	// OPTIONAL, MUST NOT occur more than once
	// Specifies the date and time that the calendar information was created.
	// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.7.1
	Created time.Time

	// OPTIONAL, MUST NOT occur more than once
	// Used to capture lengthy textual descriptions associated with the activity.
	// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.1.5
	Description string

	// OPTIONAL, MUST NOT occur more than once
	// Specifies when the calendar component begins.
	// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.2.4
	DTStart time.Time

	// OPTIONAL, MUST NOT occur more than once
	// Specifies the date and time that a to-do is expected to be completed.
	// Either DUE or DURATION may be specified in a VTODO, but not both.
	// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.2.3
	Due time.Time

	// OPTIONAL, MUST NOT occur more than once
	// Specifies a positive duration of time.
	// Either DUE or DURATION may be specified in a VTODO, but not both.
	// https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.6
	Duration time.Duration

	// OPTIONAL, MUST NOT occur more than once
	// Geo specifies the latitude and longitude of the activity specified by a calendar component.
	// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.1.6
	Geo []float64

	// OPTIONAL, MUST NOT occur more than once
	// Specifies the date and time that the information associated with the calendar component was last revised.
	// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.7.3
	LastModified time.Time

	// OPTIONAL, MUST NOT occur more than once
	// The location where the activity takes place.
	// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.1.7
	Location string

	// OPTIONAL, MUST NOT occur more than once
	// The organizer of the activity.
	// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.4.3
	Organizer *Organizer

	// OPTIONAL, MUST NOT occur more than once
	// Specifies the percentage completion of a to-do.
	// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.1.8
	PercentComplete int

	// OPTIONAL, MUST NOT occur more than once
	// Specifies the priority for a calendar component.
	// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.1.9
	Priority int

	// OPTIONAL, MUST NOT occur more than once
	// Specifies the revision sequence number of the calendar component within a sequence of revisions.
	// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.7.4
	RecurrenceID time.Time

	// OPTIONAL, MUST NOT occur more than once
	// Specifies the revision sequence number of the calendar component within a sequence of revisions.
	// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.7.4
	Sequence int

	// OPTIONAL, MUST NOT occur more than once
	// Defines the overall status or confirmation for the calendar component.
	// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.1.11
	Status TodoStatus

	// OPTIONAL, MUST NOT occur more than once
	// A short, one-line summary about the activity.
	// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.1.12
	Summary string

	// OPTIONAL, MUST NOT occur more than once
	// The time transparency for the activity.
	// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.2.7
	Transp TodoTransp

	// OPTIONAL, MUST NOT occur more than once
	// Specifies a URL associated with the activity.
	// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.4.6
	URL string

	// OPTIONAL, SHOULD NOT occur more than once
	// Specifies the recurrence rule for the todo.
	// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.5.3
	RRule *rrule.RRule

	// OPTIONAL, MAY occur more than once
	// Provides the capability to associate a document object with a calendar component.
	// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.1.1
	Attach []Attachment

	// OPTIONAL, MAY occur more than once
	// Specifies the participants that are invited to the activity.
	// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.4.1
	Attendees []Attendee

	// OPTIONAL, MAY occur more than once
	// Specifies the categories that the calendar component belongs to.
	// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.1.2
	Categories []string

	// OPTIONAL, MAY occur more than once
	// Specifies non-processing information intended to provide a comment to the calendar user.
	// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.1.4
	Comment []string

	// OPTIONAL, MAY occur more than once
	// Specifies the contact information for the activity.
	// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.4.2
	Contacts []string

	// OPTIONAL, MAY occur more than once
	// Specifies the list of date/time exceptions for a recurring calendar component.
	// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.5.1
	ExceptionDates []time.Time

	// OPTIONAL, MAY occur more than once
	// Specifies the status code returned for a scheduling request.
	// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.8.3
	RequestStatus []string

	// OPTIONAL, MAY occur more than once
	// Specifies a relationship or reference between one calendar component and another.
	// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.4.5
	Related []string

	// OPTIONAL, MAY occur more than once
	// Specifies equipment or resources anticipated for an activity.
	// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.1.10
	Resources []string

	// OPTIONAL, MAY occur more than once
	// Specifies the list of date/time values for recurring activities.
	// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.5.2
	Rdate []time.Time

	// OPTIONAL, MAY occur more than once
	// A Non-Standard Property. Can be represented by any name with a X-prefix.
	// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.8.2
	XProp map[string]string

	// OPTIONAL, MAY occur more than once
	// An IANA registered property name.
	// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.8.1
	IANAProp map[string]string

	// OPTIONAL, MAY occur more than once
	// Sub-components: VALARM
	Alarms []Alarm
}
