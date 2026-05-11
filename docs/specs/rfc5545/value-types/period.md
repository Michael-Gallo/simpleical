# Period of Time

Source: RFC 5545, Section 3.3.9

3.3.9.  Period of Time

   Value Name:  PERIOD

   Purpose:  This value type is used to identify values that contain a
      precise period of time.

   Format Definition:  This value type is defined by the following
      notation:

       period     = period-explicit / period-start

       period-explicit = date-time "/" date-time
       ; [ISO.8601.2004] complete representation basic format for a
       ; period of time consisting of a start and end.  The start MUST
       ; be before the end.

       period-start = date-time "/" dur-value
       ; [ISO.8601.2004] complete representation basic format for a
       ; period of time consisting of a start and positive duration
       ; of time.

   Description:  If the property permits, multiple "period" values are
      specified by a COMMA-separated list of values.  There are two
      forms of a period of time.  First, a period of time is identified
      by its start and its end.  This format is based on the
      [ISO.8601.2004] complete representation, basic format for "DATE-
      TIME" start of the period, followed by a SOLIDUS character
      followed by the "DATE-TIME" of the end of the period.  The start
      of the period MUST be before the end of the period.  Second, a
      period of time can also be defined by a start and a positive
      duration of time.  The format is based on the [ISO.8601.2004]
      complete representation, basic format for the "DATE-TIME" start of
      the period, followed by a SOLIDUS character, followed by the
      [ISO.8601.2004] basic format for "DURATION" of the period.

   Example:  The period starting at 18:00:00 UTC, on January 1, 1997 and
      ending at 07:00:00 UTC on January 2, 1997 would be:

       19970101T180000Z/19970102T070000Z

      The period start at 18:00:00 on January 1, 1997 and lasting 5
      hours and 30 minutes would be:

       19970101T180000Z/PT5H30M

      No additional content value encoding (i.e., BACKSLASH character
      encoding, see Section 3.3.11) is defined for this value type.
