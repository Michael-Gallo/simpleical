# Categories

Source: RFC 5545, Section 3.8.1.2

3.8.1.2.  Categories

   Property Name:  CATEGORIES

   Purpose:  This property defines the categories for a calendar
      component.

   Value Type:  TEXT

   Property Parameters:  IANA, non-standard, and language property
      parameters can be specified on this property.

   Conformance:  The property can be specified within "VEVENT", "VTODO",
      or "VJOURNAL" calendar components.


   Description:  This property is used to specify categories or subtypes
      of the calendar component.  The categories are useful in searching
      for a calendar component of a particular type and category.
      Within the "VEVENT", "VTODO", or "VJOURNAL" calendar components,
      more than one category can be specified as a COMMA-separated list
      of categories.

   Format Definition:  This property is defined by the following
      notation:

       categories = "CATEGORIES" catparam ":" text *("," text)
                    CRLF

       catparam   = *(
                  ;
                  ; The following is OPTIONAL,
                  ; but MUST NOT occur more than once.
                  ;
                  (";" languageparam ) /
                  ;
                  ; The following is OPTIONAL,
                  ; and MAY occur more than once.
                  ;
                  (";" other-param)
                  ;
                  )

   Example:  The following are examples of this property:

       CATEGORIES:APPOINTMENT,EDUCATION

       CATEGORIES:MEETING
