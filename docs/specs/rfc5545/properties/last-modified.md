# Last Modified

Source: RFC 5545, Section 3.8.7.3

3.8.7.3.  Last Modified

   Property Name:  LAST-MODIFIED

   Purpose:  This property specifies the date and time that the
      information associated with the calendar component was last
      revised in the calendar store.

         Note: This is analogous to the modification date and time for a
         file in the file system.

   Value Type:  DATE-TIME

   Property Parameters:  IANA and non-standard property parameters can
      be specified on this property.

   Conformance:  This property can be specified in the "VEVENT",
      "VTODO", "VJOURNAL", or "VTIMEZONE" calendar components.

   Description:  The property value MUST be specified in the UTC time
      format.

   Format Definition:  This property is defined by the following
      notation:

       last-mod   = "LAST-MODIFIED" lstparam ":" date-time CRLF

       lstparam   = *(";" other-param)

   Example:  The following is an example of this property:

       LAST-MODIFIED:19960817T133000Z
