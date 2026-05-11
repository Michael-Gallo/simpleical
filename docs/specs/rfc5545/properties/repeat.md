# Repeat Count

Source: RFC 5545, Section 3.8.6.2

3.8.6.2.  Repeat Count

   Property Name:  REPEAT

   Purpose:  This property defines the number of times the alarm should
      be repeated, after the initial trigger.

   Value Type:  INTEGER

   Property Parameters:  IANA and non-standard property parameters can
      be specified on this property.

   Conformance:  This property can be specified in a "VALARM" calendar
      component.

   Description:  This property defines the number of times an alarm
      should be repeated after its initial trigger.  If the alarm
      triggers more than once, then this property MUST be specified
      along with the "DURATION" property.

   Format Definition:  This property is defined by the following
      notation:

       repeat  = "REPEAT" repparam ":" integer CRLF
       ;Default is "0", zero.

       repparam   = *(";" other-param)

   Example:  The following is an example of this property for an alarm
      that repeats 4 additional times with a 5-minute delay after the
      initial triggering of the alarm:

       REPEAT:4
       DURATION:PT5M
