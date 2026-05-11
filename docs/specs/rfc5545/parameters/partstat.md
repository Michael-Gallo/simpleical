# Participation Status

Source: RFC 5545, Section 3.2.12

3.2.12.  Participation Status

   Parameter Name:  PARTSTAT

   Purpose:  To specify the participation status for the calendar user
      specified by the property.

   Format Definition:  This property parameter is defined by the
      following notation:

       partstatparam    = "PARTSTAT" "="
                         (partstat-event
                        / partstat-todo
                        / partstat-jour)

       partstat-event   = ("NEEDS-ACTION"    ; Event needs action
                        / "ACCEPTED"         ; Event accepted
                        / "DECLINED"         ; Event declined
                        / "TENTATIVE"        ; Event tentatively
                                             ; accepted
                        / "DELEGATED"        ; Event delegated
                        / x-name             ; Experimental status
                        / iana-token)        ; Other IANA-registered
                                             ; status
       ; These are the participation statuses for a "VEVENT".
       ; Default is NEEDS-ACTION.

       partstat-todo    = ("NEEDS-ACTION"    ; To-do needs action
                        / "ACCEPTED"         ; To-do accepted
                        / "DECLINED"         ; To-do declined
                        / "TENTATIVE"        ; To-do tentatively
                                             ; accepted

                        / "DELEGATED"        ; To-do delegated
                        / "COMPLETED"        ; To-do completed
                                             ; COMPLETED property has
                                             ; DATE-TIME completed
                        / "IN-PROCESS"       ; To-do in process of
                                             ; being completed
                        / x-name             ; Experimental status
                        / iana-token)        ; Other IANA-registered
                                             ; status
       ; These are the participation statuses for a "VTODO".
       ; Default is NEEDS-ACTION.



       partstat-jour    = ("NEEDS-ACTION"    ; Journal needs action
                        / "ACCEPTED"         ; Journal accepted
                        / "DECLINED"         ; Journal declined
                        / x-name             ; Experimental status
                        / iana-token)        ; Other IANA-registered
                                             ; status
       ; These are the participation statuses for a "VJOURNAL".
       ; Default is NEEDS-ACTION.

   Description:  This parameter can be specified on properties with a
      CAL-ADDRESS value type.  The parameter identifies the
      participation status for the calendar user specified by the
      property value.  The parameter values differ depending on whether
      they are associated with a group-scheduled "VEVENT", "VTODO", or
      "VJOURNAL".  The values MUST match one of the values allowed for
      the given calendar component.  If not specified on a property that
      allows this parameter, the default value is NEEDS-ACTION.
      Applications MUST treat x-name and iana-token values they don't
      recognize the same way as they would the NEEDS-ACTION value.

   Example:

       ATTENDEE;PARTSTAT=DECLINED:mailto:jsmith@example.com
