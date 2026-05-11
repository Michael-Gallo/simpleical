# Contact

Source: RFC 5545, Section 3.8.4.2

3.8.4.2.  Contact

   Property Name:  CONTACT

   Purpose:  This property is used to represent contact information or
      alternately a reference to contact information associated with the
      calendar component.

   Value Type:  TEXT

   Property Parameters:  IANA, non-standard, alternate text
      representation, and language property parameters can be specified
      on this property.


   Conformance:  This property can be specified in a "VEVENT", "VTODO",
      "VJOURNAL", or "VFREEBUSY" calendar component.

   Description:  The property value consists of textual contact
      information.  An alternative representation for the property value
      can also be specified that refers to a URI pointing to an
      alternate form, such as a vCard [RFC2426], for the contact
      information.

   Format Definition:  This property is defined by the following
      notation:

       contact    = "CONTACT" contparam ":" text CRLF

       contparam  = *(
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

   Example:  The following is an example of this property referencing
      textual contact information:

       CONTACT:Jim Dolittle\, ABC Industries\, +1-919-555-1234

      The following is an example of this property with an alternate
      representation of an LDAP URI to a directory entry containing the
      contact information:

       CONTACT;ALTREP="ldap://example.com:6666/o=ABC%20Industries\,
        c=US???(cn=Jim%20Dolittle)":Jim Dolittle\, ABC Industries\,
        +1-919-555-1234

      The following is an example of this property with an alternate
      representation of a MIME body part containing the contact
      information, such as a vCard [RFC2426] embedded in a text/
      directory media type [RFC2425]:

       CONTACT;ALTREP="CID:part3.msg970930T083000SILVER@example.com":
        Jim Dolittle\, ABC Industries\, +1-919-555-1234

      The following is an example of this property referencing a network
      resource, such as a vCard [RFC2426] object containing the contact
      information:

       CONTACT;ALTREP="http://example.com/pdi/jdoe.vcf":Jim
         Dolittle\, ABC Industries\, +1-919-555-1234
