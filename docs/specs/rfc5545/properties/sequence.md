# Sequence Number

Source: RFC 5545, Section 3.8.7.4

3.8.7.4.  Sequence Number

   Property Name:  SEQUENCE

   Purpose:  This property defines the revision sequence number of the
      calendar component within a sequence of revisions.

   Value Type:  INTEGER

   Property Parameters:  IANA and non-standard property parameters can
      be specified on this property.

   Conformance:  The property can be specified in "VEVENT", "VTODO", or
      "VJOURNAL" calendar component.

   Description:  When a calendar component is created, its sequence
      number is 0.  It is monotonically incremented by the "Organizer's"
      CUA each time the "Organizer" makes a significant revision to the
      calendar component.

      The "Organizer" includes this property in an iCalendar object that
      it sends to an "Attendee" to specify the current version of the
      calendar component.

      The "Attendee" includes this property in an iCalendar object that
      it sends to the "Organizer" to specify the version of the calendar
      component to which the "Attendee" is referring.

      A change to the sequence number is not the mechanism that an
      "Organizer" uses to request a response from the "Attendees".  The
      "RSVP" parameter on the "ATTENDEE" property is used by the
      "Organizer" to indicate that a response from the "Attendees" is
      requested.

      Recurrence instances of a recurring component MAY have different
      sequence numbers.

   Format Definition:  This property is defined by the following
      notation:

       seq = "SEQUENCE" seqparam ":" integer CRLF
       ; Default is "0"

       seqparam   = *(";" other-param)

   Example:  The following is an example of this property for a calendar
      component that was just created by the "Organizer":

       SEQUENCE:0

      The following is an example of this property for a calendar
      component that has been revised two different times by the
      "Organizer":

       SEQUENCE:2
