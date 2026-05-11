# URI

Source: RFC 5545, Section 3.3.13

3.3.13.  URI

   Value Name:  URI

   Purpose:  This value type is used to identify values that contain a
      uniform resource identifier (URI) type of reference to the
      property value.

   Format Definition:  This value type is defined by the following
      notation:

       uri = <As defined in Section 3 of [RFC3986]>

   Description:  This value type might be used to reference binary
      information, for values that are large, or otherwise undesirable
      to include directly in the iCalendar object.

      Property values with this value type MUST follow the generic URI
      syntax defined in [RFC3986].

      When a property parameter value is a URI value type, the URI MUST
      be specified as a quoted-string value.

      No additional content value encoding (i.e., BACKSLASH character
      encoding, see Section 3.3.11) is defined for this value type.

   Example:  The following is a URI for a network file:

       http://example.com/my-report.txt
