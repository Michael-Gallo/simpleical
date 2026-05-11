# Inline Encoding

Source: RFC 5545, Section 3.2.7

3.2.7.  Inline Encoding

   Parameter Name:  ENCODING

   Purpose:  To specify an alternate inline encoding for the property
      value.

   Format Definition:  This property parameter is defined by the
      following notation:






       encodingparam      = "ENCODING" "="
                          ( "8BIT"
          ; "8bit" text encoding is defined in [RFC2045]
                          / "BASE64"
          ; "BASE64" binary encoding format is defined in [RFC4648]
                          )

   Description:  This property parameter identifies the inline encoding
      used in a property value.  The default encoding is "8BIT",
      corresponding to a property value consisting of text.  The
      "BASE64" encoding type corresponds to a property value encoded
      using the "BASE64" encoding defined in [RFC2045].

      If the value type parameter is ";VALUE=BINARY", then the inline
      encoding parameter MUST be specified with the value
      ";ENCODING=BASE64".

   Example:

     ATTACH;FMTTYPE=text/plain;ENCODING=BASE64;VALUE=BINARY:TG9yZW
      0gaXBzdW0gZG9sb3Igc2l0IGFtZXQsIGNvbnNlY3RldHVyIGFkaXBpc2ljaW
      5nIGVsaXQsIHNlZCBkbyBlaXVzbW9kIHRlbXBvciBpbmNpZGlkdW50IHV0IG
      xhYm9yZSBldCBkb2xvcmUgbWFnbmEgYWxpcXVhLiBVdCBlbmltIGFkIG1pbm
      ltIHZlbmlhbSwgcXVpcyBub3N0cnVkIGV4ZXJjaXRhdGlvbiB1bGxhbWNvIG
      xhYm9yaXMgbmlzaSB1dCBhbGlxdWlwIGV4IGVhIGNvbW1vZG8gY29uc2VxdW
      F0LiBEdWlzIGF1dGUgaXJ1cmUgZG9sb3IgaW4gcmVwcmVoZW5kZXJpdCBpbi
      B2b2x1cHRhdGUgdmVsaXQgZXNzZSBjaWxsdW0gZG9sb3JlIGV1IGZ1Z2lhdC
      BudWxsYSBwYXJpYXR1ci4gRXhjZXB0ZXVyIHNpbnQgb2NjYWVjYXQgY3VwaW
      RhdGF0IG5vbiBwcm9pZGVudCwgc3VudCBpbiBjdWxwYSBxdWkgb2ZmaWNpYS
      BkZXNlcnVudCBtb2xsaXQgYW5pbSBpZCBlc3QgbGFib3J1bS4=
