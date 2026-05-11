# Trigger

Source: RFC 5545, Section 3.8.6.3

3.8.6.3.  Trigger

   Property Name:  TRIGGER

   Purpose:  This property specifies when an alarm will trigger.

   Value Type:  The default value type is DURATION.  The value type can
      be set to a DATE-TIME value type, in which case the value MUST
      specify a UTC-formatted DATE-TIME value.

   Property Parameters:  IANA, non-standard, value data type, time zone
      identifier, or trigger relationship property parameters can be
      specified on this property.  The trigger relationship property
      parameter MUST only be specified when the value type is
      "DURATION".

   Conformance:  This property MUST be specified in the "VALARM"
      calendar component.

   Description:  This property defines when an alarm will trigger.  The
      default value type is DURATION, specifying a relative time for the
      trigger of the alarm.  The default duration is relative to the
      start of an event or to-do with which the alarm is associated.
      The duration can be explicitly set to trigger from either the end
      or the start of the associated event or to-do with the "RELATED"
      parameter.  A value of START will set the alarm to trigger off the
      start of the associated event or to-do.  A value of END will set
      the alarm to trigger off the end of the associated event or to-do.

      Either a positive or negative duration may be specified for the
      "TRIGGER" property.  An alarm with a positive duration is
      triggered after the associated start or end of the event or to-do.
      An alarm with a negative duration is triggered before the
      associated start or end of the event or to-do.

      The "RELATED" property parameter is not valid if the value type of
      the property is set to DATE-TIME (i.e., for an absolute date and
      time alarm trigger).  If a value type of DATE-TIME is specified,
      then the property value MUST be specified in the UTC time format.
      If an absolute trigger is specified on an alarm for a recurring
      event or to-do, then the alarm will only trigger for the specified
      absolute DATE-TIME, along with any specified repeating instances.

      If the trigger is set relative to START, then the "DTSTART"
      property MUST be present in the associated "VEVENT" or "VTODO"
      calendar component.  If an alarm is specified for an event with
      the trigger set relative to the END, then the "DTEND" property or
      the "DTSTART" and "DURATION " properties MUST be present in the
      associated "VEVENT" calendar component.  If the alarm is specified
      for a to-do with a trigger set relative to the END, then either
      the "DUE" property or the "DTSTART" and "DURATION " properties
      MUST be present in the associated "VTODO" calendar component.

      Alarms specified in an event or to-do that is defined in terms of
      a DATE value type will be triggered relative to 00:00:00 of the
      user's configured time zone on the specified date, or relative to
      00:00:00 UTC on the specified date if no configured time zone can
      be found for the user.  For example, if "DTSTART" is a DATE value

      set to 19980205 then the duration trigger will be relative to
      19980205T000000 America/New_York for a user configured with the
      America/New_York time zone.

   Format Definition:  This property is defined by the following
      notation:

       trigger    = "TRIGGER" (trigrel / trigabs) CRLF

       trigrel    = *(
                  ;
                  ; The following are OPTIONAL,
                  ; but MUST NOT occur more than once.
                  ;
                  (";" "VALUE" "=" "DURATION") /
                  (";" trigrelparam) /
                  ;
                  ; The following is OPTIONAL,
                  ; and MAY occur more than once.
                  ;
                  (";" other-param)
                  ;
                  ) ":"  dur-value

       trigabs    = *(
                  ;
                  ; The following is REQUIRED,
                  ; but MUST NOT occur more than once.
                  ;
                  (";" "VALUE" "=" "DATE-TIME") /
                  ;
                  ; The following is OPTIONAL,
                  ; and MAY occur more than once.
                  ;
                  (";" other-param)
                  ;
                  ) ":" date-time

   Example:  A trigger set 15 minutes prior to the start of the event or
      to-do.

       TRIGGER:-PT15M

      A trigger set five minutes after the end of an event or the due
      date of a to-do.

       TRIGGER;RELATED=END:PT5M


      A trigger set to an absolute DATE-TIME.

       TRIGGER;VALUE=DATE-TIME:19980101T050000Z
