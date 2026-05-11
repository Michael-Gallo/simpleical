# IANA Considerations

Source: RFC 5545, Section 8

8.  IANA Considerations

8.1.  iCalendar Media Type Registration

   The Calendaring and Scheduling Core Object Specification is intended
   for use as a MIME content type.

   To: ietf-types@iana.org

   Subject: Registration of media type text/calendar

   Type name:  text

   Subtype name:  calendar

   Required parameters:  none

   Optional parameters:  charset, method, component, and optinfo

      The "charset" parameter is defined in [RFC2046] for subtypes of
      the "text" media type.  It is used to indicate the charset used in
      the body part.  The charset supported by this revision of
      iCalendar is UTF-8.  The use of any other charset is deprecated by
      this revision of iCalendar; however, note that this revision
      requires that compliant applications MUST accept iCalendar streams
      using either the UTF-8 or US-ASCII charset.

      The "method" parameter is used to convey the iCalendar object
      method or transaction semantics for the calendaring and scheduling
      information.  It also is an identifier for the restricted set of
      properties and values of which the iCalendar object consists.  The
      parameter is to be used as a guide for applications interpreting
      the information contained within the body part.  It SHOULD NOT be
      used to exclude or require particular pieces of information unless
      the identified method definition specifically calls for this
      behavior.  Unless specifically forbidden by a particular method
      definition, a text/calendar content type can contain any set of
      properties permitted by the Calendaring and Scheduling Core Object
      Specification.  The "method" parameter MUST be specified and MUST
      be set to the same value as the "METHOD" component property of the
      iCalendar objects of the iCalendar stream if and only if the
      iCalendar objects in the iCalendar stream all have a "METHOD"
      component property set to the same value.


      The value for the "method" parameter is defined as follows:

       method  = 1*(ALPHA / DIGIT / "-")
       ; IANA-registered iCalendar object method

      The "component" parameter conveys the type of iCalendar calendar
      component within the body part.  If the iCalendar object contains
      more than one calendar component type, then multiple component
      parameters MUST be specified.

      The value for the "component" parameter is defined as follows:

       component = "VEVENT"
                 / "VTODO"
                 / "VJOURNAL"
                 / "VFREEBUSY"
                 / "VTIMEZONE"
                 / iana-token
                 / x-name

      The "optinfo" parameter conveys optional information about the
      iCalendar object within the body part.  This parameter can only
      specify semantics already specified by the iCalendar object and
      that can be otherwise determined by parsing the body part.  In
      addition, the optional information specified by this parameter
      MUST be consistent with that information specified by the
      iCalendar object.  For example, it can be used to convey the
      "Attendee" response status to a meeting request.  The parameter
      value consists of a string value.

      The parameter can be specified multiple times.

      The value for the "optinfo" parameter is defined as follows:

       optinfo    = infovalue / qinfovalue

       infovalue  = iana-token / x-name

       qinfovalue = DQUOTE (infovalue) DQUOTE

   Encoding considerations:  This media type can contain 8bit
      characters, so the use of quoted-printable or base64 MIME Content-
      Transfer-Encodings might be necessary when iCalendar objects are
      transferred across protocols restricted to the 7bit repertoire.
      Note that a text valued property in the content entity can also
      have content encoding of special characters using a BACKSLASH
      character escapement technique.  This means that content values
      can end up being encoded twice.

   Security considerations:  See Section 7.

   Interoperability considerations:  This media type is intended to
      define a common format for conveying calendaring and scheduling
      information between different systems.  It is heavily based on the
      earlier [VCAL] industry specification.

   Published specification:  This specification.

   Applications that use this media type:  This media type is designed
      for widespread use by Internet calendaring and scheduling
      applications.  In addition, applications in the workflow and
      document management area might find this content-type applicable.
      The iTIP [2446bis], iMIP [2447bis], and CalDAV [RFC4791] Internet
      protocols directly use this media type also.

   Additional information:

      Magic number(s):  None.

      File extension(s):  The file extension of "ics" is to be used to
         designate a file containing (an arbitrary set of) calendaring
         and scheduling information consistent with this MIME content
         type.

         The file extension of "ifb" is to be used to designate a file
         containing free or busy time information consistent with this
         MIME content type.

      Macintosh file type code(s):  The file type code of "iCal" is to
         be used in Apple MacIntosh operating system environments to
         designate a file containing calendaring and scheduling
         information consistent with this MIME media type.

         The file type code of "iFBf" is to be used in Apple MacIntosh
         operating system environments to designate a file containing
         free or busy time information consistent with this MIME media
         type.

   Person & email address to contact for further information:  See the
      "Author's Address" section of this document.

   Intended usage:  COMMON

   Restrictions on usage:  There are no restrictions on where this media
      type can be used.

   Author:  See the "Author's Address" section of this document.

   Change controller:  IETF

