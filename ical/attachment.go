// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package ical

import (
	"encoding/base64"
	"fmt"
	"net/url"

	"github.com/michael-gallo/simpleical/internal/icalerr"
	"github.com/michael-gallo/simpleical/model"
)

// parseAttachment parses an ATTACH property value and parameters.
// The attachment can be either a URI (default) or binary data (with ENCODING=BASE64 and VALUE=BINARY).
// For more information see https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.1.1.
func parseAttachment(value string, params map[string]string) (*model.Attachment, error) {
	attachment := &model.Attachment{}

	// Extract VALUE parameter (defaults to "URI" if not specified).
	// Validate against the closed set allowed for ATTACH (URI | BINARY).
	if val, ok := params[model.ParamValue]; ok {
		switch v := model.AttachValue(val); v {
		case model.AttachValueURI, model.AttachValueBinary:
			attachment.Value = v
		default:
			return nil, fmt.Errorf("%w: invalid VALUE parameter %q for ATTACH", icalerr.ErrParseErrorInComponent, val)
		}
	} else {
		attachment.Value = model.AttachValueURI
	}

	// Extract ENCODING parameter if present
	if enc, ok := params[model.ParamEncoding]; ok {
		attachment.Encoding = enc
	}

	// Check if this is a binary attachment
	isBinary := false
	if attachment.Value == model.AttachValueBinary {
		if attachment.Encoding == model.EncodingBase64 {
			isBinary = true
		} else {
			return nil, fmt.Errorf("%w: ATTACH property with VALUE=BINARY must have ENCODING=BASE64", icalerr.ErrParseErrorInComponent)
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
			return nil, fmt.Errorf("%w: invalid base64 encoded data", icalerr.ErrParseErrorInComponent)
		}
		// Store the base64-encoded binary data (store original, not padded)
		attachment.Binary = value
	} else {
		// Parse as URI
		parsedURI, err := url.Parse(value)
		if err != nil {
			return nil, fmt.Errorf("%w: ATTACH property URI", icalerr.ErrParseErrorInComponent)
		}
		attachment.URI = parsedURI
	}

	// Extract FMTTYPE parameter if present
	if fmtType, ok := params[model.ParamFmtType]; ok {
		attachment.FormatType = fmtType
	}

	// Store other parameters only when unknown params exist.
	for paramName, paramValue := range params {
		if paramName == model.ParamValue ||
			paramName == model.ParamEncoding ||
			paramName == model.ParamFmtType {
			continue
		}
		if attachment.OtherParams == nil {
			attachment.OtherParams = make(map[string]string, 1)
		}
		attachment.OtherParams[paramName] = paramValue
	}

	return attachment, nil
}
