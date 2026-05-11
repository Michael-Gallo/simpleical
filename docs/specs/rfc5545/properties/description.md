# Description

Source: RFC 5545, Section 3.8.1.5

3.8.1.5.  Description

   Property Name:  DESCRIPTION

   Purpose:  This property provides a more complete description of the
      calendar component than that provided by the "SUMMARY" property.

   Value Type:  TEXT

   Property Parameters:  IANA, non-standard, alternate text
      representation, and language property parameters can be specified
      on this property.




   Conformance:  The property can be specified in the "VEVENT", "VTODO",
      "VJOURNAL", or "VALARM" calendar components.  The property can be
      specified multiple times only within a "VJOURNAL" calendar
      component.

   Description:  This property is used in the "VEVENT" and "VTODO" to
      capture lengthy textual descriptions associated with the activity.

      This property is used in the "VJOURNAL" calendar component to
      capture one or more textual journal entries.

      This property is used in the "VALARM" calendar component to
      capture the display text for a DISPLAY category of alarm, and to
      capture the body text for an EMAIL category of alarm.

   Format Definition:  This property is defined by the following
      notation:

       description = "DESCRIPTION" descparam ":" text CRLF

       descparam   = *(
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

   Example:  The following is an example of this property with formatted
      line breaks in the property value:

       DESCRIPTION:Meeting to provide technical review for "Phoenix"
         design.\nHappy Face Conference Room. Phoenix design team
         MUST attend this meeting.\nRSVP to team leader.
