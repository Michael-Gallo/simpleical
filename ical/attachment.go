// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package ical

import (
	"encoding/base64"
	"fmt"
	"net/url"

	"github.com/michael-gallo/simpleical/model"
)

// parseAttachment parses an ATTACH property value and parameters.
// The attachment can be either a URI (default) or binary data (with ENCODING=BASE64 and VALUE=BINARY).
// For more information see https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.1.1.
func parseAttachment(value string, params map[string]string) (*model.Attachment, error) {
	attachment := &model.Attachment{}

	// Extract VALUE parameter (defaults to "URI" if not specified)
	if val, ok := params["VALUE"]; ok {
		attachment.Value = val
	} else {
		attachment.Value = "URI"
	}

	// Extract ENCODING parameter if present
	if enc, ok := params["ENCODING"]; ok {
		attachment.Encoding = enc
	}

	// Check if this is a binary attachment
	isBinary := false
	if attachment.Value == "BINARY" {
		if attachment.Encoding == "BASE64" {
			isBinary = true
		} else {
			return nil, fmt.Errorf("%w: ATTACH property with VALUE=BINARY must have ENCODING=BASE64", errParseErrorInComponent)
		}
	}

	if isBinary {
		// Verify the base64 encoded data is valid
		// Base64 strings must be a multiple of 4 characters, pad if necessary
		paddedValue := value
		for len(paddedValue)%4 != 0 {
			paddedValue += "="
		}
		_, err := base64.StdEncoding.DecodeString(paddedValue)
		if err != nil {
			return nil, fmt.Errorf("%w: invalid base64 encoded data", errParseErrorInComponent)
		}
		// Store the base64-encoded binary data (store original, not padded)
		attachment.Binary = value
	} else {
		// Parse as URI
		parsedURI, err := url.Parse(value)
		if err != nil {
			return nil, fmt.Errorf("%w: ATTACH property URI", errParseErrorInComponent)
		}
		attachment.URI = parsedURI
	}

	// Extract FMTTYPE parameter if present
	if fmtType, ok := params["FMTTYPE"]; ok {
		attachment.FormatType = fmtType
	}

	// Store other parameters (excluding VALUE, ENCODING, and FMTTYPE which we've already handled)
	attachment.OtherParams = make(map[string]string)
	for paramName, paramValue := range params {
		if paramName != "VALUE" && paramName != "ENCODING" && paramName != "FMTTYPE" {
			attachment.OtherParams[paramName] = paramValue
		}
	}

	// If no other params, set to nil to avoid empty map
	if len(attachment.OtherParams) == 0 {
		attachment.OtherParams = nil
	}

	return attachment, nil
}
