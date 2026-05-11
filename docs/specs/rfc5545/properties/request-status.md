# Request Status

Source: RFC 5545, Section 3.8.8.3

3.8.8.3.  Request Status

   Property Name:  REQUEST-STATUS

   Purpose:  This property defines the status code returned for a
      scheduling request.

   Value Type:  TEXT

   Property Parameters:  IANA, non-standard, and language property
      parameters can be specified on this property.

   Conformance:  The property can be specified in the "VEVENT", "VTODO",
      "VJOURNAL", or "VFREEBUSY" calendar component.

   Description:  This property is used to return status code information
      related to the processing of an associated iCalendar object.  The
      value type for this property is TEXT.





      The value consists of a short return status component, a longer
      return status description component, and optionally a status-
      specific data component.  The components of the value are
      separated by the SEMICOLON character.

      The short return status is a PERIOD character separated pair or
      3-tuple of integers.  For example, "3.1" or "3.1.1".  The
      successive levels of integers provide for a successive level of
      status code granularity.

      The following are initial classes for the return status code.
      Individual iCalendar object methods will define specific return
      status codes for these classes.  In addition, other classes for
      the return status code may be defined using the registration
      process defined later in this memo.

   +--------+----------------------------------------------------------+
   | Short  | Longer Return Status Description                         |
   | Return |                                                          |
   | Status |                                                          |
   | Code   |                                                          |
   +--------+----------------------------------------------------------+
   | 1.xx   | Preliminary success.  This class of status code          |
   |        | indicates that the request has been initially processed  |
   |        | but that completion is pending.                          |
   |        |                                                          |
   | 2.xx   | Successful.  This class of status code indicates that    |
   |        | the request was completed successfully.  However, the    |
   |        | exact status code can indicate that a fallback has been  |
   |        | taken.                                                   |
   |        |                                                          |
   | 3.xx   | Client Error.  This class of status code indicates that  |
   |        | the request was not successful.  The error is the result |
   |        | of either a syntax or a semantic error in the client-    |
   |        | formatted request.  Request should not be retried until  |
   |        | the condition in the request is corrected.               |
   |        |                                                          |
   | 4.xx   | Scheduling Error.  This class of status code indicates   |
   |        | that the request was not successful.  Some sort of error |
   |        | occurred within the calendaring and scheduling service,  |
   |        | not directly related to the request itself.              |
   +--------+----------------------------------------------------------+







   Format Definition:  This property is defined by the following
      notation:

       rstatus    = "REQUEST-STATUS" rstatparam ":"
                    statcode ";" statdesc [";" extdata]

       rstatparam = *(
                  ;
                  ; The following is OPTIONAL,
                  ; but MUST NOT occur more than once.
                  ;
                  (";" languageparam) /
                  ;
                  ; The following is OPTIONAL,
                  ; and MAY occur more than once.
                  ;
                  (";" other-param)
                  ;
                  )

       statcode   = 1*DIGIT 1*2("." 1*DIGIT)
       ;Hierarchical, numeric return status code

       statdesc   = text
       ;Textual status description

       extdata    = text
       ;Textual exception data.  For example, the offending property
       ;name and value or complete property line.

   Example:  The following are some possible examples of this property.

      The COMMA and SEMICOLON separator characters in the property value
      are BACKSLASH character escaped because they appear in a text
      value.

       REQUEST-STATUS:2.0;Success

       REQUEST-STATUS:3.1;Invalid property value;DTSTART:96-Apr-01

       REQUEST-STATUS:2.8; Success\, repeating event ignored. Scheduled
        as a single event.;RRULE:FREQ=WEEKLY\;INTERVAL=2

       REQUEST-STATUS:4.1;Event conflict.  Date-time is busy.

       REQUEST-STATUS:3.7;Invalid calendar user;ATTENDEE:
        mailto:jsmith@example.com
