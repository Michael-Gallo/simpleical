# Organizer

Source: RFC 5545, Section 3.8.4.3

3.8.4.3.  Organizer

   Property Name:  ORGANIZER

   Purpose:  This property defines the organizer for a calendar
      component.

   Value Type:  CAL-ADDRESS

   Property Parameters:  IANA, non-standard, language, common name,
      directory entry reference, and sent-by property parameters can be
      specified on this property.

   Conformance:  This property MUST be specified in an iCalendar object
      that specifies a group-scheduled calendar entity.  This property
      MUST be specified in an iCalendar object that specifies the
      publication of a calendar user's busy time.  This property MUST
      NOT be specified in an iCalendar object that specifies only a time
      zone definition or that defines calendar components that are not
      group-scheduled components, but are components only on a single
      user's calendar.

   Description:  This property is specified within the "VEVENT",
      "VTODO", and "VJOURNAL" calendar components to specify the
      organizer of a group-scheduled calendar entity.  The property is
      specified within the "VFREEBUSY" calendar component to specify the
      calendar user requesting the free or busy time.  When publishing a
      "VFREEBUSY" calendar component, the property is used to specify
      the calendar that the published busy time came from.

      The property has the property parameters "CN", for specifying the
      common or display name associated with the "Organizer", "DIR", for
      specifying a pointer to the directory information associated with
      the "Organizer", "SENT-BY", for specifying another calendar user
      that is acting on behalf of the "Organizer".  The non-standard
      parameters may also be specified on this property.  If the
      "LANGUAGE" property parameter is specified, the identified
      language applies to the "CN" parameter value.




   Format Definition:  This property is defined by the following
      notation:

       organizer  = "ORGANIZER" orgparam ":"
                    cal-address CRLF

       orgparam   = *(
                  ;
                  ; The following are OPTIONAL,
                  ; but MUST NOT occur more than once.
                  ;
                  (";" cnparam) / (";" dirparam) / (";" sentbyparam) /
                  (";" languageparam) /
                  ;
                  ; The following is OPTIONAL,
                  ; and MAY occur more than once.
                  ;
                  (";" other-param)
                  ;
                  )

   Example:  The following is an example of this property:

       ORGANIZER;CN=John Smith:mailto:jsmith@example.com

      The following is an example of this property with a pointer to the
      directory information associated with the organizer:

       ORGANIZER;CN=JohnSmith;DIR="ldap://example.com:6666/o=DC%20Ass
        ociates,c=US???(cn=John%20Smith)":mailto:jsmith@example.com

      The following is an example of this property used by another
      calendar user who is acting on behalf of the organizer, with
      responses intended to be sent back to the organizer, not the other
      calendar user:

       ORGANIZER;SENT-BY="mailto:jane_doe@example.com":
        mailto:jsmith@example.com
