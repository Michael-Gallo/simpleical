# VCALENDAR

Source: RFC 5545, Section 3.4

3.4.  iCalendar Object

   The Calendaring and Scheduling Core Object is a collection of
   calendaring and scheduling information.  Typically, this information
   will consist of an iCalendar stream with a single iCalendar object.
   However, multiple iCalendar objects can be sequentially grouped
   together in an iCalendar stream.  The first line and last line of the
   iCalendar object MUST contain a pair of iCalendar object delimiter
   strings.  The syntax for an iCalendar stream is as follows:

       icalstream = 1*icalobject

       icalobject = "BEGIN" ":" "VCALENDAR" CRLF
                    icalbody
                    "END" ":" "VCALENDAR" CRLF

   The following is a simple example of an iCalendar object:

       BEGIN:VCALENDAR
       VERSION:2.0
       PRODID:-//hacksw/handcal//NONSGML v1.0//EN
       BEGIN:VEVENT
       UID:19970610T172345Z-AF23B2@example.com
       DTSTAMP:19970610T172345Z
       DTSTART:19970714T170000Z
       DTEND:19970715T040000Z
       SUMMARY:Bastille Day Party
       END:VEVENT
       END:VCALENDAR
