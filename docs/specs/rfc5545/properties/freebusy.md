# Free/Busy Time

Source: RFC 5545, Section 3.8.2.6

3.8.2.6.  Free/Busy Time

   Property Name:  FREEBUSY

   Purpose:  This property defines one or more free or busy time
      intervals.

   Value Type:  PERIOD

   Property Parameters:  IANA, non-standard, and free/busy time type
      property parameters can be specified on this property.

   Conformance:  The property can be specified in a "VFREEBUSY" calendar
      component.

   Description:  These time periods can be specified as either a start
      and end DATE-TIME or a start DATE-TIME and DURATION.  The date and
      time MUST be a UTC time format.

      "FREEBUSY" properties within the "VFREEBUSY" calendar component
      SHOULD be sorted in ascending order, based on start time and then
      end time, with the earliest periods first.

      The "FREEBUSY" property can specify more than one value, separated
      by the COMMA character.  In such cases, the "FREEBUSY" property
      values MUST all be of the same "FBTYPE" property parameter type
      (e.g., all values of a particular "FBTYPE" listed together in a
      single property).

   Format Definition:  This property is defined by the following
      notation:

       freebusy   = "FREEBUSY" fbparam ":" fbvalue CRLF

       fbparam    = *(
                  ;
                  ; The following is OPTIONAL,
                  ; but MUST NOT occur more than once.
                  ;
                  (";" fbtypeparam) /
                  ;
                  ; The following is OPTIONAL,
                  ; and MAY occur more than once.
                  ;
                  (";" other-param)
                  ;
                  )


       fbvalue    = period *("," period)
       ;Time value MUST be in the UTC time format.

   Example:  The following are some examples of this property:

       FREEBUSY;FBTYPE=BUSY-UNAVAILABLE:19970308T160000Z/PT8H30M

       FREEBUSY;FBTYPE=FREE:19970308T160000Z/PT3H,19970308T200000Z/PT1H

       FREEBUSY;FBTYPE=FREE:19970308T160000Z/PT3H,19970308T200000Z/PT1H
        ,19970308T230000Z/19970309T000000Z
