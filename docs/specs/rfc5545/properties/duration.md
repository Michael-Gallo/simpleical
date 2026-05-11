# Duration

Source: RFC 5545, Section 3.8.2.5

3.8.2.5.  Duration

   Property Name:  DURATION

   Purpose:  This property specifies a positive duration of time.

   Value Type:  DURATION

   Property Parameters:  IANA and non-standard property parameters can
      be specified on this property.

   Conformance:  This property can be specified in "VEVENT", "VTODO", or
      "VALARM" calendar components.

   Description:  In a "VEVENT" calendar component the property may be
      used to specify a duration of the event, instead of an explicit
      end DATE-TIME.  In a "VTODO" calendar component the property may
      be used to specify a duration for the to-do, instead of an
      explicit due DATE-TIME.  In a "VALARM" calendar component the
      property may be used to specify the delay period prior to
      repeating an alarm.  When the "DURATION" property relates to a
      "DTSTART" property that is specified as a DATE value, then the
      "DURATION" property MUST be specified as a "dur-day" or "dur-week"
      value.

   Format Definition:  This property is defined by the following
      notation:

       duration   = "DURATION" durparam ":" dur-value CRLF
                    ;consisting of a positive duration of time.

       durparam   = *(";" other-param)

   Example:  The following is an example of this property that specifies
      an interval of time of one hour and zero minutes and zero seconds:

       DURATION:PT1H0M0S

      The following is an example of this property that specifies an
      interval of time of 15 minutes.

       DURATION:PT15M