8.2.  New iCalendar Elements Registration

   This section defines the process to register new or modified
   iCalendar elements, that is, components, properties, parameters,
   value data types, and values, with IANA.

8.2.1.  iCalendar Elements Registration Procedure

   The IETF will create a mailing list, icalendar@ietf.org, which can be
   used for public discussion of iCalendar elements proposals prior to
   registration.  Use of the mailing list is strongly encouraged.  The
   IESG will appoint a designated expert who will monitor the
   icalendar@ietf.org mailing list and review registrations.

   Registration of new iCalendar elements MUST be reviewed by the
   designated expert and published in an RFC.  A Standards Track RFC is
   REQUIRED for the registration of new value data types that modify
   existing properties, as well as for the registration of participation
   status values to be used in "VEVENT" calendar components.  A
   Standards Track RFC is also REQUIRED for registration of iCalendar
   elements that modify iCalendar elements previously documented in a
   Standards Track RFC.

   The registration procedure begins when a completed registration
   template, defined in the sections below, is sent to
   icalendar@ietf.org and iana@iana.org.  The designated expert is
   expected to tell IANA and the submitter of the registration within
   two weeks whether the registration is approved, approved with minor
   changes, or rejected with cause.  When a registration is rejected
   with cause, it can be re-submitted if the concerns listed in the
   cause are addressed.  Decisions made by the designated expert can be
   appealed to the IESG Applications Area Director, then to the IESG.
   They follow the normal appeals procedure for IESG decisions.

8.2.2.  Registration Template for Components

   A component is defined by completing the following template.

   Component name:  The name of the component.

   Purpose:  The purpose of the component.  Give a short but clear
      description.

   Format definition:  The ABNF for the component definition needs to be
      specified.


   Description:  Any special notes about the component, how it is to be
      used, etc.

   Example(s):  One or more examples of instances of the component need
      to be specified.

8.2.3.  Registration Template for Properties

   A property is defined by completing the following template.

   Property name:  The name of the property.

   Purpose:  The purpose of the property.  Give a short but clear
      description.

   Value type:  Any of the valid value types for the property value need
      to be specified.  The default value type also needs to be
      specified.

   Property parameters:  Any of the valid property parameters for the
      property MUST be specified.

   Conformance:  The calendar components in which the property can
      appear MUST be specified.

   Description:  Any special notes about the property, how it is to be
      used, etc.

   Format definition:  The ABNF for the property definition needs to be
      specified.

   Example(s):  One or more examples of instances of the property need
      to be specified.

8.2.4.  Registration Template for Parameters

   A parameter is defined by completing the following template.

   Parameter name:  The name of the parameter.

   Purpose:  The purpose of the parameter.  Give a short but clear
      description.

   Format definition:  The ABNF for the parameter definition needs to be
      specified.

   Description:  Any special notes about the parameter, how it is to be
      used, etc.

   Example(s):  One or more examples of instances of the parameter need
      to be specified.

8.2.5.  Registration Template for Value Data Types

   A value data type is defined by completing the following template.

   Value name:  The name of the value type.

   Purpose:  The purpose of the value type.  Give a short but clear
      description.

   Format definition:  The ABNF for the value type definition needs to
      be specified.

   Description:  Any special notes about the value type, how it is to be
      used, etc.

   Example(s):  One or more examples of instances of the value type need
      to be specified.

