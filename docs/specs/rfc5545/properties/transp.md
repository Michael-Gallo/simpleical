# Time Transparency

Source: RFC 5545, Section 3.8.2.7

3.8.2.7.  Time Transparency

   Property Name:  TRANSP

   Purpose:  This property defines whether or not an event is
      transparent to busy time searches.

   Value Type:  TEXT

   Property Parameters:  IANA and non-standard property parameters can
      be specified on this property.

   Conformance:  This property can be specified once in a "VEVENT"
      calendar component.

   Description:  Time Transparency is the characteristic of an event
      that determines whether it appears to consume time on a calendar.
      Events that consume actual time for the individual or resource
      associated with the calendar SHOULD be recorded as OPAQUE,
      allowing them to be detected by free/busy time searches.  Other
      events, which do not take up the individual's (or resource's) time
      SHOULD be recorded as TRANSPARENT, making them invisible to free/
      busy time searches.

   Format Definition:  This property is defined by the following
      notation:

       transp     = "TRANSP" transparam ":" transvalue CRLF

       transparam = *(";" other-param)

       transvalue = "OPAQUE"
                   ;Blocks or opaque on busy time searches.
                   / "TRANSPARENT"
                   ;Transparent on busy time searches.
       ;Default value is OPAQUE

   Example:  The following is an example of this property for an event
      that is transparent or does not block on free/busy time searches:

       TRANSP:TRANSPARENT

      The following is an example of this property for an event that is
      opaque or blocks on free/busy time searches:

       TRANSP:OPAQUE
