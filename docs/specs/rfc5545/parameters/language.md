# Language

Source: RFC 5545, Section 3.2.10

3.2.10.  Language

   Parameter Name:  LANGUAGE

   Purpose:  To specify the language for text values in a property or
      property parameter.

   Format Definition:  This property parameter is defined by the
      following notation:

       languageparam = "LANGUAGE" "=" language

       language = Language-Tag
                  ; As defined in [RFC5646].

   Description:  This parameter identifies the language of the text in
      the property value and of all property parameter values of the
      property.  The value of the "LANGUAGE" property parameter is that
      defined in [RFC5646].

      For transport in a MIME entity, the Content-Language header field
      can be used to set the default language for the entire body part.
      Otherwise, no default language is assumed.

   Example:  The following are examples of this parameter on the
      "SUMMARY" and "LOCATION" properties:

       SUMMARY;LANGUAGE=en-US:Company Holiday Party

       LOCATION;LANGUAGE=en:Germany

       LOCATION;LANGUAGE=no:Tyskland