8.2.6.  Registration Template for Values

   A value is defined by completing the following template.

   Value:  The value literal.

   Purpose:  The purpose of the value.  Give a short but clear
      description.

   Conformance:  The calendar properties and/or parameters that can take
      this value need to be specified.

   Example(s):  One or more examples of instances of the value need to
      be specified.

   The following is a fictitious example of a registration of an
   iCalendar value:

   Value:  TOP-SECRET

   Purpose:  This value is used to specify the access classification of
      top-secret calendar components.

   Conformance:  This value can be used with the "CLASS" property.




   Example(s):  The following is an example of this value used with the
      "CLASS" property:

     CLASS:TOP-SECRET

8.3.  Initial iCalendar Elements Registries

   The IANA created and maintains the following registries for iCalendar
   elements with pointers to appropriate reference documents.

8.3.1.  Components Registry

   The following table has been used to initialize the components
   registry.

             +-----------+---------+-------------------------+
             | Component | Status  | Reference               |
             +-----------+---------+-------------------------+
             | VCALENDAR | Current | RFC 5545, Section 3.4   |
             |           |         |                         |
             | VEVENT    | Current | RFC 5545, Section 3.6.1 |
             |           |         |                         |
             | VTODO     | Current | RFC 5545, Section 3.6.2 |
             |           |         |                         |
             | VJOURNAL  | Current | RFC 5545, Section 3.6.3 |
             |           |         |                         |
             | VFREEBUSY | Current | RFC 5545, Section 3.6.4 |
             |           |         |                         |
             | VTIMEZONE | Current | RFC 5545, Section 3.6.5 |
             |           |         |                         |
             | VALARM    | Current | RFC 5545, Section 3.6.6 |
             |           |         |                         |
             | STANDARD  | Current | RFC 5545, Section 3.6.5 |
             |           |         |                         |
             | DAYLIGHT  | Current | RFC 5545, Section 3.6.5 |
             +-----------+---------+-------------------------+













