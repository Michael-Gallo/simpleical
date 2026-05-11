# Resources

Source: RFC 5545, Section 3.8.1.10

3.8.1.10.  Resources

   Property Name:  RESOURCES

   Purpose:  This property defines the equipment or resources
      anticipated for an activity specified by a calendar component.

   Value Type:  TEXT

   Property Parameters:  IANA, non-standard, alternate text
      representation, and language property parameters can be specified
      on this property.

   Conformance:  This property can be specified once in "VEVENT" or
      "VTODO" calendar component.

   Description:  The property value is an arbitrary text.  More than one
      resource can be specified as a COMMA-separated list of resources.

   Format Definition:  This property is defined by the following
      notation:

       resources  = "RESOURCES" resrcparam ":" text *("," text) CRLF

       resrcparam = *(
                  ;
                  ; The following are OPTIONAL,
                  ; but MUST NOT occur more than once.
                  ;
                  (";" altrepparam) / (";" languageparam) /
                  ;
                  ; The following is OPTIONAL,
                  ; and MAY occur more than once.
                  ;
                  (";" other-param)
                  ;
                  )

   Example:  The following is an example of this property:

       RESOURCES:EASEL,PROJECTOR,VCR

       RESOURCES;LANGUAGE=fr:Nettoyeur haute pression
