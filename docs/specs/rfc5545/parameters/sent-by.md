# Sent By

Source: RFC 5545, Section 3.2.18

3.2.18.  Sent By

   Parameter Name:  SENT-BY

   Purpose:  To specify the calendar user that is acting on behalf of
      the calendar user specified by the property.

   Format Definition:  This property parameter is defined by the
      following notation:

       sentbyparam        = "SENT-BY" "=" DQUOTE cal-address DQUOTE

   Description:  This parameter can be specified on properties with a
      CAL-ADDRESS value type.  The parameter specifies the calendar user
      that is acting on behalf of the calendar user specified by the
      property.  The parameter value MUST be a mailto URI as defined in
      [RFC2368].  The individual calendar address parameter values MUST
      each be specified in a quoted-string.

   Example:

       ORGANIZER;SENT-BY="mailto:sray@example.com":mailto:
        jsmith@example.com
