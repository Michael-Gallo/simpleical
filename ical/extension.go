package ical

import (
	"maps"
	"strings"

	"github.com/michael-gallo/simpleical/model"
)

// appendExtensionProperty appends an unrecognized property to XProp or IANAProp.
// Names starting with "X-" go to XProp; all others go to IANAProp.
// Property names are uppercased before this is called.
// Params are cloned because the caller reuses the parameter map across lines.
func appendExtensionProperty(xProp, ianaProp *[]model.ExtensionProperty, name, value string, params map[string]string) {
	prop := model.ExtensionProperty{
		Name:  name,
		Value: value,
	}
	if len(params) > 0 {
		prop.Params = maps.Clone(params)
	}
	if strings.HasPrefix(name, "X-") {
		*xProp = append(*xProp, prop)
		return
	}
	*ianaProp = append(*ianaProp, prop)
}
