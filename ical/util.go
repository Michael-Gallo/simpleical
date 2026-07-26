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
	colonIndex := findUnquotedColonIndex(line)
	if colonIndex == -1 {
		err = fmt.Errorf("%w: %s", icalerr.ErrInvalidPropertyLine, line)
		return "", nil, "", err
	}

	beforeColon := line[:colonIndex]
	propertyName = beforeColon
	if before, after, ok := strings.Cut(beforeColon, ";"); ok {
		propertyName = before
		paramString := after
		if paramString != "" {
			params = reusableParams
			if err := splitParametersWithReusableMap(paramString, params); err != nil {
				return "", nil, "", err
			}
		}
	}

	return propertyName, params, line[colonIndex+1:], nil
}

// splitParametersWithReusableMap splits parameters using a reusable map.
// Parameter names are uppercased (RFC 5545 case-insensitive). Duplicate names error.
func splitParametersWithReusableMap(paramString string, params map[string]string) error {
	keyStart := 0
	valStart := -1
	inQuotes := false

	store := func(rawKey, rawVal string) error {
		key := strings.ToUpper(rawKey)
		if _, exists := params[key]; exists {
			return fmt.Errorf("%w: %s", icalerr.ErrDuplicateParameter, key)
		}
		params[key] = unquoteParam(rawVal)
		return nil
	}

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
				if err := store(paramString[keyStart:valStart-1], paramString[valStart:i]); err != nil {
					return err
				}
			}
			keyStart = i + 1
			valStart = -1
		}
	}
	if valStart >= 0 {
		if err := store(paramString[keyStart:valStart-1], paramString[valStart:]); err != nil {
			return err
		}
	}
	return nil
}

// unquoteParam strips surrounding DQUOTE if present (RFC 5545 paramtext / quoted-string).
func unquoteParam(s string) string {
	if len(s) >= 2 && s[0] == '"' && s[len(s)-1] == '"' {
		return s[1 : len(s)-1]
	}
	return s
}

// parseGeo parses a GEO property value "lat;lng" into two floats with range checks.
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
	if latitude < -90 || latitude > 90 {
		return [2]float64{}, icalerr.ErrGeoLatitudeOutOfRange
	}
	if longitude < -180 || longitude > 180 {
		return [2]float64{}, icalerr.ErrGeoLongitudeOutOfRange
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
