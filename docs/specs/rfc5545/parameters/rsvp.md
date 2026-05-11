# RSVP Expectation

Source: RFC 5545, Section 3.2.17

3.2.17.  RSVP Expectation

   Parameter Name:  RSVP

   Purpose:  To specify whether there is an expectation of a favor of a
      reply from the calendar user specified by the property value.

   Format Definition:  This property parameter is defined by the
      following notation:

       rsvpparam = "RSVP" "=" ("TRUE" / "FALSE")
       ; Default is FALSE

   Description:  This parameter can be specified on properties with a
      CAL-ADDRESS value type.  The parameter identifies the expectation
      of a reply from the calendar user specified by the property value.
      This parameter is used by the "Organizer" to request a
      participation status reply from an "Attendee" of a group-scheduled
      event or to-do.  If not specified on a property that allows this
      parameter, the default value is FALSE.



   Example:

       ATTENDEE;RSVP=TRUE:mailto:jsmith@example.com
