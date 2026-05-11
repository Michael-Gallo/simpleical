# Date

Source: RFC 5545, Section 3.3.4

3.3.4.  Date

   Value Name:  DATE

   Purpose:  This value type is used to identify values that contain a
      calendar date.

   Format Definition:  This value type is defined by the following
      notation:

       date               = date-value

       date-value         = date-fullyear date-month date-mday
       date-fullyear      = 4DIGIT
       date-month         = 2DIGIT        ;01-12
       date-mday          = 2DIGIT        ;01-28, 01-29, 01-30, 01-31
                                          ;based on month/year

   Description:  If the property permits, multiple "date" values are
      specified as a COMMA-separated list of values.  The format for the
      value type is based on the [ISO.8601.2004] complete
      representation, basic format for a calendar date.  The textual
      format specifies a four-digit year, two-digit month, and two-digit
      day of the month.  There are no separator characters between the
      year, month, and day component text.

      No additional content value encoding (i.e., BACKSLASH character
      encoding, see Section 3.3.11) is defined for this value type.

   Example:  The following represents July 14, 1997:

       19970714