8.3.2.  Properties Registry

   The following table is has been used to initialize the properties
   registry.

      +------------------+------------+----------------------------+
      | Property         | Status     | Reference                  |
      +------------------+------------+----------------------------+
      | CALSCALE         | Current    | RFC 5545, Section 3.7.1    |
      | METHOD           | Current    | RFC 5545, Section 3.7.2    |
      |                  |            |                            |
      | PRODID           | Current    | RFC 5545, Section 3.7.3    |
      |                  |            |                            |
      | VERSION          | Current    | RFC 5545, Section 3.7.4    |
      |                  |            |                            |
      | ATTACH           | Current    | RFC 5545, Section 3.8.1.1  |
      |                  |            |                            |
      | CATEGORIES       | Current    | RFC 5545, Section 3.8.1.2  |
      |                  |            |                            |
      | CLASS            | Current    | RFC 5545, Section 3.8.1.3  |
      |                  |            |                            |
      | COMMENT          | Current    | RFC 5545, Section 3.8.1.4  |
      |                  |            |                            |
      | DESCRIPTION      | Current    | RFC 5545, Section 3.8.1.5  |
      |                  |            |                            |
      | GEO              | Current    | RFC 5545, Section 3.8.1.6  |
      |                  |            |                            |
      | LOCATION         | Current    | RFC 5545, Section 3.8.1.7  |
      |                  |            |                            |
      | PERCENT-COMPLETE | Current    | RFC 5545, Section 3.8.1.8  |
      |                  |            |                            |
      | PRIORITY         | Current    | RFC 5545, Section 3.8.1.9  |
      |                  |            |                            |
      | RESOURCES        | Current    | RFC 5545, Section 3.8.1.10 |
      |                  |            |                            |
      | STATUS           | Current    | RFC 5545, Section 3.8.1.11 |
      |                  |            |                            |
      | SUMMARY          | Current    | RFC 5545, Section 3.8.1.12 |
      |                  |            |                            |
      | COMPLETED        | Current    | RFC 5545, Section 3.8.2.1  |
      |                  |            |                            |
      | DTEND            | Current    | RFC 5545, Section 3.8.2.2  |
      |                  |            |                            |
      | DUE              | Current    | RFC 5545, Section 3.8.2.3  |
      |                  |            |                            |
      | DTSTART          | Current    | RFC 5545, Section 3.8.2.4  |
      |                  |            |                            |
      | DURATION         | Current    | RFC 5545, Section 3.8.2.5  |

      |                  |            |                            |
      | FREEBUSY         | Current    | RFC 5545, Section 3.8.2.6  |
      |                  |            |                            |
      | TRANSP           | Current    | RFC 5545, Section 3.8.2.7  |
      |                  |            |                            |
      | TZID             | Current    | RFC 5545, Section 3.8.3.1  |
      |                  |            |                            |
      | TZNAME           | Current    | RFC 5545, Section 3.8.3.2  |
      |                  |            |                            |
      | TZOFFSETFROM     | Current    | RFC 5545, Section 3.8.3.3  |
      |                  |            |                            |
      | TZOFFSETTO       | Current    | RFC 5545, Section 3.8.3.4  |
      |                  |            |                            |
      | TZURL            | Current    | RFC 5545, Section 3.8.3.5  |
      |                  |            |                            |
      | ATTENDEE         | Current    | RFC 5545, Section 3.8.4.1  |
      |                  |            |                            |
      | CONTACT          | Current    | RFC 5545, Section 3.8.4.2  |
      |                  |            |                            |
      | ORGANIZER        | Current    | RFC 5545, Section 3.8.4.3  |
      |                  |            |                            |
      | RECURRENCE-ID    | Current    | RFC 5545, Section 3.8.4.4  |
      |                  |            |                            |
      | RELATED-TO       | Current    | RFC 5545, Section 3.8.4.5  |
      |                  |            |                            |
      | URL              | Current    | RFC 5545, Section 3.8.4.6  |
      |                  |            |                            |
      | UID              | Current    | RFC 5545, Section 3.8.4.7  |
      |                  |            |                            |
      | EXDATE           | Current    | RFC 5545, Section 3.8.5.1  |
      |                  |            |                            |
      | EXRULE           | Deprecated | [RFC2445], Section 4.8.5.2 |
      |                  |            |                            |
      | RDATE            | Current    | RFC 5545, Section 3.8.5.2  |
      |                  |            |                            |
      | RRULE            | Current    | RFC 5545, Section 3.8.5.3  |
      |                  |            |                            |
      | ACTION           | Current    | RFC 5545, Section 3.8.6.1  |
      |                  |            |                            |
      | REPEAT           | Current    | RFC 5545, Section 3.8.6.2  |
      |                  |            |                            |
      | TRIGGER          | Current    | RFC 5545, Section 3.8.6.3  |
      |                  |            |                            |
      | CREATED          | Current    | RFC 5545, Section 3.8.7.1  |
      |                  |            |                            |
      | DTSTAMP          | Current    | RFC 5545, Section 3.8.7.2  |
      |                  |            |                            |
      | LAST-MODIFIED    | Current    | RFC 5545, Section 3.8.7.3  |

      |                  |            |                            |
      | SEQUENCE         | Current    | RFC 5545, Section 3.8.7.4  |
      |                  |            |                            |
      | REQUEST-STATUS   | Current    | RFC 5545, Section 3.8.8.3  |
      +------------------+------------+----------------------------+

