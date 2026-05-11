# Calendar Properties

Source: RFC 5545, Section 3.7

3.7.  Calendar Properties

   The Calendar Properties are attributes that apply to the iCalendar
   object, as a whole.  These properties do not appear within a calendar
   component.  They SHOULD be specified after the "BEGIN:VCALENDAR"
   delimiter string and prior to any calendar component.

3.7.1.  Calendar Scale

   Property Name:  CALSCALE

   Purpose:  This property defines the calendar scale used for the
      calendar information specified in the iCalendar object.

   Value Type:  TEXT

   Property Parameters:  IANA and non-standard property parameters can
      be specified on this property.

   Conformance:  This property can be specified once in an iCalendar
      object.  The default value is "GREGORIAN".

   Description:  This memo is based on the Gregorian calendar scale.
      The Gregorian calendar scale is assumed if this property is not
      specified in the iCalendar object.  It is expected that other
      calendar scales will be defined in other specifications or by
      future versions of this memo.

   Format Definition:  This property is defined by the following
      notation:

       calscale   = "CALSCALE" calparam ":" calvalue CRLF

       calparam   = *(";" other-param)

       calvalue   = "GREGORIAN"

   Example:  The following is an example of this property:

       CALSCALE:GREGORIAN

3.7.2.  Method

   Property Name:  METHOD

   Purpose:  This property defines the iCalendar object method
      associated with the calendar object.

   Value Type:  TEXT

   Property Parameters:  IANA and non-standard property parameters can
      be specified on this property.

   Conformance:  This property can be specified once in an iCalendar
      object.

   Description:  When used in a MIME message entity, the value of this
      property MUST be the same as the Content-Type "method" parameter
      value.  If either the "METHOD" property or the Content-Type
      "method" parameter is specified, then the other MUST also be
      specified.

      No methods are defined by this specification.  This is the subject
      of other specifications, such as the iCalendar Transport-
      independent Interoperability Protocol (iTIP) defined by [2446bis].

      If this property is not present in the iCalendar object, then a
      scheduling transaction MUST NOT be assumed.  In such cases, the
      iCalendar object is merely being used to transport a snapshot of


      some calendar information; without the intention of conveying a
      scheduling semantic.

   Format Definition:  This property is defined by the following
      notation:

       method     = "METHOD" metparam ":" metvalue CRLF

       metparam   = *(";" other-param)

       metvalue   = iana-token

   Example:  The following is a hypothetical example of this property to
      convey that the iCalendar object is a scheduling request:

       METHOD:REQUEST

3.7.3.  Product Identifier

   Property Name:  PRODID

   Purpose:  This property specifies the identifier for the product that
      created the iCalendar object.

   Value Type:  TEXT

   Property Parameters:  IANA and non-standard property parameters can
      be specified on this property.

   Conformance:  The property MUST be specified once in an iCalendar
      object.

   Description:  The vendor of the implementation SHOULD assure that
      this is a globally unique identifier; using some technique such as
      an FPI value, as defined in [ISO.9070.1991].

      This property SHOULD NOT be used to alter the interpretation of an
      iCalendar object beyond the semantics specified in this memo.  For
      example, it is not to be used to further the understanding of non-
      standard properties.

   Format Definition:  This property is defined by the following
      notation:

       prodid     = "PRODID" pidparam ":" pidvalue CRLF

       pidparam   = *(";" other-param)


       pidvalue   = text
       ;Any text that describes the product and version
       ;and that is generally assured of being unique.

   Example:  The following is an example of this property.  It does not
      imply that English is the default language.

       PRODID:-//ABC Corporation//NONSGML My Product//EN

3.7.4.  Version

   Property Name:  VERSION

   Purpose:  This property specifies the identifier corresponding to the
      highest version number or the minimum and maximum range of the
      iCalendar specification that is required in order to interpret the
      iCalendar object.

   Value Type:  TEXT

   Property Parameters:  IANA and non-standard property parameters can
      be specified on this property.

   Conformance:  This property MUST be specified once in an iCalendar
      object.

   Description:  A value of "2.0" corresponds to this memo.

   Format Definition:  This property is defined by the following
      notation:

       version    = "VERSION" verparam ":" vervalue CRLF

       verparam   = *(";" other-param)

       vervalue   = "2.0"         ;This memo
                  / maxver
                  / (minver ";" maxver)

       minver     = <A IANA-registered iCalendar version identifier>
       ;Minimum iCalendar version needed to parse the iCalendar object.

       maxver     = <A IANA-registered iCalendar version identifier>
       ;Maximum iCalendar version needed to parse the iCalendar object.





   Example:  The following is an example of this property:

       VERSION:2.0
