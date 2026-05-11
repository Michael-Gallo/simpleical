# Integer

Source: RFC 5545, Section 3.3.8

3.3.8.  Integer

   Value Name:  INTEGER

   Purpose:  This value type is used to identify properties that contain
      a signed integer value.

   Format Definition:  This value type is defined by the following
      notation:

       integer    = (["+"] / "-") 1*DIGIT

   Description:  If the property permits, multiple "integer" values are
      specified by a COMMA-separated list of values.  The valid range
      for "integer" is -2147483648 to 2147483647.  If the sign is not
      specified, then the value is assumed to be positive.

      No additional content value encoding (i.e., BACKSLASH character
      encoding, see Section 3.3.11) is defined for this value type.

   Example:

       1234567890
       -1234567890
       +1234567890
       432109876
