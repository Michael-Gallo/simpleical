# Method

Source: RFC 5545, Section 3.7.2

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
