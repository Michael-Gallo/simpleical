package ical

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/michael-gallo/simpleical/internal/icalerr"
)

// parseIcalLineWithReusableMap parses a single property line using a reusable parameter map.
// This avoids allocating a new map for every property with parameters.
func parseIcalLineWithReusableMap(line string, reusableParams map[string]string) (propertyName string, params map[string]string, value string, err error) {
	// Find the first colon that is not inside quotes
	colonIndex := findUnquotedColonIndex(line)
	if colonIndex == -1 {
		err = fmt.Errorf("%w: %s", icalerr.ErrInvalidPropertyLine, line)
		return "", nil, "", err
	}

	// Split the line at the colon
	beforeColon := line[:colonIndex]

	// The property name is the first part before any semicolon
	propertyName = beforeColon
	if before, after, ok := strings.Cut(beforeColon, ";"); ok {
		propertyName = before
		// Extract parameters from the part between property name and colon
		paramString := after
		if paramString != "" {
			// Use the reusable map (caller has already cleared it)
			params = reusableParams
			splitParametersWithReusableMap(paramString, params)
		}
	}

	return propertyName, params, line[colonIndex+1:], nil
}

// splitParametersWithReusableMap splits parameters using a reusable map.
// Byte-oriented scan avoids rune decoding; keys/values are sliced from the input string.
func splitParametersWithReusableMap(paramString string, params map[string]string) {
	keyStart := 0
	valStart := -1
	inQuotes := false

	for i := 0; i < len(paramString); i++ {
		c := paramString[i]
		switch c {
		case '"':
			inQuotes = !inQuotes
		case '=':
			if inQuotes || valStart >= 0 {
				continue
			}
			valStart = i + 1
		case ';':
			if inQuotes {
				continue
			}
			if valStart >= 0 {
				params[paramString[keyStart:valStart-1]] = unquoteParam(paramString[valStart:i])
			}
			keyStart = i + 1
			valStart = -1
		}
	}
	if valStart >= 0 {
		params[paramString[keyStart:valStart-1]] = unquoteParam(paramString[valStart:])
	}
}

// unquoteParam strips surrounding DQUOTE if present (RFC 5545 paramtext / quoted-string).
func unquoteParam(s string) string {
	if len(s) >= 2 && s[0] == '"' && s[len(s)-1] == '"' {
		return s[1 : len(s)-1]
	}
	return s
}

// parseGeo parses a GEO property value "lat;lng" into two floats.
func parseGeo(value string) ([2]float64, error) {
	latitudeString, longitudeString, found := strings.Cut(value, ";")
	if !found {
		return [2]float64{}, icalerr.ErrInvalidGeoProperty
	}
	latitude, err := strconv.ParseFloat(latitudeString, 64)
	if err != nil {
		return [2]float64{}, icalerr.ErrInvalidGeoPropertyLatitude
	}
	longitude, err := strconv.ParseFloat(longitudeString, 64)
	if err != nil {
		return [2]float64{}, icalerr.ErrInvalidGeoPropertyLongitude
	}
	return [2]float64{latitude, longitude}, nil
}

// findUnquotedColonIndex finds the first colon that is not encapsulated in quotations.
func findUnquotedColonIndex(line string) int {
	inQuotes := false
	for i := 0; i < len(line); i++ {
		c := line[i]
		if c == '"' {
			inQuotes = !inQuotes
		} else if c == ':' && !inQuotes {
			return i
		}
	}
	return -1
}
