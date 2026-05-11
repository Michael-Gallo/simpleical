# Summary

Source: RFC 5545, Section 3.8.1.12

3.8.1.12.  Summary

   Property Name:  SUMMARY

   Purpose:  This property defines a short summary or subject for the
      calendar component.

   Value Type:  TEXT

   Property Parameters:  IANA, non-standard, alternate text
      representation, and language property parameters can be specified
      on this property.

   Conformance:  The property can be specified in "VEVENT", "VTODO",
      "VJOURNAL", or "VALARM" calendar components.

   Description:  This property is used in the "VEVENT", "VTODO", and
      "VJOURNAL" calendar components to capture a short, one-line
      summary about the activity or journal entry.

      This property is used in the "VALARM" calendar component to
      capture the subject of an EMAIL category of alarm.

   Format Definition:  This property is defined by the following
      notation:




       summary    = "SUMMARY" summparam ":" text CRLF

       summparam  = *(
                  ;
                  ; The following are OPTIONAL,
                  ; but MUST NOT occur more than once.
                  ;
                  (";" altrepparam) / (";" languageparam) /
                  ;
                  ; The following is OPTIONAL,
                  ; and MAY occur more than once.
                  ;
                  (";" other-param)
                  ;
                  )

   Example:  The following is an example of this property:

       SUMMARY:Department Party
