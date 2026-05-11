# Component Properties

Source: RFC 5545, Section 3.8

3.8.  Component Properties

   The following properties can appear within calendar components, as
   specified by each component property definition.

3.8.1.  Descriptive Component Properties

   The following properties specify descriptive information about
   calendar components.

3.8.1.1.  Attachment

   Property Name:  ATTACH

   Purpose:  This property provides the capability to associate a
      document object with a calendar component.

   Value Type:  The default value type for this property is URI.  The
      value type can also be set to BINARY to indicate inline binary
      encoded content information.

   Property Parameters:  IANA, non-standard, inline encoding, and value
      data type property parameters can be specified on this property.
      The format type parameter can be specified on this property and is
      RECOMMENDED for inline binary encoded content information.

   Conformance:  This property can be specified multiple times in a
      "VEVENT", "VTODO", "VJOURNAL", or "VALARM" calendar component with
      the exception of AUDIO alarm that only allows this property to
      occur once.

   Description:  This property is used in "VEVENT", "VTODO", and
      "VJOURNAL" calendar components to associate a resource (e.g.,
      document) with the calendar component.  This property is used in
      "VALARM" calendar components to specify an audio sound resource or
      an email message attachment.  This property can be specified as a
      URI pointing to a resource or as inline binary encoded content.

      When this property is specified as inline binary encoded content,
      calendar applications MAY attempt to guess the media type of the
      resource via inspection of its content if and only if the media
      type of the resource is not given by the "FMTTYPE" parameter.  If
      the media type remains unknown, calendar applications SHOULD treat
      it as type "application/octet-stream".

   Format Definition:  This property is defined by the following
      notation:

       attach     = "ATTACH" attachparam ( ":" uri ) /
                    (
                      ";" "ENCODING" "=" "BASE64"
                      ";" "VALUE" "=" "BINARY"
                      ":" binary
                    )
                    CRLF

       attachparam = *(
                   ;
                   ; The following is OPTIONAL for a URI value,
                   ; RECOMMENDED for a BINARY value,
                   ; and MUST NOT occur more than once.
                   ;
                   (";" fmttypeparam) /
                   ;
                   ; The following is OPTIONAL,
                   ; and MAY occur more than once.
                   ;
                   (";" other-param)
                   ;
                   )

   Example:  The following are examples of this property:

       ATTACH:CID:jsmith.part3.960817T083000.xyzMail@example.com

       ATTACH;FMTTYPE=application/postscript:ftp://example.com/pub/
        reports/r-960812.ps

3.8.1.2.  Categories

   Property Name:  CATEGORIES

   Purpose:  This property defines the categories for a calendar
      component.

   Value Type:  TEXT

   Property Parameters:  IANA, non-standard, and language property
      parameters can be specified on this property.

   Conformance:  The property can be specified within "VEVENT", "VTODO",
      or "VJOURNAL" calendar components.


   Description:  This property is used to specify categories or subtypes
      of the calendar component.  The categories are useful in searching
      for a calendar component of a particular type and category.
      Within the "VEVENT", "VTODO", or "VJOURNAL" calendar components,
      more than one category can be specified as a COMMA-separated list
      of categories.

   Format Definition:  This property is defined by the following
      notation:

       categories = "CATEGORIES" catparam ":" text *("," text)
                    CRLF

       catparam   = *(
                  ;
                  ; The following is OPTIONAL,
                  ; but MUST NOT occur more than once.
                  ;
                  (";" languageparam ) /
                  ;
                  ; The following is OPTIONAL,
                  ; and MAY occur more than once.
                  ;
                  (";" other-param)
                  ;
                  )

   Example:  The following are examples of this property:

       CATEGORIES:APPOINTMENT,EDUCATION

       CATEGORIES:MEETING

3.8.1.3.  Classification

   Property Name:  CLASS

   Purpose:  This property defines the access classification for a
      calendar component.

   Value Type:  TEXT

   Property Parameters:  IANA and non-standard property parameters can
      be specified on this property.

   Conformance:  The property can be specified once in a "VEVENT",
      "VTODO", or "VJOURNAL" calendar components.


   Description:  An access classification is only one component of the
      general security system within a calendar application.  It
      provides a method of capturing the scope of the access the
      calendar owner intends for information within an individual
      calendar entry.  The access classification of an individual
      iCalendar component is useful when measured along with the other
      security components of a calendar system (e.g., calendar user
      authentication, authorization, access rights, access role, etc.).
      Hence, the semantics of the individual access classifications
      cannot be completely defined by this memo alone.  Additionally,
      due to the "blind" nature of most exchange processes using this
      memo, these access classifications cannot serve as an enforcement
      statement for a system receiving an iCalendar object.  Rather,
      they provide a method for capturing the intention of the calendar
      owner for the access to the calendar component.  If not specified
      in a component that allows this property, the default value is
      PUBLIC.  Applications MUST treat x-name and iana-token values they
      don't recognize the same way as they would the PRIVATE value.

   Format Definition:  This property is defined by the following
      notation:

       class      = "CLASS" classparam ":" classvalue CRLF

       classparam = *(";" other-param)

       classvalue = "PUBLIC" / "PRIVATE" / "CONFIDENTIAL" / iana-token
                  / x-name
       ;Default is PUBLIC

   Example:  The following is an example of this property:

       CLASS:PUBLIC

