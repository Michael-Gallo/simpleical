# IANA Properties

Source: RFC 5545, Section 3.8.8.1

3.8.8.1.  IANA Properties

   Property Name:  An IANA-registered property name

   Value Type:  The default value type is TEXT.  The value type can be
      set to any value type.

   Property Parameters:  Any parameter can be specified on this
      property.

   Description:  This specification allows other properties registered
      with IANA to be specified in any calendar components.  Compliant
      applications are expected to be able to parse these other IANA-
      registered properties but can ignore them.

   Format Definition:  This property is defined by the following
      notation:

       iana-prop = iana-token *(";" icalparameter) ":" value CRLF

   Example:  The following are examples of properties that might be
      registered to IANA:

       DRESSCODE:CASUAL

       NON-SMOKING;VALUE=BOOLEAN:TRUE
