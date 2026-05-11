# Classification

Source: RFC 5545, Section 3.8.1.3

3.8.1.3.  Classification

   Property Name:  CLASS

   Purpose:  This property defines the access classification for a
      calendar component.

   Value Type:  TEXT

   Property Parameters:  IANA and non-standard property parameters can
      be specified on this property.

   Conformance:  The property can be specified once in a "VEVENT",
      "VTODO", or "VJOURNAL" calendar components.


   Description:  An access classification is only one component of the
      general security system within a calendar application.  It
      provides a method of capturing the scope of the access the
      calendar owner intends for information within an individual
      calendar entry.  The access classification of an individual
      iCalendar component is useful when measured along with the other
      security components of a calendar system (e.g., calendar user
      authentication, authorization, access rights, access role, etc.).
      Hence, the semantics of the individual access classifications
      cannot be completely defined by this memo alone.  Additionally,
      due to the "blind" nature of most exchange processes using this
      memo, these access classifications cannot serve as an enforcement
      statement for a system receiving an iCalendar object.  Rather,
      they provide a method for capturing the intention of the calendar
      owner for the access to the calendar component.  If not specified
      in a component that allows this property, the default value is
      PUBLIC.  Applications MUST treat x-name and iana-token values they
      don't recognize the same way as they would the PRIVATE value.

   Format Definition:  This property is defined by the following
      notation:

       class      = "CLASS" classparam ":" classvalue CRLF

       classparam = *(";" other-param)

       classvalue = "PUBLIC" / "PRIVATE" / "CONFIDENTIAL" / iana-token
                  / x-name
       ;Default is PUBLIC

   Example:  The following is an example of this property:

       CLASS:PUBLIC
