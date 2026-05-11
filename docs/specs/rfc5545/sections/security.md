# Security Considerations

Source: RFC 5545, Section 7

7.  Security Considerations

   Because calendaring and scheduling information is very privacy-
   sensitive, the protocol used for the transmission of calendaring and
   scheduling information should have capabilities to protect the
   information from possible threats, such as eavesdropping, replay,
   message insertion, deletion, modification, and man-in-the-middle
   attacks.

   As this document only defines the data format and media type of text/
   calendar that is independent of any calendar service or protocol, it
   is up to the actual protocol specifications such as iTIP [2446bis],

   iMIP [2447bis], and "Calendaring Extensions to WebDAV (CalDAV)"
   [RFC4791] to describe the threats that the above attacks present, as
   well as ways in which to mitigate them.
