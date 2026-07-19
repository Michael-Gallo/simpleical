package ical

import (
	"strings"

	"github.com/michael-gallo/simpleical/model"
)

// appendExtensionProperty appends an unrecognized property to XProp or IANAProp.
// Names starting with "X-" (case-insensitive) go to XProp; all others go to IANAProp.
// Params are cloned because the caller reuses the parameter map across lines.
func appendExtensionProperty(xProp, ianaProp *[]model.ExtensionProperty, name, value string, params map[string]string) {
	prop := model.ExtensionProperty{
		Name:  name,
		Value: value,
	}
	if len(params) > 0 {
		prop.Params = make(map[string]string, len(params))
		for k, v := range params {
			prop.Params[k] = v
		}
	}
	if isXName(name) {
		*xProp = append(*xProp, prop)
		return
	}
	*ianaProp = append(*ianaProp, prop)
}

// isXName reports whether name is an RFC 5545 x-name (prefix "X-").
func isXName(name string) bool {
	return len(name) >= 2 && strings.EqualFold(name[:2], "X-")
}
