# Text

Source: RFC 5545, Section 3.3.11

3.3.11.  Text

   Value Name:  TEXT

   Purpose:  This value type is used to identify values that contain
      human-readable text.

   Format Definition:  This value type is defined by the following
      notation:



       text       = *(TSAFE-CHAR / ":" / DQUOTE / ESCAPED-CHAR)
          ; Folded according to description above

       ESCAPED-CHAR = ("\\" / "\;" / "\," / "\N" / "\n")
          ; \\ encodes \, \N or \n encodes newline
          ; \; encodes ;, \, encodes ,

       TSAFE-CHAR = WSP / %x21 / %x23-2B / %x2D-39 / %x3C-5B /
                    %x5D-7E / NON-US-ASCII
          ; Any character except CONTROLs not needed by the current
          ; character set, DQUOTE, ";", ":", "\", ","

   Description:  If the property permits, multiple TEXT values are
      specified by a COMMA-separated list of values.

      The language in which the text is represented can be controlled by
      the "LANGUAGE" property parameter.

      An intentional formatted text line break MUST only be included in
      a "TEXT" property value by representing the line break with the
      character sequence of BACKSLASH, followed by a LATIN SMALL LETTER
      N or a LATIN CAPITAL LETTER N, that is "\n" or "\N".

      The "TEXT" property values may also contain special characters
      that are used to signify delimiters, such as a COMMA character for
      lists of values or a SEMICOLON character for structured values.
      In order to support the inclusion of these special characters in
      "TEXT" property values, they MUST be escaped with a BACKSLASH
      character.  A BACKSLASH character in a "TEXT" property value MUST
      be escaped with another BACKSLASH character.  A COMMA character in
      a "TEXT" property value MUST be escaped with a BACKSLASH
      character.  A SEMICOLON character in a "TEXT" property value MUST
      be escaped with a BACKSLASH character.  However, a COLON character
      in a "TEXT" property value SHALL NOT be escaped with a BACKSLASH
      character.

   Example:  A multiple line value of:

       Project XYZ Final Review
       Conference Room - 3B
       Come Prepared.

      would be represented as:

       Project XYZ Final Review\nConference Room - 3B\nCome Prepared.
