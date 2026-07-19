package ical

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/michael-gallo/simpleical/icaldur"
	"github.com/michael-gallo/simpleical/internal/icalerr"
	"github.com/michael-gallo/simpleical/model"
)

// setOnceProperty ensures that set-once properties have consistent error handling
func setOnceProperty[T comparable](field *T, value T, propertyName string, componentType string) error {
	var zero T
	if *field != zero {
		return fmt.Errorf(icalerr.ErrDuplicatePropertyInComponentFormat, icalerr.ErrDuplicatePropertyInComponent, propertyName, componentType)
	}
	*field = value
	return nil
}

// setOnceIntProperty sets an int field only if it hasn't been set before.
// this is intended for properties that according to the spec must only be set once
func setOnceIntProperty(field *int, value, propertyName string, componentType string) error {
	intValue, err := strconv.Atoi(value)
	if err != nil {
		return fmt.Errorf("%w: %s property %s in iCal", icalerr.ErrParseErrorInComponent, componentType, propertyName)
	}
	return setOnceProperty(field, intValue, propertyName, componentType)
}

// dateValueAllowed reports whether propertyName may use VALUE=DATE per RFC 5545.
func dateValueAllowed(propertyName string) bool {
	switch propertyName {
	case string(model.EventTokenDtstart),
		string(model.EventTokenDtend),
		string(model.TodoTokenDue),
		string(model.TodoTokenRecurrenceID),
		string(model.TodoTokenExceptionDates),
		string(model.TodoTokenRdate):
		return true
	default:
		return false
	}
}

func timeValueType(propertyName string, params map[string]string) string {
	if !dateValueAllowed(propertyName) {
		return ""
	}
	return params[model.ParamValue]
}

// setOnceTimePropertyWithParams sets a time.Time field only if it hasn't been set before.
// VALUE=DATE is honored only for properties that permit DATE values.
func setOnceTimePropertyWithParams(field *time.Time, value string, params map[string]string, propertyName string, componentType string) error {
	parsedTime, err := icaldur.ParseIcalTimeOrDate(value, timeValueType(propertyName, params))
	if err != nil {
		return fmt.Errorf("%w: %s property %s in iCal", icalerr.ErrParseErrorInComponent, componentType, propertyName)
	}
	return setOnceProperty(field, parsedTime, propertyName, componentType)
}

// setOnceDurationProperty sets a duration field only if it hasn't been set before.
// this is intended for properties that according to the spec must only be set once
func setOnceDurationProperty(field *time.Duration, value, propertyName string, componentType string) error {
	duration, err := icaldur.ParseICalDuration(value)
	if err != nil {
		return fmt.Errorf("%w: %s property %s in iCal", icalerr.ErrParseErrorInComponent, componentType, propertyName)
	}
	return setOnceProperty(field, duration, propertyName, componentType)
}

// appendTimePropertyWithParams appends a parsed time.
// VALUE=DATE is honored only for properties that permit DATE values.
func appendTimePropertyWithParams(field *[]time.Time, value string, params map[string]string, propertyName string, componentType string) error {
	parsedTime, err := icaldur.ParseIcalTimeOrDate(value, timeValueType(propertyName, params))
	if err != nil {
		return fmt.Errorf("%w: %s property %s in iCal", icalerr.ErrParseErrorInComponent, componentType, propertyName)
	}
	*field = append(*field, parsedTime)
	return nil
}

// appendCommaSeparatedTimePropertyWithParams appends each comma-separated DATE/DATE-TIME
// from value (as used by EXDATE and RDATE).
func appendCommaSeparatedTimePropertyWithParams(field *[]time.Time, value string, params map[string]string, propertyName string, componentType string) error {
	for part := range strings.SplitSeq(value, ",") {
		if err := appendTimePropertyWithParams(field, part, params, propertyName, componentType); err != nil {
			return err
		}
	}
	return nil
}
