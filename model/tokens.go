// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package model

// SectionToken represents the names of the top level components in a VCALENDAR
// https://datatracker.ietf.org/doc/html/rfc5545#section-3.6
type SectionToken string

const (
	SectionTokenVCalendar SectionToken = "VCALENDAR"
	SectionTokenVEvent    SectionToken = "VEVENT"
	SectionTokenVTodo     SectionToken = "VTODO"
	SectionTokenVJournal  SectionToken = "VJOURNAL"
	SectionTokenVTimezone SectionToken = "VTIMEZONE"
	SectionTokenVFreebusy SectionToken = "VFREEBUSY"
	SectionTokenVAlarm    SectionToken = "VALARM"
	SectionTokenVStandard SectionToken = "STANDARD"
	SectionTokenVDaylight SectionToken = "DAYLIGHT"
)

// Property names as defined by RFC 5545.
// Property names are shared across components, so these are untyped string
// constants rather than per-component enums.
// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8
const (
	PropAction          = "ACTION"
	PropAttach          = "ATTACH"
	PropAttendee        = "ATTENDEE"
	PropCategories      = "CATEGORIES"
	PropClass           = "CLASS"
	PropComment         = "COMMENT"
	PropCompleted       = "COMPLETED"
	PropContact         = "CONTACT"
	PropCreated         = "CREATED"
	PropDescription     = "DESCRIPTION"
	PropDTEnd           = "DTEND"
	PropDTStamp         = "DTSTAMP"
	PropDTStart         = "DTSTART"
	PropDue             = "DUE"
	PropDuration        = "DURATION"
	PropExDate          = "EXDATE"
	PropFreeBusy        = "FREEBUSY"
	PropGeo             = "GEO"
	PropLastModified    = "LAST-MODIFIED"
	PropLocation        = "LOCATION"
	PropOrganizer       = "ORGANIZER"
	PropPercentComplete = "PERCENT-COMPLETE"
	PropPriority        = "PRIORITY"
	PropRDate           = "RDATE"
	PropRecurrenceID    = "RECURRENCE-ID"
	PropRelatedTo       = "RELATED-TO"
	PropRepeat          = "REPEAT"
	PropRequestStatus   = "REQUEST-STATUS"
	PropResources       = "RESOURCES"
	PropRRule           = "RRULE"
	PropSequence        = "SEQUENCE"
	PropStatus          = "STATUS"
	PropSummary         = "SUMMARY"
	PropTransp          = "TRANSP"
	PropTrigger         = "TRIGGER"
	PropTZID            = "TZID"
	PropTZName          = "TZNAME"
	PropTZOffsetFrom    = "TZOFFSETFROM"
	PropTZOffsetTo      = "TZOFFSETTO"
	PropTZURL           = "TZURL"
	PropUID             = "UID"
	PropURL             = "URL"
)
