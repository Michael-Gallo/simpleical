# Delegatees

Source: RFC 5545, Section 3.2.5

3.2.5.  Delegatees

   Parameter Name:  DELEGATED-TO

   Purpose:  To specify the calendar users to whom the calendar user
      specified by the property has delegated participation.

   Format Definition:  This property parameter is defined by the
      following notation:

       deltoparam = "DELEGATED-TO" "=" DQUOTE cal-address DQUOTE
                    *("," DQUOTE cal-address DQUOTE)

   Description:  This parameter can be specified on properties with a
      CAL-ADDRESS value type.  This parameter specifies those calendar
      users whom have been delegated participation in a group-scheduled
      event or to-do by the calendar user specified by the property.
      The individual calendar address parameter values MUST each be
      specified in a quoted-string.





   Example:

       ATTENDEE;DELEGATED-TO="mailto:jdoe@example.com","mailto:jqpublic
        @example.com":mailto:jsmith@example.com
