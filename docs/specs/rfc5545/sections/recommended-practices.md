# Recommended Practices

Source: RFC 5545, Section 5

5.  Recommended Practices

   These recommended practices should be followed in order to assure
   consistent handling of the following cases for an iCalendar object.

   1.  Content lines longer than 75 octets SHOULD be folded.

   2.  When the combination of the "RRULE" and "RDATE" properties in a
       recurring component produces multiple instances having the same
       start DATE-TIME value, they should be collapsed to, and
       considered as, a single instance.  If the "RDATE" property is
       specified as a PERIOD value the duration of the recurrence
       instance will be the one specified by the "RDATE" property, and
       not the duration of the recurrence instance defined by the
       "DTSTART" property.

   3.  When a calendar user receives multiple requests for the same
       calendar component (e.g., REQUEST for a "VEVENT" calendar

       component) as a result of being on multiple mailing lists
       specified by "ATTENDEE" properties in the request, they SHOULD
       respond to only one of the requests.  The calendar user SHOULD
       also specify (using the "MEMBER" parameter of the "ATTENDEE"
       property) of which mailing list they are a member.

   4.  An implementation can truncate a "SUMMARY" property value to 255
       octets, but it MUST NOT truncate the value in the middle of a
       UTF-8 multi-octet sequence.

   5.  If seconds of the minute are not supported by an implementation,
       then a value of "00" SHOULD be specified for the seconds
       component in a time value.

   6.  "TZURL" values SHOULD NOT be specified as a file URI type.  This
       URI form can be useful within an organization, but is problematic
       in the Internet.

   7.  Some possible English values for "CATEGORIES" property include:
       "ANNIVERSARY", "APPOINTMENT", "BUSINESS", "EDUCATION", "HOLIDAY",
       "MEETING", "MISCELLANEOUS", "NON-WORKING HOURS", "NOT IN OFFICE",
       "PERSONAL", "PHONE CALL", "SICK DAY", "SPECIAL OCCASION",
       "TRAVEL", "VACATION".  Categories can be specified in any
       registered language.

   8.  Some possible English values for the "RESOURCES" property
       include: "CATERING", "CHAIRS", "COMPUTER PROJECTOR", "EASEL",
       "OVERHEAD PROJECTOR", "SPEAKER PHONE", "TABLE", "TV", "VCR",
       "VIDEO PHONE", "VEHICLE".  Resources can be specified in any
       registered language.
