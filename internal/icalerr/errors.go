// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package icalerr

import "errors"

// Calendar-level errors.
var (
	ErrNoCalendarFound                   = errors.New("empty calendar sent")
	ErrInvalidCalendarFormatMissingBegin = errors.New("invalid calendar format: must start with BEGIN:VCALENDAR")
	ErrInvalidCalendarFormatMissingEnd   = errors.New("invalid calendar format: must end with END:VCALENDAR")
	ErrInvalidCalendarEmptyLine          = errors.New("invalid calendar format: must not contain empty lines")
	ErrNestedBeginVCalendar              = errors.New("invalid calendar format: BEGIN:VCALENDAR cannot appear inside a calendar")
	ErrContentAfterEndBlock              = errors.New("content after END:VCALENDAR")
	ErrTemplateInvalidEndBlock           = errors.New("invalid end block")
	ErrUnexpectedEndBlock                = errors.New("unexpected end block: no matching begin or wrong nesting")
	ErrUnexpectedBeginBlock              = errors.New("unexpected begin block: not inside expected parent")
	ErrMissingCalendarVersionProperty    = errors.New("calendar must have a VERSION property")
	ErrMissingCalendarProdIDProperty     = errors.New("calendar must have a PRODID property")
	ErrMissingCalendarComponent          = errors.New("calendar must contain at least one component")
	ErrCalendarPropertyAfterComponent    = errors.New("calendar property after component is not allowed")
	ErrDuplicateParameter                = errors.New("duplicate property parameter")
	ErrInvalidTextEscape                 = errors.New("invalid TEXT escape sequence")
	ErrComponentNotAllowedHere           = errors.New("component begin not allowed in current state")
)

// General parsing errors.
var (
	ErrInvalidPropertyLine       = errors.New("invalid property line in iCal data")
	ErrInvalidOrganizer          = errors.New("invalid organizer property")
	ErrInvalidAttendee           = errors.New("invalid attendee property")
	ErrInvalidRRule              = errors.New("invalid rrule property")
	ErrPropertyWhenNotInCalendar = errors.New("property found when not in a calendar")
)

// Event-specific errors.
var (
	ErrMissingEventUIDProperty     = errors.New("event must have a UID property")
	ErrMissingEventDTStampProperty = errors.New("event must have a DTSTAMP property")
	ErrMissingEventDTStartProperty = errors.New("event must have a DTSTART property if no METHOD property is present for the top level calendar")

	ErrInvalidDurationPropertyDtend = errors.New("invalid duration property in iCal Event: DTEND and DURATION are mutually exclusive")
	ErrDateDurationMustBeDayOrWeek  = errors.New("DURATION with DATE DTSTART must be day or week granularity")
	ErrMismatchedDateValueTypes     = errors.New("DTSTART and DTEND/DUE must use matching DATE or DATE-TIME value types")
	ErrPositiveDurationRequired     = errors.New("duration must be positive")
	ErrPeriodEndNotAfterStart       = errors.New("PERIOD end must be after start")
	ErrInvalidPriority              = errors.New("PRIORITY must be an integer from 0 to 9")
	ErrInvalidPercentComplete       = errors.New("PERCENT-COMPLETE must be an integer from 0 to 100")
	ErrInvalidEnumValue             = errors.New("invalid enumerated property value")
	ErrUTCValueRequired             = errors.New("property value must be in UTC form")
	ErrUnknownTZID                  = errors.New("TZID parameter does not reference a VTIMEZONE in this calendar")
	ErrDuplicateTimezoneTZID        = errors.New("duplicate VTIMEZONE TZID in calendar")

	ErrInvalidGeoProperty          = errors.New("invalid event property in iCal Event: GEO must be two floats separated by a semicolon")
	ErrInvalidGeoPropertyLatitude  = errors.New("invalid latitude in iCal Event: GEO must be a float")
	ErrInvalidGeoPropertyLongitude = errors.New("invalid longitude in iCal Event: GEO must be a float")
	ErrGeoLatitudeOutOfRange       = errors.New("GEO latitude must be between -90 and 90")
	ErrGeoLongitudeOutOfRange      = errors.New("GEO longitude must be between -180 and 180")
)

