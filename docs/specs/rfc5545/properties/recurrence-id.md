# Recurrence ID

Source: RFC 5545, Section 3.8.4.4

3.8.4.4.  Recurrence ID

   Property Name:  RECURRENCE-ID

   Purpose:  This property is used in conjunction with the "UID" and
      "SEQUENCE" properties to identify a specific instance of a
      recurring "VEVENT", "VTODO", or "VJOURNAL" calendar component.
      The property value is the original value of the "DTSTART" property
      of the recurrence instance.

   Value Type:  The default value type is DATE-TIME.  The value type can
      be set to a DATE value type.  This property MUST have the same
      value type as the "DTSTART" property contained within the
      recurring component.  Furthermore, this property MUST be specified
      as a date with local time if and only if the "DTSTART" property
      contained within the recurring component is specified as a date
      with local time.

   Property Parameters:  IANA, non-standard, value data type, time zone
      identifier, and recurrence identifier range parameters can be
      specified on this property.

   Conformance:  This property can be specified in an iCalendar object
      containing a recurring calendar component.

   Description:  The full range of calendar components specified by a
      recurrence set is referenced by referring to just the "UID"
      property value corresponding to the calendar component.  The
      "RECURRENCE-ID" property allows the reference to an individual
      instance within the recurrence set.

      If the value of the "DTSTART" property is a DATE type value, then
      the value MUST be the calendar date for the recurrence instance.

      The DATE-TIME value is set to the time when the original
      recurrence instance would occur; meaning that if the intent is to
      change a Friday meeting to Thursday, the DATE-TIME is still set to
      the original Friday meeting.

      The "RECURRENCE-ID" property is used in conjunction with the "UID"
      and "SEQUENCE" properties to identify a particular instance of a
      recurring event, to-do, or journal.  For a given pair of "UID" and
      "SEQUENCE" property values, the "RECURRENCE-ID" value for a
      recurrence instance is fixed.

      The "RANGE" parameter is used to specify the effective range of
      recurrence instances from the instance specified by the
      "RECURRENCE-ID" property value.  The value for the range parameter
      can only be "THISANDFUTURE" to indicate a range defined by the
      given recurrence instance and all subsequent instances.
      Subsequent instances are determined by their "RECURRENCE-ID" value
      and not their current scheduled start time.  Subsequent instances
      defined in separate components are not impacted by the given
      recurrence instance.  When the given recurrence instance is
      rescheduled, all subsequent instances are also rescheduled by the
      same time difference.  For instance, if the given recurrence
      instance is rescheduled to start 2 hours later, then all
      subsequent instances are also rescheduled 2 hours later.

      Similarly, if the duration of the given recurrence instance is
      modified, then all subsequence instances are also modified to have
      this same duration.

         Note: The "RANGE" parameter may not be appropriate to
         reschedule specific subsequent instances of complex recurring
         calendar component.  Assuming an unbounded recurring calendar
         component scheduled to occur on Mondays and Wednesdays, the
         "RANGE" parameter could not be used to reschedule only the
         future Monday instances to occur on Tuesday instead.  In such
         cases, the calendar application could simply truncate the
         unbounded recurring calendar component (i.e., with the "COUNT"
         or "UNTIL" rule parts), and create two new unbounded recurring
         calendar components for the future instances.

   Format Definition:  This property is defined by the following
      notation:

       recurid    = "RECURRENCE-ID" ridparam ":" ridval CRLF

       ridparam   = *(
                  ;
                  ; The following are OPTIONAL,
                  ; but MUST NOT occur more than once.
                  ;
                  (";" "VALUE" "=" ("DATE-TIME" / "DATE")) /
                  (";" tzidparam) / (";" rangeparam) /
                  ;
                  ; The following is OPTIONAL,
                  ; and MAY occur more than once.
                  ;
                  (";" other-param)
                  ;
                  )

       ridval     = date-time / date
       ;Value MUST match value type

   Example:  The following are examples of this property:

       RECURRENCE-ID;VALUE=DATE:19960401

       RECURRENCE-ID;RANGE=THISANDFUTURE:19960120T120000Z
