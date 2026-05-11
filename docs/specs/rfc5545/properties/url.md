# Uniform Resource Locator

Source: RFC 5545, Section 3.8.4.6

3.8.4.6.  Uniform Resource Locator

   Property Name:  URL

   Purpose:  This property defines a Uniform Resource Locator (URL)
      associated with the iCalendar object.

   Value Type:  URI

   Property Parameters:  IANA and non-standard property parameters can
      be specified on this property.

   Conformance:  This property can be specified once in the "VEVENT",
      "VTODO", "VJOURNAL", or "VFREEBUSY" calendar components.

   Description:  This property may be used in a calendar component to
      convey a location where a more dynamic rendition of the calendar
      information associated with the calendar component can be found.
      This memo does not attempt to standardize the form of the URI, nor
      the format of the resource pointed to by the property value.  If
      the URL property and Content-Location MIME header are both
      specified, they MUST point to the same resource.


   Format Definition:  This property is defined by the following
      notation:

       url        = "URL" urlparam ":" uri CRLF

       urlparam   = *(";" other-param)

   Example:  The following is an example of this property:

       URL:http://example.com/pub/calendars/jsmith/mytime.ics