// Todo-specific errors.
var (
	ErrMissingTodoUIDProperty = errors.New("todo must have a UID property")

	ErrMissingTodoDTStampProperty = errors.New("todo must have a DTSTAMP property")

	ErrInvalidDurationPropertyDue = errors.New("invalid duration property in iCal Todo: DUE and DURATION are mutually exclusive")
	ErrDurationRequiresDTStart    = errors.New("invalid duration property in iCal Todo: DURATION requires DTSTART to be set")
)

// Journal-specific errors.
var (
	ErrMissingJournalUIDProperty = errors.New("journal must have a UID property")

	ErrMissingJournalDTStampProperty = errors.New("journal must have a DTSTAMP property")
)

// FreeBusy-specific errors.
var (
	ErrMissingFreeBusyUIDProperty = errors.New("freebusy must have a UID property")

	ErrInvalidFreeBusyFormat = errors.New("invalid FREEBUSY property format")

	ErrMissingFreeBusyDTStampProperty = errors.New("freebusy must have a DTSTAMP property")
)

// Timezone-specific errors.
var (
	ErrInvalidTimezoneProperty     = errors.New("invalid timezone property")
	ErrMissingTimezoneTZIDProperty = errors.New("timezone must have a TZID property")
	ErrMissingTimezoneObservance   = errors.New("timezone must have at least one STANDARD or DAYLIGHT sub-component")
	ErrMissingObservanceDTStart    = errors.New("STANDARD/DAYLIGHT must have a DTSTART property")
	ErrMissingObservanceOffsetFrom = errors.New("STANDARD/DAYLIGHT must have a TZOFFSETFROM property")
	ErrMissingObservanceOffsetTo   = errors.New("STANDARD/DAYLIGHT must have a TZOFFSETTO property")
	// ErrTimezoneLocalTimeRequired is returned when DTSTART or RDATE in a
	// STANDARD/DAYLIGHT sub-component uses a UTC timestamp (trailing "Z").
	// RFC 5545 requires these values to be local wall time.
	ErrTimezoneLocalTimeRequired = errors.New("timezone STANDARD/DAYLIGHT DTSTART and RDATE must be specified as local time")
)

// Alarm-specific errors.
var (
	ErrUnknownAlarmAction = errors.New("unknown alarm action")

	ErrMissingAlarmActionProperty = errors.New("alarm must have an ACTION property")

	ErrMissingAlarmTriggerProperty = errors.New("alarm must have a TRIGGER property")

	ErrMissingAlarmDescriptionForDisplay = errors.New("DISPLAY alarm must have a DESCRIPTION property")

	ErrMissingAlarmDescriptionForEmail = errors.New("EMAIL alarm must have a DESCRIPTION property")

	ErrMissingAlarmSummaryForEmail = errors.New("EMAIL alarm must have a SUMMARY property")

	ErrMissingAlarmAttendeesForEmail = errors.New("EMAIL alarm must have at least one ATTENDEE property")

	ErrAlarmAttachTooManyForAudio  = errors.New("AUDIO alarm must not have more than one ATTACH property")
	ErrAlarmDurationRepeatCoupling = errors.New("DURATION and REPEAT must both be present or both absent")
	ErrInvalidAlarmTrigger         = errors.New("invalid TRIGGER property")
	ErrTodoTranspNotAllowed        = errors.New("TRANSP is not allowed on VTODO")
)

// Property Setter errors.

const ErrDuplicatePropertyInComponentFormat = "%w: %s set twice in component %s"

var (
	ErrDuplicatePropertyInComponent = errors.New("duplicate property error")
	ErrParseErrorInComponent        = errors.New("parse error in component")
)
