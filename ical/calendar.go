package ical

import (
	"github.com/michael-gallo/simpleical/internal/icalerr"
	"github.com/michael-gallo/simpleical/model"
)

// parseCalendarProperty parses a single property line and sets its value in the provided vcalendar.
func parseCalendarProperty(propertyName string, value string, params map[string]string, calendar *model.Calendar) error {
	switch propertyName {
	case "VERSION":
		return setOnceProperty(&calendar.Version, value, propertyName, "VCALENDAR")
	case "PRODID":
		return setOnceProperty(&calendar.ProdID, value, propertyName, "VCALENDAR")
	case "CALSCALE":
		return setOnceProperty(&calendar.CalScale, value, propertyName, "VCALENDAR")
	case "METHOD":
		return setOnceProperty(&calendar.Method, value, propertyName, "VCALENDAR")
	default:
		appendExtensionProperty(&calendar.XProp, &calendar.IANAProp, propertyName, value, params)
		return nil
	}
}

func validateCalendar(calendar *model.Calendar) error {
	if calendar.Version == "" {
		return icalerr.ErrMissingCalendarVersionProperty
	}
	if calendar.ProdID == "" {
		return icalerr.ErrMissingCalendarProdIDProperty
	}
	return nil
}
