# Date-Time Start

Source: RFC 5545, Section 3.8.2.4

3.8.2.4.  Date-Time Start

   Property Name:  DTSTART

   Purpose:  This property specifies when the calendar component begins.

   Value Type:  The default value type is DATE-TIME.  The time value
      MUST be one of the forms defined for the DATE-TIME value type.
      The value type can be set to a DATE value type.

   Property Parameters:  IANA, non-standard, value data type, and time
      zone identifier property parameters can be specified on this
      property.

   Conformance:  This property can be specified once in the "VEVENT",
      "VTODO", or "VFREEBUSY" calendar components as well as in the

      "STANDARD" and "DAYLIGHT" sub-components.  This property is
      REQUIRED in all types of recurring calendar components that
      specify the "RRULE" property.  This property is also REQUIRED in
      "VEVENT" calendar components contained in iCalendar objects that
      don't specify the "METHOD" property.

   Description:  Within the "VEVENT" calendar component, this property
      defines the start date and time for the event.

      Within the "VFREEBUSY" calendar component, this property defines
      the start date and time for the free or busy time information.
      The time MUST be specified in UTC time.

      Within the "STANDARD" and "DAYLIGHT" sub-components, this property
      defines the effective start date and time for a time zone
      specification.  This property is REQUIRED within each "STANDARD"
      and "DAYLIGHT" sub-components included in "VTIMEZONE" calendar
      components and MUST be specified as a date with local time without
      the "TZID" property parameter.

   Format Definition:  This property is defined by the following
      notation:

       dtstart    = "DTSTART" dtstparam ":" dtstval CRLF

       dtstparam  = *(
                  ;
                  ; The following are OPTIONAL,
                  ; but MUST NOT occur more than once.
                  ;
                  (";" "VALUE" "=" ("DATE-TIME" / "DATE")) /
                  (";" tzidparam) /
                  ;
                  ; The following is OPTIONAL,
                  ; and MAY occur more than once.
                  ;
                  (";" other-param)
                  ;
                  )

       dtstval    = date-time / date
       ;Value MUST match value type

   Example:  The following is an example of this property:

       DTSTART:19980118T073000Z
