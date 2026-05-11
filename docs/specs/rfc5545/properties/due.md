# Date-Time Due

Source: RFC 5545, Section 3.8.2.3

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
