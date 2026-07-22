// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package model

import (
	"net/url"
)

// Class is the CLASS property value shared by VEVENT, VTODO, and VJOURNAL.
// See: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.1.3
type Class string

const (
	ClassPublic       Class = "PUBLIC"
	ClassPrivate      Class = "PRIVATE"
	ClassConfidential Class = "CONFIDENTIAL"
)

// Transp is the TRANSP property value shared by VEVENT and VTODO.
// See: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.2.7
type Transp string

const (
	TranspTransparent Transp = "TRANSPARENT"
	TranspOpaque      Transp = "OPAQUE"
)

// Organizer represents an ORGANIZER component in the iCalendar format, used in VEVENT, VTODO, and VJOURNAL
// for more information see https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.4.3
type Organizer struct {
	// denoted by CN
	// See: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.2
	CommonName string
	// Note: Any Valid URI
	// See: https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.3
	CalAddress *url.URL

	// denoted by DIR
	// A directory entry reference
	// See: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.6
	Directory *url.URL

	// denoted by SENT-BY
	// See https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.18
	SentBy *url.URL

	// denoted by LANGUAGE
	// https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.10
	// no validation is done on the string at this time, but it is intended to be a valid tag under RFC5646
	// See: https://datatracker.ietf.org/doc/html/rfc5646
	Language string

	OtherParams map[string]string
}

// Attendee represents an ATTENDEE property in the iCalendar format.
// Used in VEVENT, VTODO, VJOURNAL, VFREEBUSY, and VALARM.
// See: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.4.1
type Attendee struct {
	// CalAddress is the attendee's calendar address (property value).
	CalAddress *url.URL

	// CommonName is the CN parameter.
	// See: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.2
	CommonName string

	// Directory is the DIR parameter.
	// See: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.6
	Directory *url.URL

	// SentBy is the SENT-BY parameter.
	// See: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.18
	SentBy *url.URL

	// Language is the LANGUAGE parameter.
	// See: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.10
	Language string

	// CUType is the CUTYPE parameter.
	// See: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.3
	CUType string

	// Role is the ROLE parameter.
	// See: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.16
	Role string

	// PartStat is the PARTSTAT parameter.
	// See: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.12
	PartStat string

	// RSVP is the RSVP parameter.
	// See: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.17
	RSVP bool

	// Member is the MEMBER parameter (one or more calendar addresses).
	// See: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.11
	Member []*url.URL

	// DelegatedTo is the DELEGATED-TO parameter (one or more calendar addresses).
	// See: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.5
	DelegatedTo []*url.URL

	// DelegatedFrom is the DELEGATED-FROM parameter (one or more calendar addresses).
	// See: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.4
	DelegatedFrom []*url.URL

	// OtherParams contains any other parameters on the ATTENDEE property.
	OtherParams map[string]string
}

// Attachment represents an ATTACH property in the iCalendar format.
// This property provides the capability to associate a document object with a calendar component.
// It can be specified as a URI pointing to a resource or as inline binary encoded content.
// For more information see https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.1.1.
type Attachment struct {
	// Value is the VALUE parameter, which can be "URI" (default) or "BINARY".
	// When Value is "BINARY", Encoding must be "BASE64".
	Value AttachValue

	// Encoding is the ENCODING parameter, which is "BASE64" for binary attachments.
	// This will be empty for URI attachments.
	Encoding string

	// URI is the URI pointing to the attachment resource.
	// This is set when the attachment is specified as a URI (the default value type).
	// This will be nil for binary attachments.
	URI *url.URL

	// Binary is the base64-encoded string for inline binary encoded content.
	// This is set when the attachment has ENCODING=BASE64 and VALUE=BINARY parameters.
	// This will be empty for URI attachments.
	Binary string

	// FormatType is the FMTTYPE parameter value.
	// This is optional for URI attachments but recommended for binary attachments.
	// It specifies the media type of the resource (e.g., "application/postscript").
	FormatType string

	// OtherParams contains any other parameters that may be specified on the ATTACH property.
	// This includes IANA and non-standard parameters.
	OtherParams map[string]string
}
