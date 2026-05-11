# Time Zone URL

Source: RFC 5545, Section 3.8.3.5

3.8.3.5.  Time Zone URL

   Property Name:  TZURL

   Purpose:  This property provides a means for a "VTIMEZONE" component
      to point to a network location that can be used to retrieve an up-
      to-date version of itself.

   Value Type:  URI

   Property Parameters:  IANA and non-standard property parameters can
      be specified on this property.

   Conformance:  This property can be specified in a "VTIMEZONE"
      calendar component.

   Description:  This property provides a means for a "VTIMEZONE"
      component to point to a network location that can be used to
      retrieve an up-to-date version of itself.  This provides a hook to
      handle changes government bodies impose upon time zone
      definitions.  Retrieval of this resource results in an iCalendar
      object containing a single "VTIMEZONE" component and a "METHOD"
      property set to PUBLISH.

   Format Definition:  This property is defined by the following
      notation:

       tzurl      = "TZURL" tzurlparam ":" uri CRLF

       tzurlparam = *(";" other-param)

   Example:  The following is an example of this property:

    TZURL:http://timezones.example.org/tz/America-Los_Angeles.ics
