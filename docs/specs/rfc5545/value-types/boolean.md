# Boolean

Source: RFC 5545, Section 3.3.2

3.3.2.  Boolean

   Value Name:  BOOLEAN

   Purpose:  This value type is used to identify properties that contain
      either a "TRUE" or "FALSE" Boolean value.

   Format Definition:  This value type is defined by the following
      notation:

       boolean    = "TRUE" / "FALSE"

   Description:  These values are case-insensitive text.  No additional
      content value encoding (i.e., BACKSLASH character encoding, see
      Section 3.3.11) is defined for this value type.

   Example:  The following is an example of a hypothetical property that
      has a BOOLEAN value type:

       TRUE
