# Version

Source: RFC 5545, Section 3.7.4

3.7.4.  Version

   Property Name:  VERSION

   Purpose:  This property specifies the identifier corresponding to the
      highest version number or the minimum and maximum range of the
      iCalendar specification that is required in order to interpret the
      iCalendar object.

   Value Type:  TEXT

   Property Parameters:  IANA and non-standard property parameters can
      be specified on this property.

   Conformance:  This property MUST be specified once in an iCalendar
      object.

   Description:  A value of "2.0" corresponds to this memo.

   Format Definition:  This property is defined by the following
      notation:

       version    = "VERSION" verparam ":" vervalue CRLF

       verparam   = *(";" other-param)

       vervalue   = "2.0"         ;This memo
                  / maxver
                  / (minver ";" maxver)

       minver     = <A IANA-registered iCalendar version identifier>
       ;Minimum iCalendar version needed to parse the iCalendar object.

       maxver     = <A IANA-registered iCalendar version identifier>
       ;Maximum iCalendar version needed to parse the iCalendar object.





   Example:  The following is an example of this property:

       VERSION:2.0
