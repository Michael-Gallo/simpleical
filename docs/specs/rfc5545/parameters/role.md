# Participation Role

Source: RFC 5545, Section 3.2.16

3.2.16.  Participation Role

   Parameter Name:  ROLE

   Purpose:  To specify the participation role for the calendar user
      specified by the property.

   Format Definition:  This property parameter is defined by the
      following notation:



       roleparam  = "ROLE" "="
                   ("CHAIR"             ; Indicates chair of the
                                        ; calendar entity
                  / "REQ-PARTICIPANT"   ; Indicates a participant whose
                                        ; participation is required
                  / "OPT-PARTICIPANT"   ; Indicates a participant whose
                                        ; participation is optional
                  / "NON-PARTICIPANT"   ; Indicates a participant who
                                        ; is copied for information
                                        ; purposes only
                  / x-name              ; Experimental role
                  / iana-token)         ; Other IANA role
       ; Default is REQ-PARTICIPANT

   Description:  This parameter can be specified on properties with a
      CAL-ADDRESS value type.  The parameter specifies the participation
      role for the calendar user specified by the property in the group
      schedule calendar component.  If not specified on a property that
      allows this parameter, the default value is REQ-PARTICIPANT.
      Applications MUST treat x-name and iana-token values they don't
      recognize the same way as they would the REQ-PARTICIPANT value.

   Example:

       ATTENDEE;ROLE=CHAIR:mailto:mrbig@example.com
