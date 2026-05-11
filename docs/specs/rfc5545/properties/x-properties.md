# Non-Standard Properties

Source: RFC 5545, Section 3.8.8.2

3.8.8.2.  Non-Standard Properties

   Property Name:  Any property name with a "X-" prefix

   Purpose:  This class of property provides a framework for defining
      non-standard properties.

   Value Type:  The default value type is TEXT.  The value type can be
      set to any value type.

   Property Parameters:  IANA, non-standard, and language property
      parameters can be specified on this property.

   Conformance:  This property can be specified in any calendar
      component.

   Description:  The MIME Calendaring and Scheduling Content Type
      provides a "standard mechanism for doing non-standard things".
      This extension support is provided for implementers to "push the
      envelope" on the existing version of the memo.  Extension
      properties are specified by property and/or property parameter

      names that have the prefix text of "X-" (the two-character
      sequence: LATIN CAPITAL LETTER X character followed by the HYPHEN-
      MINUS character).  It is recommended that vendors concatenate onto
      this sentinel another short prefix text to identify the vendor.
      This will facilitate readability of the extensions and minimize
      possible collision of names between different vendors.  User
      agents that support this content type are expected to be able to
      parse the extension properties and property parameters but can
      ignore them.

      At present, there is no registration authority for names of
      extension properties and property parameters.  The value type for
      this property is TEXT.  Optionally, the value type can be any of
      the other valid value types.

   Format Definition:  This property is defined by the following
      notation:

       x-prop = x-name *(";" icalparameter) ":" value CRLF

   Example:  The following might be the ABC vendor's extension for an
      audio-clip form of subject property:

       X-ABC-MMSUBJ;VALUE=URI;FMTTYPE=audio/basic:http://www.example.
        org/mysubj.au