3.8.1.4.  Comment

   Property Name:  COMMENT

   Purpose:  This property specifies non-processing information intended
      to provide a comment to the calendar user.

   Value Type:  TEXT

   Property Parameters:  IANA, non-standard, alternate text
      representation, and language property parameters can be specified
      on this property.



   Conformance:  This property can be specified multiple times in
      "VEVENT", "VTODO", "VJOURNAL", and "VFREEBUSY" calendar components
      as well as in the "STANDARD" and "DAYLIGHT" sub-components.

   Description:  This property is used to specify a comment to the
      calendar user.

   Format Definition:  This property is defined by the following
      notation:

       comment    = "COMMENT" commparam ":" text CRLF

       commparam  = *(
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

       COMMENT:The meeting really needs to include both ourselves
         and the customer. We can't hold this meeting without them.
         As a matter of fact\, the venue for the meeting ought to be at
         their site. - - John

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

3.8.1.6.  Geographic Position

   Property Name:  GEO

   Purpose:  This property specifies information related to the global
      position for the activity specified by a calendar component.


   Value Type:  FLOAT.  The value MUST be two SEMICOLON-separated FLOAT
      values.

   Property Parameters:  IANA and non-standard property parameters can
      be specified on this property.

   Conformance:  This property can be specified in "VEVENT" or "VTODO"
      calendar components.

   Description:  This property value specifies latitude and longitude,
      in that order (i.e., "LAT LON" ordering).  The longitude
      represents the location east or west of the prime meridian as a
      positive or negative real number, respectively.  The longitude and
      latitude values MAY be specified up to six decimal places, which
      will allow for accuracy to within one meter of geographical
      position.  Receiving applications MUST accept values of this
      precision and MAY truncate values of greater precision.

      Values for latitude and longitude shall be expressed as decimal
      fractions of degrees.  Whole degrees of latitude shall be
      represented by a two-digit decimal number ranging from 0 through
      90.  Whole degrees of longitude shall be represented by a decimal
      number ranging from 0 through 180.  When a decimal fraction of a
      degree is specified, it shall be separated from the whole number
      of degrees by a decimal point.

      Latitudes north of the equator shall be specified by a plus sign
      (+), or by the absence of a minus sign (-), preceding the digits
      designating degrees.  Latitudes south of the Equator shall be
      designated by a minus sign (-) preceding the digits designating
      degrees.  A point on the Equator shall be assigned to the Northern
      Hemisphere.

      Longitudes east of the prime meridian shall be specified by a plus
      sign (+), or by the absence of a minus sign (-), preceding the
      digits designating degrees.  Longitudes west of the meridian shall
      be designated by minus sign (-) preceding the digits designating
      degrees.  A point on the prime meridian shall be assigned to the
      Eastern Hemisphere.  A point on the 180th meridian shall be
      assigned to the Western Hemisphere.  One exception to this last
      convention is permitted.  For the special condition of describing
      a band of latitude around the earth, the East Bounding Coordinate
      data element shall be assigned the value +180 (180) degrees.

      Any spatial address with a latitude of +90 (90) or -90 degrees
      will specify the position at the North or South Pole,
      respectively.  The component for longitude may have any legal
      value.

      With the exception of the special condition described above, this
      form is specified in [ANSI INCITS 61-1986].

      The simple formula for converting degrees-minutes-seconds into
      decimal degrees is:

      decimal = degrees + minutes/60 + seconds/3600.

   Format Definition:  This property is defined by the following
      notation:

       geo        = "GEO" geoparam ":" geovalue CRLF

       geoparam   = *(";" other-param)

       geovalue   = float ";" float
       ;Latitude and Longitude components

   Example:  The following is an example of this property:

       GEO:37.386013;-122.082932

3.8.1.7.  Location

   Property Name:  LOCATION

   Purpose:  This property defines the intended venue for the activity
      defined by a calendar component.

   Value Type:  TEXT

   Property Parameters:  IANA, non-standard, alternate text
      representation, and language property parameters can be specified
      on this property.

   Conformance:  This property can be specified in "VEVENT" or "VTODO"
      calendar component.

   Description:  Specific venues such as conference or meeting rooms may
      be explicitly specified using this property.  An alternate
      representation may be specified that is a URI that points to
      directory information with more structured specification of the
      location.  For example, the alternate representation may specify
      either an LDAP URL [RFC4516] pointing to an LDAP server entry or a
      CID URL [RFC2392] pointing to a MIME body part containing a
      Virtual-Information Card (vCard) [RFC2426] for the location.



   Format Definition:  This property is defined by the following
      notation:

       location   = "LOCATION"  locparam ":" text CRLF

       locparam   = *(
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

   Example:  The following are some examples of this property:

       LOCATION:Conference Room - F123\, Bldg. 002

       LOCATION;ALTREP="http://xyzcorp.com/conf-rooms/f123.vcf":
        Conference Room - F123\, Bldg. 002

3.8.1.8.  Percent Complete

   Property Name:  PERCENT-COMPLETE

   Purpose:  This property is used by an assignee or delegatee of a
      to-do to convey the percent completion of a to-do to the
      "Organizer".

   Value Type:  INTEGER

   Property Parameters:  IANA and non-standard property parameters can
      be specified on this property.

   Conformance:  This property can be specified once in a "VTODO"
      calendar component.

   Description:  The property value is a positive integer between 0 and
      100.  A value of "0" indicates the to-do has not yet been started.
      A value of "100" indicates that the to-do has been completed.
      Integer values in between indicate the percent partially complete.



      When a to-do is assigned to multiple individuals, the property
      value indicates the percent complete for that portion of the to-do
      assigned to the assignee or delegatee.  For example, if a to-do is
      assigned to both individuals "A" and "B".  A reply from "A" with a
      percent complete of "70" indicates that "A" has completed 70% of
      the to-do assigned to them.  A reply from "B" with a percent
      complete of "50" indicates "B" has completed 50% of the to-do
      assigned to them.

   Format Definition:  This property is defined by the following
      notation:

       percent = "PERCENT-COMPLETE" pctparam ":" integer CRLF

       pctparam   = *(";" other-param)

   Example:  The following is an example of this property to show 39%
      completion:

       PERCENT-COMPLETE:39

3.8.1.9.  Priority

   Property Name:  PRIORITY

   Purpose:  This property defines the relative priority for a calendar
      component.

   Value Type:  INTEGER

   Property Parameters:  IANA and non-standard property parameters can
      be specified on this property.

   Conformance:  This property can be specified in "VEVENT" and "VTODO"
      calendar components.

   Description:  This priority is specified as an integer in the range 0
      to 9.  A value of 0 specifies an undefined priority.  A value of 1
      is the highest priority.  A value of 2 is the second highest
      priority.  Subsequent numbers specify a decreasing ordinal
      priority.  A value of 9 is the lowest priority.

      A CUA with a three-level priority scheme of "HIGH", "MEDIUM", and
      "LOW" is mapped into this property such that a property value in
      the range of 1 to 4 specifies "HIGH" priority.  A value of 5 is
      the normal or "MEDIUM" priority.  A value in the range of 6 to 9
      is "LOW" priority.


      A CUA with a priority schema of "A1", "A2", "A3", "B1", "B2", ...,
      "C3" is mapped into this property such that a property value of 1
      specifies "A1", a property value of 2 specifies "A2", a property
      value of 3 specifies "A3", and so forth up to a property value of
      9 specifies "C3".

      Other integer values are reserved for future use.

      Within a "VEVENT" calendar component, this property specifies a
      priority for the event.  This property may be useful when more
      than one event is scheduled for a given time period.

      Within a "VTODO" calendar component, this property specified a
      priority for the to-do.  This property is useful in prioritizing
      multiple action items for a given time period.

   Format Definition:  This property is defined by the following
      notation:

       priority   = "PRIORITY" prioparam ":" priovalue CRLF
       ;Default is zero (i.e., undefined).

       prioparam  = *(";" other-param)

       priovalue   = integer       ;Must be in the range [0..9]
          ; All other values are reserved for future use.

   Example:  The following is an example of a property with the highest
      priority:

       PRIORITY:1

      The following is an example of a property with a next highest
      priority:

       PRIORITY:2

      The following is an example of a property with no priority.  This
      is equivalent to not specifying the "PRIORITY" property:

       PRIORITY:0








3.8.1.10.  Resources

   Property Name:  RESOURCES

   Purpose:  This property defines the equipment or resources
      anticipated for an activity specified by a calendar component.

   Value Type:  TEXT

   Property Parameters:  IANA, non-standard, alternate text
      representation, and language property parameters can be specified
      on this property.

   Conformance:  This property can be specified once in "VEVENT" or
      "VTODO" calendar component.

   Description:  The property value is an arbitrary text.  More than one
      resource can be specified as a COMMA-separated list of resources.

   Format Definition:  This property is defined by the following
      notation:

       resources  = "RESOURCES" resrcparam ":" text *("," text) CRLF

       resrcparam = *(
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

       RESOURCES:EASEL,PROJECTOR,VCR

       RESOURCES;LANGUAGE=fr:Nettoyeur haute pression






3.8.1.11.  Status

   Property Name:  STATUS

   Purpose:  This property defines the overall status or confirmation
      for the calendar component.

   Value Type:  TEXT

   Property Parameters:  IANA and non-standard property parameters can
      be specified on this property.

   Conformance:  This property can be specified once in "VEVENT",
      "VTODO", or "VJOURNAL" calendar components.

   Description:  In a group-scheduled calendar component, the property
      is used by the "Organizer" to provide a confirmation of the event
      to the "Attendees".  For example in a "VEVENT" calendar component,
      the "Organizer" can indicate that a meeting is tentative,
      confirmed, or cancelled.  In a "VTODO" calendar component, the
      "Organizer" can indicate that an action item needs action, is
      completed, is in process or being worked on, or has been
      cancelled.  In a "VJOURNAL" calendar component, the "Organizer"
      can indicate that a journal entry is draft, final, or has been
      cancelled or removed.

   Format Definition:  This property is defined by the following
      notation:

       status          = "STATUS" statparam ":" statvalue CRLF

       statparam       = *(";" other-param)

       statvalue       = (statvalue-event
                       /  statvalue-todo
                       /  statvalue-jour)

       statvalue-event = "TENTATIVE"    ;Indicates event is tentative.
                       / "CONFIRMED"    ;Indicates event is definite.
                       / "CANCELLED"    ;Indicates event was cancelled.
       ;Status values for a "VEVENT"

       statvalue-todo  = "NEEDS-ACTION" ;Indicates to-do needs action.
                       / "COMPLETED"    ;Indicates to-do completed.
                       / "IN-PROCESS"   ;Indicates to-do in process of.
                       / "CANCELLED"    ;Indicates to-do was cancelled.
       ;Status values for "VTODO".


       statvalue-jour  = "DRAFT"        ;Indicates journal is draft.
                       / "FINAL"        ;Indicates journal is final.
                       / "CANCELLED"    ;Indicates journal is removed.
      ;Status values for "VJOURNAL".

   Example:  The following is an example of this property for a "VEVENT"
      calendar component:

       STATUS:TENTATIVE

      The following is an example of this property for a "VTODO"
      calendar component:

       STATUS:NEEDS-ACTION

      The following is an example of this property for a "VJOURNAL"
      calendar component:

       STATUS:DRAFT

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

3.8.2.  Date and Time Component Properties

   The following properties specify date and time related information in
   calendar components.

3.8.2.1.  Date-Time Completed

   Property Name:  COMPLETED

   Purpose:  This property defines the date and time that a to-do was
      actually completed.

   Value Type:  DATE-TIME

   Property Parameters:  IANA and non-standard property parameters can
      be specified on this property.

   Conformance:  The property can be specified in a "VTODO" calendar
      component.  The value MUST be specified as a date with UTC time.

   Description:  This property defines the date and time that a to-do
      was actually completed.

   Format Definition:  This property is defined by the following
      notation:




       completed  = "COMPLETED" compparam ":" date-time CRLF

       compparam  = *(";" other-param)

   Example:  The following is an example of this property:

    COMPLETED:19960401T150000Z

3.8.2.2.  Date-Time End

   Property Name:  DTEND

   Purpose:  This property specifies the date and time that a calendar
      component ends.

   Value Type:  The default value type is DATE-TIME.  The value type can
      be set to a DATE value type.

   Property Parameters:  IANA, non-standard, value data type, and time
      zone identifier property parameters can be specified on this
      property.

   Conformance:  This property can be specified in "VEVENT" or
      "VFREEBUSY" calendar components.

   Description:  Within the "VEVENT" calendar component, this property
      defines the date and time by which the event ends.  The value type
      of this property MUST be the same as the "DTSTART" property, and
      its value MUST be later in time than the value of the "DTSTART"
      property.  Furthermore, this property MUST be specified as a date
      with local time if and only if the "DTSTART" property is also
      specified as a date with local time.

      Within the "VFREEBUSY" calendar component, this property defines
      the end date and time for the free or busy time information.  The
      time MUST be specified in the UTC time format.  The value MUST be
      later in time than the value of the "DTSTART" property.

   Format Definition:  This property is defined by the following
      notation:









       dtend      = "DTEND" dtendparam ":" dtendval CRLF

       dtendparam = *(
                  ;
                  ; The following are OPTIONAL,
                  ; but MUST NOT occur more than once.
                  ;
                  (";" "VALUE" "=" ("DATE-TIME" / "DATE")) /
                  (";" tzidparam) /
                  ;
                  ; The following is OPTIONAL,
                  ; and MAY occur more than once.
                  ;
                  (";" other-param)
                  ;
                  )

       dtendval   = date-time / date
       ;Value MUST match value type

   Example:  The following is an example of this property:

       DTEND:19960401T150000Z

       DTEND;VALUE=DATE:19980704

3.8.2.3.  Date-Time Due

   Property Name:  DUE

   Purpose:  This property defines the date and time that a to-do is
      expected to be completed.

   Value Type:  The default value type is DATE-TIME.  The value type can
      be set to a DATE value type.

   Property Parameters:  IANA, non-standard, value data type, and time
      zone identifier property parameters can be specified on this
      property.

   Conformance:  The property can be specified once in a "VTODO"
      calendar component.

   Description:  This property defines the date and time before which a
      to-do is expected to be completed.  For cases where this property
      is specified in a "VTODO" calendar component that also specifies a
      "DTSTART" property, the value type of this property MUST be the
      same as the "DTSTART" property, and the value of this property

      MUST be later in time than the value of the "DTSTART" property.
      Furthermore, this property MUST be specified as a date with local
      time if and only if the "DTSTART" property is also specified as a
      date with local time.

   Format Definition:  This property is defined by the following
      notation:

       due        = "DUE" dueparam ":" dueval CRLF

       dueparam   = *(
                  ;
                  ; The following are OPTIONAL,
                  ; but MUST NOT occur more than once.
                  ;
                  (";" "VALUE" "=" ("DATE-TIME" / "DATE")) /
                  (";" tzidparam) /
                  ;
                  ; The following is OPTIONAL,
                  ; and MAY occur more than once.
                  ;
                  (";" other-param)
                  ;
                  )

       dueval     = date-time / date
       ;Value MUST match value type

   Example:  The following is an example of this property:

       DUE:19980430T000000Z

3.8.2.4.  Date-Time Start

   Property Name:  DTSTART

   Purpose:  This property specifies when the calendar component begins.

   Value Type:  The default value type is DATE-TIME.  The time value
      MUST be one of the forms defined for the DATE-TIME value type.
      The value type can be set to a DATE value type.

   Property Parameters:  IANA, non-standard, value data type, and time
      zone identifier property parameters can be specified on this
      property.

   Conformance:  This property can be specified once in the "VEVENT",
      "VTODO", or "VFREEBUSY" calendar components as well as in the

      "STANDARD" and "DAYLIGHT" sub-components.  This property is
      REQUIRED in all types of recurring calendar components that
      specify the "RRULE" property.  This property is also REQUIRED in
      "VEVENT" calendar components contained in iCalendar objects that
      don't specify the "METHOD" property.

   Description:  Within the "VEVENT" calendar component, this property
      defines the start date and time for the event.

      Within the "VFREEBUSY" calendar component, this property defines
      the start date and time for the free or busy time information.
      The time MUST be specified in UTC time.

      Within the "STANDARD" and "DAYLIGHT" sub-components, this property
      defines the effective start date and time for a time zone
      specification.  This property is REQUIRED within each "STANDARD"
      and "DAYLIGHT" sub-components included in "VTIMEZONE" calendar
      components and MUST be specified as a date with local time without
      the "TZID" property parameter.

   Format Definition:  This property is defined by the following
      notation:

       dtstart    = "DTSTART" dtstparam ":" dtstval CRLF

       dtstparam  = *(
                  ;
                  ; The following are OPTIONAL,
                  ; but MUST NOT occur more than once.
                  ;
                  (";" "VALUE" "=" ("DATE-TIME" / "DATE")) /
                  (";" tzidparam) /
                  ;
                  ; The following is OPTIONAL,
                  ; and MAY occur more than once.
                  ;
                  (";" other-param)
                  ;
                  )

       dtstval    = date-time / date
       ;Value MUST match value type

   Example:  The following is an example of this property:

       DTSTART:19980118T073000Z



3.8.2.5.  Duration

   Property Name:  DURATION

   Purpose:  This property specifies a positive duration of time.

   Value Type:  DURATION

   Property Parameters:  IANA and non-standard property parameters can
      be specified on this property.

   Conformance:  This property can be specified in "VEVENT", "VTODO", or
      "VALARM" calendar components.

   Description:  In a "VEVENT" calendar component the property may be
      used to specify a duration of the event, instead of an explicit
      end DATE-TIME.  In a "VTODO" calendar component the property may
      be used to specify a duration for the to-do, instead of an
      explicit due DATE-TIME.  In a "VALARM" calendar component the
      property may be used to specify the delay period prior to
      repeating an alarm.  When the "DURATION" property relates to a
      "DTSTART" property that is specified as a DATE value, then the
      "DURATION" property MUST be specified as a "dur-day" or "dur-week"
      value.

   Format Definition:  This property is defined by the following
      notation:

       duration   = "DURATION" durparam ":" dur-value CRLF
                    ;consisting of a positive duration of time.

       durparam   = *(";" other-param)

   Example:  The following is an example of this property that specifies
      an interval of time of one hour and zero minutes and zero seconds:

       DURATION:PT1H0M0S

      The following is an example of this property that specifies an
      interval of time of 15 minutes.

       DURATION:PT15M







3.8.2.6.  Free/Busy Time

   Property Name:  FREEBUSY

   Purpose:  This property defines one or more free or busy time
      intervals.

   Value Type:  PERIOD

   Property Parameters:  IANA, non-standard, and free/busy time type
      property parameters can be specified on this property.

   Conformance:  The property can be specified in a "VFREEBUSY" calendar
      component.

   Description:  These time periods can be specified as either a start
      and end DATE-TIME or a start DATE-TIME and DURATION.  The date and
      time MUST be a UTC time format.

      "FREEBUSY" properties within the "VFREEBUSY" calendar component
      SHOULD be sorted in ascending order, based on start time and then
      end time, with the earliest periods first.

      The "FREEBUSY" property can specify more than one value, separated
      by the COMMA character.  In such cases, the "FREEBUSY" property
      values MUST all be of the same "FBTYPE" property parameter type
      (e.g., all values of a particular "FBTYPE" listed together in a
      single property).

   Format Definition:  This property is defined by the following
      notation:

       freebusy   = "FREEBUSY" fbparam ":" fbvalue CRLF

       fbparam    = *(
                  ;
                  ; The following is OPTIONAL,
                  ; but MUST NOT occur more than once.
                  ;
                  (";" fbtypeparam) /
                  ;
                  ; The following is OPTIONAL,
                  ; and MAY occur more than once.
                  ;
                  (";" other-param)
                  ;
                  )


       fbvalue    = period *("," period)
       ;Time value MUST be in the UTC time format.

   Example:  The following are some examples of this property:

       FREEBUSY;FBTYPE=BUSY-UNAVAILABLE:19970308T160000Z/PT8H30M

       FREEBUSY;FBTYPE=FREE:19970308T160000Z/PT3H,19970308T200000Z/PT1H

       FREEBUSY;FBTYPE=FREE:19970308T160000Z/PT3H,19970308T200000Z/PT1H
        ,19970308T230000Z/19970309T000000Z

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

3.8.3.  Time Zone Component Properties

   The following properties specify time zone information in calendar
   components.

3.8.3.1.  Time Zone Identifier

   Property Name:  TZID

   Purpose:  This property specifies the text value that uniquely
      identifies the "VTIMEZONE" calendar component in the scope of an
      iCalendar object.

   Value Type:  TEXT

   Property Parameters:  IANA and non-standard property parameters can
      be specified on this property.

   Conformance:  This property MUST be specified in a "VTIMEZONE"
      calendar component.

   Description:  This is the label by which a time zone calendar
      component is referenced by any iCalendar properties whose value
      type is either DATE-TIME or TIME and not intended to specify a UTC
      or a "floating" time.  The presence of the SOLIDUS character as a
      prefix, indicates that this "TZID" represents an unique ID in a
      globally defined time zone registry (when such registry is
      defined).

         Note: This document does not define a naming convention for
         time zone identifiers.  Implementers may want to use the naming
         conventions defined in existing time zone specifications such
         as the public-domain TZ database [TZDB].  The specification of
         globally unique time zone identifiers is not addressed by this
         document and is left for future study.




   Format Definition:  This property is defined by the following
      notation:

       tzid       = "TZID" tzidpropparam ":" [tzidprefix] text CRLF

       tzidpropparam      = *(";" other-param)

       ;tzidprefix        = "/"
       ; Defined previously. Just listed here for reader convenience.

   Example:  The following are examples of non-globally unique time zone
      identifiers:

       TZID:America/New_York

       TZID:America/Los_Angeles

      The following is an example of a fictitious globally unique time
      zone identifier:

       TZID:/example.org/America/New_York

3.8.3.2.  Time Zone Name

   Property Name:  TZNAME

   Purpose:  This property specifies the customary designation for a
      time zone description.

   Value Type:  TEXT

   Property Parameters:  IANA, non-standard, and language property
      parameters can be specified on this property.

   Conformance:  This property can be specified in "STANDARD" and
      "DAYLIGHT" sub-components.

   Description:  This property specifies a customary name that can be
      used when displaying dates that occur during the observance
      defined by the time zone sub-component.

   Format Definition:  This property is defined by the following
      notation:






       tzname     = "TZNAME" tznparam ":" text CRLF

       tznparam   = *(
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

   Example:  The following are examples of this property:

       TZNAME:EST

       TZNAME;LANGUAGE=fr-CA:HNE

3.8.3.3.  Time Zone Offset From

   Property Name:  TZOFFSETFROM

   Purpose:  This property specifies the offset that is in use prior to
      this time zone observance.

   Value Type:  UTC-OFFSET

   Property Parameters:  IANA and non-standard property parameters can
      be specified on this property.

   Conformance:  This property MUST be specified in "STANDARD" and
      "DAYLIGHT" sub-components.

   Description:  This property specifies the offset that is in use prior
      to this time observance.  It is used to calculate the absolute
      time at which the transition to a given observance takes place.
      This property MUST only be specified in a "VTIMEZONE" calendar
      component.  A "VTIMEZONE" calendar component MUST include this
      property.  The property value is a signed numeric indicating the
      number of hours and possibly minutes from UTC.  Positive numbers
      represent time zones east of the prime meridian, or ahead of UTC.
      Negative numbers represent time zones west of the prime meridian,
      or behind UTC.


   Format Definition:  This property is defined by the following
      notation:

       tzoffsetfrom       = "TZOFFSETFROM" frmparam ":" utc-offset
                            CRLF

       frmparam   = *(";" other-param)

   Example:  The following are examples of this property:

       TZOFFSETFROM:-0500

       TZOFFSETFROM:+1345

3.8.3.4.  Time Zone Offset To

   Property Name:  TZOFFSETTO

   Purpose:  This property specifies the offset that is in use in this
      time zone observance.

   Value Type:  UTC-OFFSET

   Property Parameters:  IANA and non-standard property parameters can
      be specified on this property.

   Conformance:  This property MUST be specified in "STANDARD" and
      "DAYLIGHT" sub-components.

   Description:  This property specifies the offset that is in use in
      this time zone observance.  It is used to calculate the absolute
      time for the new observance.  The property value is a signed
      numeric indicating the number of hours and possibly minutes from
      UTC.  Positive numbers represent time zones east of the prime
      meridian, or ahead of UTC.  Negative numbers represent time zones
      west of the prime meridian, or behind UTC.

   Format Definition:  This property is defined by the following
      notation:

       tzoffsetto = "TZOFFSETTO" toparam ":" utc-offset CRLF

       toparam    = *(";" other-param)






   Example:  The following are examples of this property:

       TZOFFSETTO:-0400

       TZOFFSETTO:+1245

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

3.8.4.  Relationship Component Properties

   The following properties specify relationship information in calendar
   components.




3.8.4.1.  Attendee

   Property Name:  ATTENDEE

   Purpose:  This property defines an "Attendee" within a calendar
      component.

   Value Type:  CAL-ADDRESS

   Property Parameters:  IANA, non-standard, language, calendar user
      type, group or list membership, participation role, participation
      status, RSVP expectation, delegatee, delegator, sent by, common
      name, or directory entry reference property parameters can be
      specified on this property.

   Conformance:  This property MUST be specified in an iCalendar object
      that specifies a group-scheduled calendar entity.  This property
      MUST NOT be specified in an iCalendar object when publishing the
      calendar information (e.g., NOT in an iCalendar object that
      specifies the publication of a calendar user's busy time, event,
      to-do, or journal).  This property is not specified in an
      iCalendar object that specifies only a time zone definition or
      that defines calendar components that are not group-scheduled
      components, but are components only on a single user's calendar.

   Description:  This property MUST only be specified within calendar
      components to specify participants, non-participants, and the
      chair of a group-scheduled calendar entity.  The property is
      specified within an "EMAIL" category of the "VALARM" calendar
      component to specify an email address that is to receive the email
      type of iCalendar alarm.

      The property parameter "CN" is for the common or displayable name
      associated with the calendar address; "ROLE", for the intended
      role that the attendee will have in the calendar component;
      "PARTSTAT", for the status of the attendee's participation;
      "RSVP", for indicating whether the favor of a reply is requested;
      "CUTYPE", to indicate the type of calendar user; "MEMBER", to
      indicate the groups that the attendee belongs to; "DELEGATED-TO",
      to indicate the calendar users that the original request was
      delegated to; and "DELEGATED-FROM", to indicate whom the request
      was delegated from; "SENT-BY", to indicate whom is acting on
      behalf of the "ATTENDEE"; and "DIR", to indicate the URI that
      points to the directory information corresponding to the attendee.
      These property parameters can be specified on an "ATTENDEE"
      property in either a "VEVENT", "VTODO", or "VJOURNAL" calendar
      component.  They MUST NOT be specified in an "ATTENDEE" property
      in a "VFREEBUSY" or "VALARM" calendar component.  If the

      "LANGUAGE" property parameter is specified, the identified
      language applies to the "CN" parameter.

      A recipient delegated a request MUST inherit the "RSVP" and "ROLE"
      values from the attendee that delegated the request to them.

      Multiple attendees can be specified by including multiple
      "ATTENDEE" properties within the calendar component.

   Format Definition:  This property is defined by the following
      notation:

       attendee   = "ATTENDEE" attparam ":" cal-address CRLF

       attparam   = *(
                  ;
                  ; The following are OPTIONAL,
                  ; but MUST NOT occur more than once.
                  ;
                  (";" cutypeparam) / (";" memberparam) /
                  (";" roleparam) / (";" partstatparam) /
                  (";" rsvpparam) / (";" deltoparam) /
                  (";" delfromparam) / (";" sentbyparam) /
                  (";" cnparam) / (";" dirparam) /
                  (";" languageparam) /
                  ;
                  ; The following is OPTIONAL,
                  ; and MAY occur more than once.
                  ;
                  (";" other-param)
                  ;
                  )

   Example:  The following are examples of this property's use for a
      to-do:

       ATTENDEE;MEMBER="mailto:DEV-GROUP@example.com":
        mailto:joecool@example.com
       ATTENDEE;DELEGATED-FROM="mailto:immud@example.com":
        mailto:ildoit@example.com









      The following is an example of this property used for specifying
      multiple attendees to an event:

       ATTENDEE;ROLE=REQ-PARTICIPANT;PARTSTAT=TENTATIVE;CN=Henry
        Cabot:mailto:hcabot@example.com
       ATTENDEE;ROLE=REQ-PARTICIPANT;DELEGATED-FROM="mailto:bob@
        example.com";PARTSTAT=ACCEPTED;CN=Jane Doe:mailto:jdoe@
        example.com

      The following is an example of this property with a URI to the
      directory information associated with the attendee:

       ATTENDEE;CN=John Smith;DIR="ldap://example.com:6666/o=ABC%
        20Industries,c=US???(cn=Jim%20Dolittle)":mailto:jimdo@
        example.com

      The following is an example of this property with "delegatee" and
      "delegator" information for an event:

       ATTENDEE;ROLE=REQ-PARTICIPANT;PARTSTAT=TENTATIVE;DELEGATED-FROM=
        "mailto:iamboss@example.com";CN=Henry Cabot:mailto:hcabot@
        example.com
       ATTENDEE;ROLE=NON-PARTICIPANT;PARTSTAT=DELEGATED;DELEGATED-TO=
        "mailto:hcabot@example.com";CN=The Big Cheese:mailto:iamboss
        @example.com
       ATTENDEE;ROLE=REQ-PARTICIPANT;PARTSTAT=ACCEPTED;CN=Jane Doe
        :mailto:jdoe@example.com

   Example:  The following is an example of this property's use when
      another calendar user is acting on behalf of the "Attendee":

       ATTENDEE;SENT-BY=mailto:jan_doe@example.com;CN=John Smith:
        mailto:jsmith@example.com

3.8.4.2.  Contact

   Property Name:  CONTACT

   Purpose:  This property is used to represent contact information or
      alternately a reference to contact information associated with the
      calendar component.

   Value Type:  TEXT

   Property Parameters:  IANA, non-standard, alternate text
      representation, and language property parameters can be specified
      on this property.


   Conformance:  This property can be specified in a "VEVENT", "VTODO",
      "VJOURNAL", or "VFREEBUSY" calendar component.

   Description:  The property value consists of textual contact
      information.  An alternative representation for the property value
      can also be specified that refers to a URI pointing to an
      alternate form, such as a vCard [RFC2426], for the contact
      information.

   Format Definition:  This property is defined by the following
      notation:

       contact    = "CONTACT" contparam ":" text CRLF

       contparam  = *(
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

   Example:  The following is an example of this property referencing
      textual contact information:

       CONTACT:Jim Dolittle\, ABC Industries\, +1-919-555-1234

      The following is an example of this property with an alternate
      representation of an LDAP URI to a directory entry containing the
      contact information:

       CONTACT;ALTREP="ldap://example.com:6666/o=ABC%20Industries\,
        c=US???(cn=Jim%20Dolittle)":Jim Dolittle\, ABC Industries\,
        +1-919-555-1234

      The following is an example of this property with an alternate
      representation of a MIME body part containing the contact
      information, such as a vCard [RFC2426] embedded in a text/
      directory media type [RFC2425]:

       CONTACT;ALTREP="CID:part3.msg970930T083000SILVER@example.com":
        Jim Dolittle\, ABC Industries\, +1-919-555-1234

      The following is an example of this property referencing a network
      resource, such as a vCard [RFC2426] object containing the contact
      information:

       CONTACT;ALTREP="http://example.com/pdi/jdoe.vcf":Jim
         Dolittle\, ABC Industries\, +1-919-555-1234

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

3.8.4.4.  Recurrence ID

   Property Name:  RECURRENCE-ID

   Purpose:  This property is used in conjunction with the "UID" and
      "SEQUENCE" properties to identify a specific instance of a
      recurring "VEVENT", "VTODO", or "VJOURNAL" calendar component.
      The property value is the original value of the "DTSTART" property
      of the recurrence instance.

   Value Type:  The default value type is DATE-TIME.  The value type can
      be set to a DATE value type.  This property MUST have the same
      value type as the "DTSTART" property contained within the
      recurring component.  Furthermore, this property MUST be specified
      as a date with local time if and only if the "DTSTART" property
      contained within the recurring component is specified as a date
      with local time.

   Property Parameters:  IANA, non-standard, value data type, time zone
      identifier, and recurrence identifier range parameters can be
      specified on this property.

   Conformance:  This property can be specified in an iCalendar object
      containing a recurring calendar component.

   Description:  The full range of calendar components specified by a
      recurrence set is referenced by referring to just the "UID"
      property value corresponding to the calendar component.  The
      "RECURRENCE-ID" property allows the reference to an individual
      instance within the recurrence set.

      If the value of the "DTSTART" property is a DATE type value, then
      the value MUST be the calendar date for the recurrence instance.

      The DATE-TIME value is set to the time when the original
      recurrence instance would occur; meaning that if the intent is to
      change a Friday meeting to Thursday, the DATE-TIME is still set to
      the original Friday meeting.

      The "RECURRENCE-ID" property is used in conjunction with the "UID"
      and "SEQUENCE" properties to identify a particular instance of a
      recurring event, to-do, or journal.  For a given pair of "UID" and
      "SEQUENCE" property values, the "RECURRENCE-ID" value for a
      recurrence instance is fixed.

      The "RANGE" parameter is used to specify the effective range of
      recurrence instances from the instance specified by the
      "RECURRENCE-ID" property value.  The value for the range parameter
      can only be "THISANDFUTURE" to indicate a range defined by the
      given recurrence instance and all subsequent instances.
      Subsequent instances are determined by their "RECURRENCE-ID" value
      and not their current scheduled start time.  Subsequent instances
      defined in separate components are not impacted by the given
      recurrence instance.  When the given recurrence instance is
      rescheduled, all subsequent instances are also rescheduled by the
      same time difference.  For instance, if the given recurrence
      instance is rescheduled to start 2 hours later, then all
      subsequent instances are also rescheduled 2 hours later.

      Similarly, if the duration of the given recurrence instance is
      modified, then all subsequence instances are also modified to have
      this same duration.

         Note: The "RANGE" parameter may not be appropriate to
         reschedule specific subsequent instances of complex recurring
         calendar component.  Assuming an unbounded recurring calendar
         component scheduled to occur on Mondays and Wednesdays, the
         "RANGE" parameter could not be used to reschedule only the
         future Monday instances to occur on Tuesday instead.  In such
         cases, the calendar application could simply truncate the
         unbounded recurring calendar component (i.e., with the "COUNT"
         or "UNTIL" rule parts), and create two new unbounded recurring
         calendar components for the future instances.

   Format Definition:  This property is defined by the following
      notation:

       recurid    = "RECURRENCE-ID" ridparam ":" ridval CRLF

       ridparam   = *(
                  ;
                  ; The following are OPTIONAL,
                  ; but MUST NOT occur more than once.
                  ;
                  (";" "VALUE" "=" ("DATE-TIME" / "DATE")) /
                  (";" tzidparam) / (";" rangeparam) /
                  ;
                  ; The following is OPTIONAL,
                  ; and MAY occur more than once.
                  ;
                  (";" other-param)
                  ;
                  )

       ridval     = date-time / date
       ;Value MUST match value type

   Example:  The following are examples of this property:

       RECURRENCE-ID;VALUE=DATE:19960401

       RECURRENCE-ID;RANGE=THISANDFUTURE:19960120T120000Z






3.8.4.5.  Related To

   Property Name:  RELATED-TO

   Purpose:  This property is used to represent a relationship or
      reference between one calendar component and another.

   Value Type:  TEXT

   Property Parameters:  IANA, non-standard, and relationship type
      property parameters can be specified on this property.

   Conformance:  This property can be specified in the "VEVENT",
      "VTODO", and "VJOURNAL" calendar components.

   Description:  The property value consists of the persistent, globally
      unique identifier of another calendar component.  This value would
      be represented in a calendar component by the "UID" property.

      By default, the property value points to another calendar
      component that has a PARENT relationship to the referencing
      object.  The "RELTYPE" property parameter is used to either
      explicitly state the default PARENT relationship type to the
      referenced calendar component or to override the default PARENT
      relationship type and specify either a CHILD or SIBLING
      relationship.  The PARENT relationship indicates that the calendar
      component is a subordinate of the referenced calendar component.
      The CHILD relationship indicates that the calendar component is a
      superior of the referenced calendar component.  The SIBLING
      relationship indicates that the calendar component is a peer of
      the referenced calendar component.

      Changes to a calendar component referenced by this property can
      have an implicit impact on the related calendar component.  For
      example, if a group event changes its start or end date or time,
      then the related, dependent events will need to have their start
      and end dates changed in a corresponding way.  Similarly, if a
      PARENT calendar component is cancelled or deleted, then there is
      an implied impact to the related CHILD calendar components.  This
      property is intended only to provide information on the
      relationship of calendar components.  It is up to the target
      calendar system to maintain any property implications of this
      relationship.






   Format Definition:  This property is defined by the following
      notation:

       related    = "RELATED-TO" relparam ":" text CRLF

       relparam   = *(
                  ;
                  ; The following is OPTIONAL,
                  ; but MUST NOT occur more than once.
                  ;
                  (";" reltypeparam) /
                  ;
                  ; The following is OPTIONAL,
                  ; and MAY occur more than once.
                  ;
                  (";" other-param)
                  ;
                  )

      The following is an example of this property:

       RELATED-TO:jsmith.part7.19960817T083000.xyzMail@example.com

       RELATED-TO:19960401-080045-4000F192713-0052@example.com

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

3.8.4.7.  Unique Identifier

   Property Name:  UID

   Purpose:  This property defines the persistent, globally unique
      identifier for the calendar component.

   Value Type:  TEXT

   Property Parameters:  IANA and non-standard property parameters can
      be specified on this property.

   Conformance:  The property MUST be specified in the "VEVENT",
      "VTODO", "VJOURNAL", or "VFREEBUSY" calendar components.

   Description:  The "UID" itself MUST be a globally unique identifier.
      The generator of the identifier MUST guarantee that the identifier
      is unique.  There are several algorithms that can be used to
      accomplish this.  A good method to assure uniqueness is to put the
      domain name or a domain literal IP address of the host on which
      the identifier was created on the right-hand side of an "@", and
      on the left-hand side, put a combination of the current calendar
      date and time of day (i.e., formatted in as a DATE-TIME value)
      along with some other currently unique (perhaps sequential)
      identifier available on the system (for example, a process id
      number).  Using a DATE-TIME value on the left-hand side and a
      domain name or domain literal on the right-hand side makes it
      possible to guarantee uniqueness since no two hosts should be
      using the same domain name or IP address at the same time.  Though
      other algorithms will work, it is RECOMMENDED that the right-hand
      side contain some domain identifier (either of the host itself or
      otherwise) such that the generator of the message identifier can
      guarantee the uniqueness of the left-hand side within the scope of
      that domain.

      This is the method for correlating scheduling messages with the
      referenced "VEVENT", "VTODO", or "VJOURNAL" calendar component.

      The full range of calendar components specified by a recurrence
      set is referenced by referring to just the "UID" property value
      corresponding to the calendar component.  The "RECURRENCE-ID"
      property allows the reference to an individual instance within the
      recurrence set.

      This property is an important method for group-scheduling
      applications to match requests with later replies, modifications,
      or deletion requests.  Calendaring and scheduling applications
      MUST generate this property in "VEVENT", "VTODO", and "VJOURNAL"
      calendar components to assure interoperability with other group-
      scheduling applications.  This identifier is created by the
      calendar system that generates an iCalendar object.

      Implementations MUST be able to receive and persist values of at
      least 255 octets for this property, but they MUST NOT truncate
      values in the middle of a UTF-8 multi-octet sequence.

   Format Definition:  This property is defined by the following
      notation:

       uid        = "UID" uidparam ":" text CRLF

       uidparam   = *(";" other-param)

   Example:  The following is an example of this property:

       UID:19960401T080045Z-4000F192713-0052@example.com

3.8.5.  Recurrence Component Properties

   The following properties specify recurrence information in calendar
   components.

3.8.5.1.  Exception Date-Times

   Property Name:  EXDATE

   Purpose:  This property defines the list of DATE-TIME exceptions for
      recurring events, to-dos, journal entries, or time zone
      definitions.

   Value Type:  The default value type for this property is DATE-TIME.
      The value type can be set to DATE.

   Property Parameters:  IANA, non-standard, value data type, and time
      zone identifier property parameters can be specified on this
      property.

   Conformance:  This property can be specified in recurring "VEVENT",
      "VTODO", and "VJOURNAL" calendar components as well as in the
      "STANDARD" and "DAYLIGHT" sub-components of the "VTIMEZONE"
      calendar component.

   Description:  The exception dates, if specified, are used in
      computing the recurrence set.  The recurrence set is the complete
      set of recurrence instances for a calendar component.  The
      recurrence set is generated by considering the initial "DTSTART"
      property along with the "RRULE", "RDATE", and "EXDATE" properties
      contained within the recurring component.  The "DTSTART" property
      defines the first instance in the recurrence set.  The "DTSTART"
      property value SHOULD match the pattern of the recurrence rule, if
      specified.  The recurrence set generated with a "DTSTART" property
      value that doesn't match the pattern of the rule is undefined.
      The final recurrence set is generated by gathering all of the
      start DATE-TIME values generated by any of the specified "RRULE"
      and "RDATE" properties, and then excluding any start DATE-TIME
      values specified by "EXDATE" properties.  This implies that start
      DATE-TIME values specified by "EXDATE" properties take precedence
      over those specified by inclusion properties (i.e., "RDATE" and
      "RRULE").  When duplicate instances are generated by the "RRULE"
      and "RDATE" properties, only one recurrence is considered.
      Duplicate instances are ignored.

      The "EXDATE" property can be used to exclude the value specified
      in "DTSTART".  However, in such cases, the original "DTSTART" date
      MUST still be maintained by the calendaring and scheduling system
      because the original "DTSTART" value has inherent usage
      dependencies by other properties such as the "RECURRENCE-ID".

   Format Definition:  This property is defined by the following
      notation:
















       exdate     = "EXDATE" exdtparam ":" exdtval *("," exdtval) CRLF

       exdtparam  = *(
                  ;
                  ; The following are OPTIONAL,
                  ; but MUST NOT occur more than once.
                  ;
                  (";" "VALUE" "=" ("DATE-TIME" / "DATE")) /
                  ;
                  (";" tzidparam) /
                  ;
                  ; The following is OPTIONAL,
                  ; and MAY occur more than once.
                  ;
                  (";" other-param)
                  ;
                  )

       exdtval    = date-time / date
       ;Value MUST match value type

   Example:  The following is an example of this property:

       EXDATE:19960402T010000Z,19960403T010000Z,19960404T010000Z

3.8.5.2.  Recurrence Date-Times

   Property Name:  RDATE

   Purpose:  This property defines the list of DATE-TIME values for
      recurring events, to-dos, journal entries, or time zone
      definitions.

   Value Type:  The default value type for this property is DATE-TIME.
      The value type can be set to DATE or PERIOD.

   Property Parameters:  IANA, non-standard, value data type, and time
      zone identifier property parameters can be specified on this
      property.

   Conformance:  This property can be specified in recurring "VEVENT",
      "VTODO", and "VJOURNAL" calendar components as well as in the
      "STANDARD" and "DAYLIGHT" sub-components of the "VTIMEZONE"
      calendar component.

   Description:  This property can appear along with the "RRULE"
      property to define an aggregate set of repeating occurrences.
      When they both appear in a recurring component, the recurrence

      instances are defined by the union of occurrences defined by both
      the "RDATE" and "RRULE".

      The recurrence dates, if specified, are used in computing the
      recurrence set.  The recurrence set is the complete set of
      recurrence instances for a calendar component.  The recurrence set
      is generated by considering the initial "DTSTART" property along
      with the "RRULE", "RDATE", and "EXDATE" properties contained
      within the recurring component.  The "DTSTART" property defines
      the first instance in the recurrence set.  The "DTSTART" property
      value SHOULD match the pattern of the recurrence rule, if
      specified.  The recurrence set generated with a "DTSTART" property
      value that doesn't match the pattern of the rule is undefined.
      The final recurrence set is generated by gathering all of the
      start DATE-TIME values generated by any of the specified "RRULE"
      and "RDATE" properties, and then excluding any start DATE-TIME
      values specified by "EXDATE" properties.  This implies that start
      DATE-TIME values specified by "EXDATE" properties take precedence
      over those specified by inclusion properties (i.e., "RDATE" and
      "RRULE").  Where duplicate instances are generated by the "RRULE"
      and "RDATE" properties, only one recurrence is considered.
      Duplicate instances are ignored.

   Format Definition:  This property is defined by the following
      notation:

       rdate      = "RDATE" rdtparam ":" rdtval *("," rdtval) CRLF

       rdtparam   = *(
                  ;
                  ; The following are OPTIONAL,
                  ; but MUST NOT occur more than once.
                  ;
                  (";" "VALUE" "=" ("DATE-TIME" / "DATE" / "PERIOD")) /
                  (";" tzidparam) /
                  ;
                  ; The following is OPTIONAL,
                  ; and MAY occur more than once.
                  ;
                  (";" other-param)
                  ;
                  )

       rdtval     = date-time / date / period
       ;Value MUST match value type




   Example:  The following are examples of this property:

       RDATE:19970714T123000Z
       RDATE;TZID=America/New_York:19970714T083000

       RDATE;VALUE=PERIOD:19960403T020000Z/19960403T040000Z,
        19960404T010000Z/PT3H

       RDATE;VALUE=DATE:19970101,19970120,19970217,19970421
        19970526,19970704,19970901,19971014,19971128,19971129,19971225

3.8.5.3.  Recurrence Rule

   Property Name:  RRULE

   Purpose:  This property defines a rule or repeating pattern for
      recurring events, to-dos, journal entries, or time zone
      definitions.

   Value Type:  RECUR

   Property Parameters:  IANA and non-standard property parameters can
      be specified on this property.

   Conformance:  This property can be specified in recurring "VEVENT",
      "VTODO", and "VJOURNAL" calendar components as well as in the
      "STANDARD" and "DAYLIGHT" sub-components of the "VTIMEZONE"
      calendar component, but it SHOULD NOT be specified more than once.
      The recurrence set generated with multiple "RRULE" properties is
      undefined.

   Description:  The recurrence rule, if specified, is used in computing
      the recurrence set.  The recurrence set is the complete set of
      recurrence instances for a calendar component.  The recurrence set
      is generated by considering the initial "DTSTART" property along
      with the "RRULE", "RDATE", and "EXDATE" properties contained
      within the recurring component.  The "DTSTART" property defines
      the first instance in the recurrence set.  The "DTSTART" property
      value SHOULD be synchronized with the recurrence rule, if
      specified.  The recurrence set generated with a "DTSTART" property
      value not synchronized with the recurrence rule is undefined.  The
      final recurrence set is generated by gathering all of the start
      DATE-TIME values generated by any of the specified "RRULE" and
      "RDATE" properties, and then excluding any start DATE-TIME values
      specified by "EXDATE" properties.  This implies that start DATE-
      TIME values specified by "EXDATE" properties take precedence over
      those specified by inclusion properties (i.e., "RDATE" and
      "RRULE").  Where duplicate instances are generated by the "RRULE"

      and "RDATE" properties, only one recurrence is considered.
      Duplicate instances are ignored.

      The "DTSTART" property specified within the iCalendar object
      defines the first instance of the recurrence.  In most cases, a
      "DTSTART" property of DATE-TIME value type used with a recurrence
      rule, should be specified as a date with local time and time zone
      reference to make sure all the recurrence instances start at the
      same local time regardless of time zone changes.

      If the duration of the recurring component is specified with the
      "DTEND" or "DUE" property, then the same exact duration will apply
      to all the members of the generated recurrence set.  Else, if the
      duration of the recurring component is specified with the
      "DURATION" property, then the same nominal duration will apply to
      all the members of the generated recurrence set and the exact
      duration of each recurrence instance will depend on its specific
      start time.  For example, recurrence instances of a nominal
      duration of one day will have an exact duration of more or less
      than 24 hours on a day where a time zone shift occurs.  The
      duration of a specific recurrence may be modified in an exception
      component or simply by using an "RDATE" property of PERIOD value
      type.

   Format Definition:  This property is defined by the following
      notation:

       rrule      = "RRULE" rrulparam ":" recur CRLF

       rrulparam  = *(";" other-param)

   Example:  All examples assume the Eastern United States time zone.

      Daily for 10 occurrences:

       DTSTART;TZID=America/New_York:19970902T090000
       RRULE:FREQ=DAILY;COUNT=10

       ==> (1997 9:00 AM EDT) September 2-11

      Daily until December 24, 1997:

       DTSTART;TZID=America/New_York:19970902T090000
       RRULE:FREQ=DAILY;UNTIL=19971224T000000Z

       ==> (1997 9:00 AM EDT) September 2-30;October 1-25
           (1997 9:00 AM EST) October 26-31;November 1-30;December 1-23


      Every other day - forever:

       DTSTART;TZID=America/New_York:19970902T090000
       RRULE:FREQ=DAILY;INTERVAL=2

       ==> (1997 9:00 AM EDT) September 2,4,6,8...24,26,28,30;
                              October 2,4,6...20,22,24
           (1997 9:00 AM EST) October 26,28,30;
                              November 1,3,5,7...25,27,29;
                              December 1,3,...

      Every 10 days, 5 occurrences:

       DTSTART;TZID=America/New_York:19970902T090000
       RRULE:FREQ=DAILY;INTERVAL=10;COUNT=5

       ==> (1997 9:00 AM EDT) September 2,12,22;
                              October 2,12

      Every day in January, for 3 years:

       DTSTART;TZID=America/New_York:19980101T090000

       RRULE:FREQ=YEARLY;UNTIL=20000131T140000Z;
        BYMONTH=1;BYDAY=SU,MO,TU,WE,TH,FR,SA
       or
       RRULE:FREQ=DAILY;UNTIL=20000131T140000Z;BYMONTH=1

       ==> (1998 9:00 AM EST)January 1-31
           (1999 9:00 AM EST)January 1-31
           (2000 9:00 AM EST)January 1-31

      Weekly for 10 occurrences:

       DTSTART;TZID=America/New_York:19970902T090000
       RRULE:FREQ=WEEKLY;COUNT=10

       ==> (1997 9:00 AM EDT) September 2,9,16,23,30;October 7,14,21
           (1997 9:00 AM EST) October 28;November 4










      Weekly until December 24, 1997:

       DTSTART;TZID=America/New_York:19970902T090000
       RRULE:FREQ=WEEKLY;UNTIL=19971224T000000Z

       ==> (1997 9:00 AM EDT) September 2,9,16,23,30;
                              October 7,14,21
           (1997 9:00 AM EST) October 28;
                              November 4,11,18,25;
                              December 2,9,16,23

      Every other week - forever:

       DTSTART;TZID=America/New_York:19970902T090000
       RRULE:FREQ=WEEKLY;INTERVAL=2;WKST=SU

       ==> (1997 9:00 AM EDT) September 2,16,30;
                              October 14
           (1997 9:00 AM EST) October 28;
                              November 11,25;
                              December 9,23
           (1998 9:00 AM EST) January 6,20;
                              February 3, 17
           ...

      Weekly on Tuesday and Thursday for five weeks:

       DTSTART;TZID=America/New_York:19970902T090000
       RRULE:FREQ=WEEKLY;UNTIL=19971007T000000Z;WKST=SU;BYDAY=TU,TH

       or

       RRULE:FREQ=WEEKLY;COUNT=10;WKST=SU;BYDAY=TU,TH

       ==> (1997 9:00 AM EDT) September 2,4,9,11,16,18,23,25,30;
                              October 2

      Every other week on Monday, Wednesday, and Friday until December
      24, 1997, starting on Monday, September 1, 1997:

       DTSTART;TZID=America/New_York:19970901T090000
       RRULE:FREQ=WEEKLY;INTERVAL=2;UNTIL=19971224T000000Z;WKST=SU;
        BYDAY=MO,WE,FR

       ==> (1997 9:00 AM EDT) September 1,3,5,15,17,19,29;
                              October 1,3,13,15,17
           (1997 9:00 AM EST) October 27,29,31;
                              November 10,12,14,24,26,28;

                              December 8,10,12,22

      Every other week on Tuesday and Thursday, for 8 occurrences:

       DTSTART;TZID=America/New_York:19970902T090000
       RRULE:FREQ=WEEKLY;INTERVAL=2;COUNT=8;WKST=SU;BYDAY=TU,TH

       ==> (1997 9:00 AM EDT) September 2,4,16,18,30;
                              October 2,14,16

      Monthly on the first Friday for 10 occurrences:

       DTSTART;TZID=America/New_York:19970905T090000
       RRULE:FREQ=MONTHLY;COUNT=10;BYDAY=1FR

       ==> (1997 9:00 AM EDT) September 5;October 3
           (1997 9:00 AM EST) November 7;December 5
           (1998 9:00 AM EST) January 2;February 6;March 6;April 3
           (1998 9:00 AM EDT) May 1;June 5

      Monthly on the first Friday until December 24, 1997:

       DTSTART;TZID=America/New_York:19970905T090000
       RRULE:FREQ=MONTHLY;UNTIL=19971224T000000Z;BYDAY=1FR

       ==> (1997 9:00 AM EDT) September 5; October 3
           (1997 9:00 AM EST) November 7; December 5

      Every other month on the first and last Sunday of the month for 10
      occurrences:

       DTSTART;TZID=America/New_York:19970907T090000
       RRULE:FREQ=MONTHLY;INTERVAL=2;COUNT=10;BYDAY=1SU,-1SU

       ==> (1997 9:00 AM EDT) September 7,28
           (1997 9:00 AM EST) November 2,30
           (1998 9:00 AM EST) January 4,25;March 1,29
           (1998 9:00 AM EDT) May 3,31

      Monthly on the second-to-last Monday of the month for 6 months:

       DTSTART;TZID=America/New_York:19970922T090000
       RRULE:FREQ=MONTHLY;COUNT=6;BYDAY=-2MO

       ==> (1997 9:00 AM EDT) September 22;October 20
           (1997 9:00 AM EST) November 17;December 22
           (1998 9:00 AM EST) January 19;February 16


      Monthly on the third-to-the-last day of the month, forever:

       DTSTART;TZID=America/New_York:19970928T090000
       RRULE:FREQ=MONTHLY;BYMONTHDAY=-3

       ==> (1997 9:00 AM EDT) September 28
           (1997 9:00 AM EST) October 29;November 28;December 29
           (1998 9:00 AM EST) January 29;February 26
           ...

      Monthly on the 2nd and 15th of the month for 10 occurrences:

       DTSTART;TZID=America/New_York:19970902T090000
       RRULE:FREQ=MONTHLY;COUNT=10;BYMONTHDAY=2,15

       ==> (1997 9:00 AM EDT) September 2,15;October 2,15
           (1997 9:00 AM EST) November 2,15;December 2,15
           (1998 9:00 AM EST) January 2,15

      Monthly on the first and last day of the month for 10 occurrences:

       DTSTART;TZID=America/New_York:19970930T090000
       RRULE:FREQ=MONTHLY;COUNT=10;BYMONTHDAY=1,-1

       ==> (1997 9:00 AM EDT) September 30;October 1
           (1997 9:00 AM EST) October 31;November 1,30;December 1,31
           (1998 9:00 AM EST) January 1,31;February 1

      Every 18 months on the 10th thru 15th of the month for 10
      occurrences:

       DTSTART;TZID=America/New_York:19970910T090000
       RRULE:FREQ=MONTHLY;INTERVAL=18;COUNT=10;BYMONTHDAY=10,11,12,
        13,14,15

       ==> (1997 9:00 AM EDT) September 10,11,12,13,14,15
           (1999 9:00 AM EST) March 10,11,12,13

      Every Tuesday, every other month:

       DTSTART;TZID=America/New_York:19970902T090000
       RRULE:FREQ=MONTHLY;INTERVAL=2;BYDAY=TU

       ==> (1997 9:00 AM EDT) September 2,9,16,23,30
           (1997 9:00 AM EST) November 4,11,18,25
           (1998 9:00 AM EST) January 6,13,20,27;March 3,10,17,24,31
           ...


      Yearly in June and July for 10 occurrences:

       DTSTART;TZID=America/New_York:19970610T090000
       RRULE:FREQ=YEARLY;COUNT=10;BYMONTH=6,7

       ==> (1997 9:00 AM EDT) June 10;July 10
           (1998 9:00 AM EDT) June 10;July 10
           (1999 9:00 AM EDT) June 10;July 10
           (2000 9:00 AM EDT) June 10;July 10
           (2001 9:00 AM EDT) June 10;July 10

         Note: Since none of the BYDAY, BYMONTHDAY, or BYYEARDAY
         components are specified, the day is gotten from "DTSTART".

      Every other year on January, February, and March for 10
      occurrences:

       DTSTART;TZID=America/New_York:19970310T090000
       RRULE:FREQ=YEARLY;INTERVAL=2;COUNT=10;BYMONTH=1,2,3

       ==> (1997 9:00 AM EST) March 10
           (1999 9:00 AM EST) January 10;February 10;March 10
           (2001 9:00 AM EST) January 10;February 10;March 10
           (2003 9:00 AM EST) January 10;February 10;March 10

      Every third year on the 1st, 100th, and 200th day for 10
      occurrences:

       DTSTART;TZID=America/New_York:19970101T090000
       RRULE:FREQ=YEARLY;INTERVAL=3;COUNT=10;BYYEARDAY=1,100,200

       ==> (1997 9:00 AM EST) January 1
           (1997 9:00 AM EDT) April 10;July 19
           (2000 9:00 AM EST) January 1
           (2000 9:00 AM EDT) April 9;July 18
           (2003 9:00 AM EST) January 1
           (2003 9:00 AM EDT) April 10;July 19
           (2006 9:00 AM EST) January 1

      Every 20th Monday of the year, forever:

       DTSTART;TZID=America/New_York:19970519T090000
       RRULE:FREQ=YEARLY;BYDAY=20MO

       ==> (1997 9:00 AM EDT) May 19
           (1998 9:00 AM EDT) May 18
           (1999 9:00 AM EDT) May 17
           ...

      Monday of week number 20 (where the default start of the week is
      Monday), forever:

       DTSTART;TZID=America/New_York:19970512T090000
       RRULE:FREQ=YEARLY;BYWEEKNO=20;BYDAY=MO

       ==> (1997 9:00 AM EDT) May 12
           (1998 9:00 AM EDT) May 11
           (1999 9:00 AM EDT) May 17
           ...

      Every Thursday in March, forever:

       DTSTART;TZID=America/New_York:19970313T090000
       RRULE:FREQ=YEARLY;BYMONTH=3;BYDAY=TH

       ==> (1997 9:00 AM EST) March 13,20,27
           (1998 9:00 AM EST) March 5,12,19,26
           (1999 9:00 AM EST) March 4,11,18,25
           ...

      Every Thursday, but only during June, July, and August, forever:

       DTSTART;TZID=America/New_York:19970605T090000
       RRULE:FREQ=YEARLY;BYDAY=TH;BYMONTH=6,7,8

       ==> (1997 9:00 AM EDT) June 5,12,19,26;July 3,10,17,24,31;
                              August 7,14,21,28
           (1998 9:00 AM EDT) June 4,11,18,25;July 2,9,16,23,30;
                              August 6,13,20,27
           (1999 9:00 AM EDT) June 3,10,17,24;July 1,8,15,22,29;
                              August 5,12,19,26
           ...

      Every Friday the 13th, forever:

       DTSTART;TZID=America/New_York:19970902T090000
       EXDATE;TZID=America/New_York:19970902T090000
       RRULE:FREQ=MONTHLY;BYDAY=FR;BYMONTHDAY=13

       ==> (1998 9:00 AM EST) February 13;March 13;November 13
           (1999 9:00 AM EDT) August 13
           (2000 9:00 AM EDT) October 13
           ...





      The first Saturday that follows the first Sunday of the month,
      forever:

       DTSTART;TZID=America/New_York:19970913T090000
       RRULE:FREQ=MONTHLY;BYDAY=SA;BYMONTHDAY=7,8,9,10,11,12,13

       ==> (1997 9:00 AM EDT) September 13;October 11
           (1997 9:00 AM EST) November 8;December 13
           (1998 9:00 AM EST) January 10;February 7;March 7
           (1998 9:00 AM EDT) April 11;May 9;June 13...
           ...

      Every 4 years, the first Tuesday after a Monday in November,
      forever (U.S. Presidential Election day):

       DTSTART;TZID=America/New_York:19961105T090000
       RRULE:FREQ=YEARLY;INTERVAL=4;BYMONTH=11;BYDAY=TU;
        BYMONTHDAY=2,3,4,5,6,7,8

        ==> (1996 9:00 AM EST) November 5
            (2000 9:00 AM EST) November 7
            (2004 9:00 AM EST) November 2
            ...

      The third instance into the month of one of Tuesday, Wednesday, or
      Thursday, for the next 3 months:

       DTSTART;TZID=America/New_York:19970904T090000
       RRULE:FREQ=MONTHLY;COUNT=3;BYDAY=TU,WE,TH;BYSETPOS=3

       ==> (1997 9:00 AM EDT) September 4;October 7
           (1997 9:00 AM EST) November 6

      The second-to-last weekday of the month:

       DTSTART;TZID=America/New_York:19970929T090000
       RRULE:FREQ=MONTHLY;BYDAY=MO,TU,WE,TH,FR;BYSETPOS=-2

       ==> (1997 9:00 AM EDT) September 29
           (1997 9:00 AM EST) October 30;November 27;December 30
           (1998 9:00 AM EST) January 29;February 26;March 30
           ...







      Every 3 hours from 9:00 AM to 5:00 PM on a specific day:

       DTSTART;TZID=America/New_York:19970902T090000
       RRULE:FREQ=HOURLY;INTERVAL=3;UNTIL=19970902T170000Z

       ==> (September 2, 1997 EDT) 09:00,12:00,15:00

      Every 15 minutes for 6 occurrences:

       DTSTART;TZID=America/New_York:19970902T090000
       RRULE:FREQ=MINUTELY;INTERVAL=15;COUNT=6

       ==> (September 2, 1997 EDT) 09:00,09:15,09:30,09:45,10:00,10:15

      Every hour and a half for 4 occurrences:

       DTSTART;TZID=America/New_York:19970902T090000
       RRULE:FREQ=MINUTELY;INTERVAL=90;COUNT=4

       ==> (September 2, 1997 EDT) 09:00,10:30;12:00;13:30

      Every 20 minutes from 9:00 AM to 4:40 PM every day:

       DTSTART;TZID=America/New_York:19970902T090000
       RRULE:FREQ=DAILY;BYHOUR=9,10,11,12,13,14,15,16;BYMINUTE=0,20,40
       or
       RRULE:FREQ=MINUTELY;INTERVAL=20;BYHOUR=9,10,11,12,13,14,15,16

       ==> (September 2, 1997 EDT) 9:00,9:20,9:40,10:00,10:20,
                                   ... 16:00,16:20,16:40
           (September 3, 1997 EDT) 9:00,9:20,9:40,10:00,10:20,
                                   ...16:00,16:20,16:40
           ...

      An example where the days generated makes a difference because of
      WKST:

       DTSTART;TZID=America/New_York:19970805T090000
       RRULE:FREQ=WEEKLY;INTERVAL=2;COUNT=4;BYDAY=TU,SU;WKST=MO

       ==> (1997 EDT) August 5,10,19,24

      changing only WKST from MO to SU, yields different results...

       DTSTART;TZID=America/New_York:19970805T090000
       RRULE:FREQ=WEEKLY;INTERVAL=2;COUNT=4;BYDAY=TU,SU;WKST=SU

       ==> (1997 EDT) August 5,17,19,31

      An example where an invalid date (i.e., February 30) is ignored.

       DTSTART;TZID=America/New_York:20070115T090000
       RRULE:FREQ=MONTHLY;BYMONTHDAY=15,30;COUNT=5

       ==> (2007 EST) January 15,30
           (2007 EST) February 15
           (2007 EDT) March 15,30

3.8.6.  Alarm Component Properties

   The following properties specify alarm information in calendar
   components.

3.8.6.1.  Action

   Property Name:  ACTION

   Purpose:  This property defines the action to be invoked when an
      alarm is triggered.

   Value Type:  TEXT

   Property Parameters:  IANA and non-standard property parameters can
      be specified on this property.

   Conformance:  This property MUST be specified once in a "VALARM"
      calendar component.

   Description:  Each "VALARM" calendar component has a particular type
      of action with which it is associated.  This property specifies
      the type of action.  Applications MUST ignore alarms with x-name
      and iana-token values they don't recognize.

   Format Definition:  This property is defined by the following
      notation:

       action      = "ACTION" actionparam ":" actionvalue CRLF

       actionparam = *(";" other-param)


       actionvalue = "AUDIO" / "DISPLAY" / "EMAIL"
                   / iana-token / x-name

   Example:  The following are examples of this property in a "VALARM"
      calendar component:


       ACTION:AUDIO

       ACTION:DISPLAY

3.8.6.2.  Repeat Count

   Property Name:  REPEAT

   Purpose:  This property defines the number of times the alarm should
      be repeated, after the initial trigger.

   Value Type:  INTEGER

   Property Parameters:  IANA and non-standard property parameters can
      be specified on this property.

   Conformance:  This property can be specified in a "VALARM" calendar
      component.

   Description:  This property defines the number of times an alarm
      should be repeated after its initial trigger.  If the alarm
      triggers more than once, then this property MUST be specified
      along with the "DURATION" property.

   Format Definition:  This property is defined by the following
      notation:

       repeat  = "REPEAT" repparam ":" integer CRLF
       ;Default is "0", zero.

       repparam   = *(";" other-param)

   Example:  The following is an example of this property for an alarm
      that repeats 4 additional times with a 5-minute delay after the
      initial triggering of the alarm:

       REPEAT:4
       DURATION:PT5M

3.8.6.3.  Trigger

   Property Name:  TRIGGER

   Purpose:  This property specifies when an alarm will trigger.

   Value Type:  The default value type is DURATION.  The value type can
      be set to a DATE-TIME value type, in which case the value MUST
      specify a UTC-formatted DATE-TIME value.

   Property Parameters:  IANA, non-standard, value data type, time zone
      identifier, or trigger relationship property parameters can be
      specified on this property.  The trigger relationship property
      parameter MUST only be specified when the value type is
      "DURATION".

   Conformance:  This property MUST be specified in the "VALARM"
      calendar component.

   Description:  This property defines when an alarm will trigger.  The
      default value type is DURATION, specifying a relative time for the
      trigger of the alarm.  The default duration is relative to the
      start of an event or to-do with which the alarm is associated.
      The duration can be explicitly set to trigger from either the end
      or the start of the associated event or to-do with the "RELATED"
      parameter.  A value of START will set the alarm to trigger off the
      start of the associated event or to-do.  A value of END will set
      the alarm to trigger off the end of the associated event or to-do.

      Either a positive or negative duration may be specified for the
      "TRIGGER" property.  An alarm with a positive duration is
      triggered after the associated start or end of the event or to-do.
      An alarm with a negative duration is triggered before the
      associated start or end of the event or to-do.

      The "RELATED" property parameter is not valid if the value type of
      the property is set to DATE-TIME (i.e., for an absolute date and
      time alarm trigger).  If a value type of DATE-TIME is specified,
      then the property value MUST be specified in the UTC time format.
      If an absolute trigger is specified on an alarm for a recurring
      event or to-do, then the alarm will only trigger for the specified
      absolute DATE-TIME, along with any specified repeating instances.

      If the trigger is set relative to START, then the "DTSTART"
      property MUST be present in the associated "VEVENT" or "VTODO"
      calendar component.  If an alarm is specified for an event with
      the trigger set relative to the END, then the "DTEND" property or
      the "DTSTART" and "DURATION " properties MUST be present in the
      associated "VEVENT" calendar component.  If the alarm is specified
      for a to-do with a trigger set relative to the END, then either
      the "DUE" property or the "DTSTART" and "DURATION " properties
      MUST be present in the associated "VTODO" calendar component.

      Alarms specified in an event or to-do that is defined in terms of
      a DATE value type will be triggered relative to 00:00:00 of the
      user's configured time zone on the specified date, or relative to
      00:00:00 UTC on the specified date if no configured time zone can
      be found for the user.  For example, if "DTSTART" is a DATE value

      set to 19980205 then the duration trigger will be relative to
      19980205T000000 America/New_York for a user configured with the
      America/New_York time zone.

   Format Definition:  This property is defined by the following
      notation:

       trigger    = "TRIGGER" (trigrel / trigabs) CRLF

       trigrel    = *(
                  ;
                  ; The following are OPTIONAL,
                  ; but MUST NOT occur more than once.
                  ;
                  (";" "VALUE" "=" "DURATION") /
                  (";" trigrelparam) /
                  ;
                  ; The following is OPTIONAL,
                  ; and MAY occur more than once.
                  ;
                  (";" other-param)
                  ;
                  ) ":"  dur-value

       trigabs    = *(
                  ;
                  ; The following is REQUIRED,
                  ; but MUST NOT occur more than once.
                  ;
                  (";" "VALUE" "=" "DATE-TIME") /
                  ;
                  ; The following is OPTIONAL,
                  ; and MAY occur more than once.
                  ;
                  (";" other-param)
                  ;
                  ) ":" date-time

   Example:  A trigger set 15 minutes prior to the start of the event or
      to-do.

       TRIGGER:-PT15M

      A trigger set five minutes after the end of an event or the due
      date of a to-do.

       TRIGGER;RELATED=END:PT5M


      A trigger set to an absolute DATE-TIME.

       TRIGGER;VALUE=DATE-TIME:19980101T050000Z

3.8.7.  Change Management Component Properties

   The following properties specify change management information in
   calendar components.

3.8.7.1.  Date-Time Created

   Property Name:  CREATED

   Purpose:  This property specifies the date and time that the calendar
      information was created by the calendar user agent in the calendar
      store.

         Note: This is analogous to the creation date and time for a
         file in the file system.

   Value Type:  DATE-TIME

   Property Parameters:  IANA and non-standard property parameters can
      be specified on this property.

   Conformance:  The property can be specified once in "VEVENT",
      "VTODO", or "VJOURNAL" calendar components.  The value MUST be
      specified as a date with UTC time.

   Description:  This property specifies the date and time that the
      calendar information was created by the calendar user agent in the
      calendar store.

   Format Definition:  This property is defined by the following
      notation:

       created    = "CREATED" creaparam ":" date-time CRLF

       creaparam  = *(";" other-param)

   Example:  The following is an example of this property:

       CREATED:19960329T133000Z






3.8.7.2.  Date-Time Stamp

   Property Name:  DTSTAMP

   Purpose:  In the case of an iCalendar object that specifies a
      "METHOD" property, this property specifies the date and time that
      the instance of the iCalendar object was created.  In the case of
      an iCalendar object that doesn't specify a "METHOD" property, this
      property specifies the date and time that the information
      associated with the calendar component was last revised in the
      calendar store.

   Value Type:  DATE-TIME

   Property Parameters:  IANA and non-standard property parameters can
      be specified on this property.

   Conformance:  This property MUST be included in the "VEVENT",
      "VTODO", "VJOURNAL", or "VFREEBUSY" calendar components.

   Description:  The value MUST be specified in the UTC time format.

      This property is also useful to protocols such as [2447bis] that
      have inherent latency issues with the delivery of content.  This
      property will assist in the proper sequencing of messages
      containing iCalendar objects.

      In the case of an iCalendar object that specifies a "METHOD"
      property, this property differs from the "CREATED" and "LAST-
      MODIFIED" properties.  These two properties are used to specify
      when the particular calendar data in the calendar store was
      created and last modified.  This is different than when the
      iCalendar object representation of the calendar service
      information was created or last modified.

      In the case of an iCalendar object that doesn't specify a "METHOD"
      property, this property is equivalent to the "LAST-MODIFIED"
      property.

   Format Definition:  This property is defined by the following
      notation:

       dtstamp    = "DTSTAMP" stmparam ":" date-time CRLF

       stmparam   = *(";" other-param)




   Example:

       DTSTAMP:19971210T080000Z

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

3.8.7.4.  Sequence Number

   Property Name:  SEQUENCE

   Purpose:  This property defines the revision sequence number of the
      calendar component within a sequence of revisions.

   Value Type:  INTEGER

   Property Parameters:  IANA and non-standard property parameters can
      be specified on this property.

   Conformance:  The property can be specified in "VEVENT", "VTODO", or
      "VJOURNAL" calendar component.

   Description:  When a calendar component is created, its sequence
      number is 0.  It is monotonically incremented by the "Organizer's"
      CUA each time the "Organizer" makes a significant revision to the
      calendar component.

      The "Organizer" includes this property in an iCalendar object that
      it sends to an "Attendee" to specify the current version of the
      calendar component.

      The "Attendee" includes this property in an iCalendar object that
      it sends to the "Organizer" to specify the version of the calendar
      component to which the "Attendee" is referring.

      A change to the sequence number is not the mechanism that an
      "Organizer" uses to request a response from the "Attendees".  The
      "RSVP" parameter on the "ATTENDEE" property is used by the
      "Organizer" to indicate that a response from the "Attendees" is
      requested.

      Recurrence instances of a recurring component MAY have different
      sequence numbers.

   Format Definition:  This property is defined by the following
      notation:

       seq = "SEQUENCE" seqparam ":" integer CRLF
       ; Default is "0"

       seqparam   = *(";" other-param)

   Example:  The following is an example of this property for a calendar
      component that was just created by the "Organizer":

       SEQUENCE:0

      The following is an example of this property for a calendar
      component that has been revised two different times by the
      "Organizer":

       SEQUENCE:2

3.8.8.  Miscellaneous Component Properties

   The following properties specify information about a number of
   miscellaneous features of calendar components.

3.8.8.1.  IANA Properties

   Property Name:  An IANA-registered property name

   Value Type:  The default value type is TEXT.  The value type can be
      set to any value type.

   Property Parameters:  Any parameter can be specified on this
      property.

   Description:  This specification allows other properties registered
      with IANA to be specified in any calendar components.  Compliant
      applications are expected to be able to parse these other IANA-
      registered properties but can ignore them.

   Format Definition:  This property is defined by the following
      notation:

       iana-prop = iana-token *(";" icalparameter) ":" value CRLF

   Example:  The following are examples of properties that might be
      registered to IANA:

       DRESSCODE:CASUAL

       NON-SMOKING;VALUE=BOOLEAN:TRUE

3.8.8.2.  Non-Standard Properties

   Property Name:  Any property name with a "X-" prefix

   Purpose:  This class of property provides a framework for defining
      non-standard properties.

   Value Type:  The default value type is TEXT.  The value type can be
      set to any value type.

   Property Parameters:  IANA, non-standard, and language property
      parameters can be specified on this property.

   Conformance:  This property can be specified in any calendar
      component.

   Description:  The MIME Calendaring and Scheduling Content Type
      provides a "standard mechanism for doing non-standard things".
      This extension support is provided for implementers to "push the
      envelope" on the existing version of the memo.  Extension
      properties are specified by property and/or property parameter

      names that have the prefix text of "X-" (the two-character
      sequence: LATIN CAPITAL LETTER X character followed by the HYPHEN-
      MINUS character).  It is recommended that vendors concatenate onto
      this sentinel another short prefix text to identify the vendor.
      This will facilitate readability of the extensions and minimize
      possible collision of names between different vendors.  User
      agents that support this content type are expected to be able to
      parse the extension properties and property parameters but can
      ignore them.

      At present, there is no registration authority for names of
      extension properties and property parameters.  The value type for
      this property is TEXT.  Optionally, the value type can be any of
      the other valid value types.

   Format Definition:  This property is defined by the following
      notation:

       x-prop = x-name *(";" icalparameter) ":" value CRLF

   Example:  The following might be the ABC vendor's extension for an
      audio-clip form of subject property:

       X-ABC-MMSUBJ;VALUE=URI;FMTTYPE=audio/basic:http://www.example.
        org/mysubj.au

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
