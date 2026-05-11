# Format Type

Source: RFC 5545, Section 3.2.8

3.2.8.  Format Type

   Parameter Name:  FMTTYPE

   Purpose:  To specify the content type of a referenced object.

   Format Definition:  This property parameter is defined by the
      following notation:

       fmttypeparam = "FMTTYPE" "=" type-name "/" subtype-name
                      ; Where "type-name" and "subtype-name" are
                      ; defined in Section 4.2 of [RFC4288].

   Description:  This parameter can be specified on properties that are
      used to reference an object.  The parameter specifies the media
      type [RFC4288] of the referenced object.  For example, on the
      "ATTACH" property, an FTP type URI value does not, by itself,

      necessarily convey the type of content associated with the
      resource.  The parameter value MUST be the text for either an
      IANA-registered media type or a non-standard media type.

   Example:

       ATTACH;FMTTYPE=application/msword:ftp://example.com/pub/docs/
        agenda.doc
