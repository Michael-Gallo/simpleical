# Property

Source: RFC 5545, Section 3.5

3.5.  Property

   A property is the definition of an individual attribute describing a
   calendar object or a calendar component.  A property takes the form
   defined by the "contentline" notation defined in Section 3.1.

   The following is an example of a property:

       DTSTART:19960415T133000Z

   This memo imposes no ordering of properties within an iCalendar
   object.

   Property names, parameter names, and enumerated parameter values are
   case-insensitive.  For example, the property name "DUE" is the same
   as "due" and "Due", DTSTART;TZID=America/New_York:19980714T120000 is
   the same as DtStart;TzID=America/New_York:19980714T120000.
