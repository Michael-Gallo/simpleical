# Time Zone Offset From

Source: RFC 5545, Section 3.8.3.3

3.8.3.3.  Time Zone Offset From

   Property Name:  TZOFFSETFROM

   Purpose:  This property specifies the offset that is in use prior to
      this time zone observance.

   Value Type:  UTC-OFFSET

   Property Parameters:  IANA and non-standard property parameters can
      be specified on this property.

   Conformance:  This property MUST be specified in "STANDARD" and
      "DAYLIGHT" sub-components.

   Description:  This property specifies the offset that is in use prior
      to this time observance.  It is used to calculate the absolute
      time at which the transition to a given observance takes place.
      This property MUST only be specified in a "VTIMEZONE" calendar
      component.  A "VTIMEZONE" calendar component MUST include this
      property.  The property value is a signed numeric indicating the
      number of hours and possibly minutes from UTC.  Positive numbers
      represent time zones east of the prime meridian, or ahead of UTC.
      Negative numbers represent time zones west of the prime meridian,
      or behind UTC.


   Format Definition:  This property is defined by the following
      notation:

       tzoffsetfrom       = "TZOFFSETFROM" frmparam ":" utc-offset
                            CRLF

       frmparam   = *(";" other-param)

   Example:  The following are examples of this property:

       TZOFFSETFROM:-0500

       TZOFFSETFROM:+1345
