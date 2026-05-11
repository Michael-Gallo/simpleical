# Binary

Source: RFC 5545, Section 3.3.1

3.3.1.  Binary

   Value Name:  BINARY

   Purpose:  This value type is used to identify properties that contain
      a character encoding of inline binary data.  For example, an
      inline attachment of a document might be included in an iCalendar
      object.

   Format Definition:  This value type is defined by the following
      notation:

       binary     = *(4b-char) [b-end]
       ; A "BASE64" encoded character string, as defined by [RFC4648].

       b-end      = (2b-char "==") / (3b-char "=")

       b-char = ALPHA / DIGIT / "+" / "/"

   Description:  Property values with this value type MUST also include
      the inline encoding parameter sequence of ";ENCODING=BASE64".
      That is, all inline binary data MUST first be character encoded
      using the "BASE64" encoding method defined in [RFC2045].  No
      additional content value encoding (i.e., BACKSLASH character
      encoding, see Section 3.3.11) is defined for this value type.













   Example:  The following is an example of a "BASE64" encoded binary
      value data:

     ATTACH;FMTTYPE=image/vnd.microsoft.icon;ENCODING=BASE64;VALUE
      =BINARY:AAABAAEAEBAQAAEABAAoAQAAFgAAACgAAAAQAAAAIAAAAAEABAAA
      AAAAAAAAAAAAAAAAAAAAAAAAAAAAAACAAAAAgIAAAICAgADAwMAA////AAAA
      AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
      AAAAAAAAAAAAAAAAAAAAAAMwAAAAAAABNEMQAAAAAAAkQgAAAAAAJEREQgAA
      ACECQ0QgEgAAQxQzM0E0AABERCRCREQAADRDJEJEQwAAAhA0QwEQAAAAAERE
      AAAAAAAAREQAAAAAAAAkQgAAAAAAAAMgAAAAAAAAAAAAAAAAAAAAAAAAAAAA
      AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
      AAAAAAAAAAAA
