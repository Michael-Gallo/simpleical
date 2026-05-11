# Date-Time Completed

Source: RFC 5545, Section 3.8.2.1

3.8.2.1.  Date-Time Completed

   Property Name:  COMPLETED

   Purpose:  This property defines the date and time that a to-do was
      actually completed.

   Value Type:  DATE-TIME

   Property Parameters:  IANA and non-standard property parameters can
      be specified on this property.

   Conformance:  The property can be specified in a "VTODO" calendar
      component.  The value MUST be specified as a date with UTC time.

   Description:  This property defines the date and time that a to-do
      was actually completed.

   Format Definition:  This property is defined by the following
      notation:




       completed  = "COMPLETED" compparam ":" date-time CRLF

       compparam  = *(";" other-param)

   Example:  The following is an example of this property:

    COMPLETED:19960401T150000Z
