# Time Zone Identifier

Source: RFC 5545, Section 3.8.3.1

3.8.3.1.  Time Zone Identifier

   Property Name:  TZID

   Purpose:  This property specifies the text value that uniquely
      identifies the "VTIMEZONE" calendar component in the scope of an
      iCalendar object.

   Value Type:  TEXT

   Property Parameters:  IANA and non-standard property parameters can
      be specified on this property.

   Conformance:  This property MUST be specified in a "VTIMEZONE"
      calendar component.

   Description:  This is the label by which a time zone calendar
      component is referenced by any iCalendar properties whose value
      type is either DATE-TIME or TIME and not intended to specify a UTC
      or a "floating" time.  The presence of the SOLIDUS character as a
      prefix, indicates that this "TZID" represents an unique ID in a
      globally defined time zone registry (when such registry is
      defined).

         Note: This document does not define a naming convention for
         time zone identifiers.  Implementers may want to use the naming
         conventions defined in existing time zone specifications such
         as the public-domain TZ database [TZDB].  The specification of
         globally unique time zone identifiers is not addressed by this
         document and is left for future study.




   Format Definition:  This property is defined by the following
      notation:

       tzid       = "TZID" tzidpropparam ":" [tzidprefix] text CRLF

       tzidpropparam      = *(";" other-param)

       ;tzidprefix        = "/"
       ; Defined previously. Just listed here for reader convenience.

   Example:  The following are examples of non-globally unique time zone
      identifiers:

       TZID:America/New_York

       TZID:America/Los_Angeles

      The following is an example of a fictitious globally unique time
      zone identifier:

       TZID:/example.org/America/New_York