8.3.3.  Parameters Registry

   The following table has been used to initialize the parameters
   registry.

          +----------------+---------+--------------------------+
          | Parameter      | Status  | Reference                |
          +----------------+---------+--------------------------+
          | ALTREP         | Current | RFC 5545, Section 3.2.1  |
          |                |         |                          |
          | CN             | Current | RFC 5545, Section 3.2.2  |
          |                |         |                          |
          | CUTYPE         | Current | RFC 5545, Section 3.2.3  |
          |                |         |                          |
          | DELEGATED-FROM | Current | RFC 5545, Section 3.2.4  |
          |                |         |                          |
          | DELEGATED-TO   | Current | RFC 5545, Section 3.2.5  |
          |                |         |                          |
          | DIR            | Current | RFC 5545, Section 3.2.6  |
          |                |         |                          |
          | ENCODING       | Current | RFC 5545, Section 3.2.7  |
          |                |         |                          |
          | FMTTYPE        | Current | RFC 5545, Section 3.2.8  |
          |                |         |                          |
          | FBTYPE         | Current | RFC 5545, Section 3.2.9  |
          |                |         |                          |
          | LANGUAGE       | Current | RFC 5545, Section 3.2.10 |
          |                |         |                          |
          | MEMBER         | Current | RFC 5545, Section 3.2.11 |
          |                |         |                          |
          | PARTSTAT       | Current | RFC 5545, Section 3.2.12 |
          |                |         |                          |
          | RANGE          | Current | RFC 5545, Section 3.2.13 |
          |                |         |                          |
          | RELATED        | Current | RFC 5545, Section 3.2.14 |
          |                |         |                          |
          | RELTYPE        | Current | RFC 5545, Section 3.2.15 |
          |                |         |                          |
          | ROLE           | Current | RFC 5545, Section 3.2.16 |
          |                |         |                          |
          | RSVP           | Current | RFC 5545, Section 3.2.17 |
          |                |         |                          |

          | SENT-BY        | Current | RFC 5545, Section 3.2.18 |
          |                |         |                          |
          | TZID           | Current | RFC 5545, Section 3.2.19 |
          |                |         |                          |
          | VALUE          | Current | RFC 5545, Section 3.2.20 |
          +----------------+---------+--------------------------+

8.3.4.  Value Data Types Registry

   The following table has been used to initialize the value data types
   registry.

         +-----------------+---------+--------------------------+
         | Value Data Type | Status  | Reference                |
         +-----------------+---------+--------------------------+
         | BINARY          | Current | RFC 5545, Section 3.3.1  |
         |                 |         |                          |
         | BOOLEAN         | Current | RFC 5545, Section 3.3.2  |
         |                 |         |                          |
         | CAL-ADDRESS     | Current | RFC 5545, Section 3.3.3  |
         |                 |         |                          |
         | DATE            | Current | RFC 5545, Section 3.3.4  |
         |                 |         |                          |
         | DATE-TIME       | Current | RFC 5545, Section 3.3.5  |
         |                 |         |                          |
         | DURATION        | Current | RFC 5545, Section 3.3.6  |
         |                 |         |                          |
         | FLOAT           | Current | RFC 5545, Section 3.3.7  |
         |                 |         |                          |
         | INTEGER         | Current | RFC 5545, Section 3.3.8  |
         |                 |         |                          |
         | PERIOD          | Current | RFC 5545, Section 3.3.9  |
         |                 |         |                          |
         | RECUR           | Current | RFC 5545, Section 3.3.10 |
         |                 |         |                          |
         | TEXT            | Current | RFC 5545, Section 3.3.11 |
         |                 |         |                          |
         | TIME            | Current | RFC 5545, Section 3.3.12 |
         |                 |         |                          |
         | URI             | Current | RFC 5545, Section 3.3.13 |
         |                 |         |                          |
         | UTC-OFFSET      | Current | RFC 5545, Section 3.3.14 |
         +-----------------+---------+--------------------------+






8.3.5.  Calendar User Types Registry

   The following table has been used to initialize the calendar user
   types registry.

        +--------------------+---------+-------------------------+
        | Calendar User Type | Status  | Reference               |
        +--------------------+---------+-------------------------+
        | INDIVIDUAL         | Current | RFC 5545, Section 3.2.3 |
        |                    |         |                         |
        | GROUP              | Current | RFC 5545, Section 3.2.3 |
        |                    |         |                         |
        | RESOURCE           | Current | RFC 5545, Section 3.2.3 |
        |                    |         |                         |
        | ROOM               | Current | RFC 5545, Section 3.2.3 |
        |                    |         |                         |
        | UNKNOWN            | Current | RFC 5545, Section 3.2.3 |
        +--------------------+---------+-------------------------+

