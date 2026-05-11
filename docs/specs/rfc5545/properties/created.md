# Date-Time Created

Source: RFC 5545, Section 3.8.7.1

3.8.7.1.  Date-Time Created

   Property Name:  CREATED

   Purpose:  This property specifies the date and time that the calendar
      information was created by the calendar user agent in the calendar
      store.

         Note: This is analogous to the creation date and time for a
         file in the file system.

   Value Type:  DATE-TIME

   Property Parameters:  IANA and non-standard property parameters can
      be specified on this property.

   Conformance:  The property can be specified once in "VEVENT",
      "VTODO", or "VJOURNAL" calendar components.  The value MUST be
      specified as a date with UTC time.

   Description:  This property specifies the date and time that the
      calendar information was created by the calendar user agent in the
      calendar store.

   Format Definition:  This property is defined by the following
      notation:

       created    = "CREATED" creaparam ":" date-time CRLF

       creaparam  = *(";" other-param)

   Example:  The following is an example of this property:

       CREATED:19960329T133000Z
