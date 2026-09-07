// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package model

// Property parameter names as defined by RFC 5545.
// Parameter names are an open set (IANA tokens and x-params are allowed),
// so these are untyped string constants rather than a closed enum.
// https://datatracker.ietf.org/doc/html/rfc5545#section-3.2
const (
	ParamValue         = "VALUE"
	ParamEncoding      = "ENCODING"
	ParamFmtType       = "FMTTYPE"
	ParamCN            = "CN"
	ParamDir           = "DIR"
	ParamLanguage      = "LANGUAGE"
	ParamSentBy        = "SENT-BY"
	ParamCUType        = "CUTYPE"
	ParamRole          = "ROLE"
	ParamPartStat      = "PARTSTAT"
	ParamRSVP          = "RSVP"
	ParamMember        = "MEMBER"
	ParamDelegatedTo   = "DELEGATED-TO"
	ParamDelegatedFrom = "DELEGATED-FROM"
	ParamTZID          = "TZID"
	ParamAltrep        = "ALTREP"
	ParamRange         = "RANGE"
	ParamRelated       = "RELATED"
	ParamRelType       = "RELTYPE"
	ParamFBType        = "FBTYPE"
)

// Boolean parameter values used by RSVP and similar parameters.
const (
	ParamTrue  = "TRUE"
	ParamFalse = "FALSE"
)

// AttachValue represents the closed set of values the VALUE parameter can take
// on an ATTACH property: "URI" (default) or "BINARY".
// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.1.1
type AttachValue string

const (
	AttachValueURI    AttachValue = "URI"
	AttachValueBinary AttachValue = "BINARY"
)

// EncodingBase64 is the ENCODING parameter value for base64-encoded binary
// attachment content. It is the only encoding this parser handles.
// https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.1
const EncodingBase64 = "BASE64"