8.3.6.  Free/Busy Time Types Registry

   The following table has been used to initialize the free/busy time
   types registry.

        +---------------------+---------+-------------------------+
        | Free/Busy Time Type | Status  | Reference               |
        +---------------------+---------+-------------------------+
        | FREE                | Current | RFC 5545, Section 3.2.9 |
        |                     |         |                         |
        | BUSY                | Current | RFC 5545, Section 3.2.9 |
        |                     |         |                         |
        | BUSY-UNAVAILABLE    | Current | RFC 5545, Section 3.2.9 |
        |                     |         |                         |
        | BUSY-TENTATIVE      | Current | RFC 5545, Section 3.2.9 |
        +---------------------+---------+-------------------------+














8.3.7.  Participation Statuses Registry

   The following table has been used to initialize the participation
   statuses registry.

        +--------------------+---------+--------------------------+
        | Participant Status | Status  | Reference                |
        +--------------------+---------+--------------------------+
        | NEEDS-ACTION       | Current | RFC 5545, Section 3.2.12 |
        |                    |         |                          |
        | ACCEPTED           | Current | RFC 5545, Section 3.2.12 |
        |                    |         |                          |
        | DECLINED           | Current | RFC 5545, Section 3.2.12 |
        |                    |         |                          |
        | TENTATIVE          | Current | RFC 5545, Section 3.2.12 |
        |                    |         |                          |
        | DELEGATED          | Current | RFC 5545, Section 3.2.12 |
        |                    |         |                          |
        | COMPLETED          | Current | RFC 5545, Section 3.2.12 |
        |                    |         |                          |
        | IN-PROCESS         | Current | RFC 5545, Section 3.2.12 |
        +--------------------+---------+--------------------------+

8.3.8.  Relationship Types Registry

   The following table has been used to initialize the relationship
   types registry.

        +-------------------+---------+--------------------------+
        | Relationship Type | Status  | Reference                |
        +-------------------+---------+--------------------------+
        | CHILD             | Current | RFC 5545, Section 3.2.15 |
        |                   |         |                          |
        | PARENT            | Current | RFC 5545, Section 3.2.15 |
        |                   |         |                          |
        | SIBLING           | Current | RFC 5545, Section 3.2.15 |
        +-------------------+---------+--------------------------+












8.3.9.  Participation Roles Registry

   The following table has been used to initialize the participation
   roles registry.

         +-----------------+---------+--------------------------+
         | Role Type       | Status  | Reference                |
         +-----------------+---------+--------------------------+
         | CHAIR           | Current | RFC 5545, Section 3.2.16 |
         |                 |         |                          |
         | REQ-PARTICIPANT | Current | RFC 5545, Section 3.2.16 |
         |                 |         |                          |
         | OPT-PARTICIPANT | Current | RFC 5545, Section 3.2.16 |
         |                 |         |                          |
         | NON-PARTICIPANT | Current | RFC 5545, Section 3.2.16 |
         +-----------------+---------+--------------------------+

8.3.10.  Actions Registry

   The following table has been used to initialize the actions registry.

          +-----------+------------+----------------------------+
          | Action    | Status     | Reference                  |
          +-----------+------------+----------------------------+
          | AUDIO     | Current    | RFC 5545, Section 3.8.6.1  |
          |           |            |                            |
          | DISPLAY   | Current    | RFC 5545, Section 3.8.6.1  |
          |           |            |                            |
          | EMAIL     | Current    | RFC 5545, Section 3.8.6.1  |
          |           |            |                            |
          | PROCEDURE | Deprecated | [RFC2445], Section 4.8.6.1 |
          +-----------+------------+----------------------------+

8.3.11.  Classifications Registry

   The following table has been used to initialize the classifications
   registry.

         +----------------+---------+---------------------------+
         | Classification | Status  | Reference                 |
         +----------------+---------+---------------------------+
         | PUBLIC         | Current | RFC 5545, Section 3.8.1.3 |
         |                |         |                           |
         | PRIVATE        | Current | RFC 5545, Section 3.8.1.3 |
         |                |         |                           |
         | CONFIDENTIAL   | Current | RFC 5545, Section 3.8.1.3 |
         +----------------+---------+---------------------------+


8.3.12.  Methods Registry

   No values are defined in this document for the "METHOD" property.
