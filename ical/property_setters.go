package ical

import (
	"fmt"
	"strconv"
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

// setOnceTimeProperty sets a time.Time field only if it hasn't been set before.
// this is intended for properties that according to the spec must only be set once
func setOnceTimeProperty(field *time.Time, value, propertyName string, componentType string) error {
	return setOnceTimePropertyWithParams(field, value, nil, propertyName, componentType)
}

// setOnceTimePropertyWithParams is like setOnceTimeProperty but honors VALUE=DATE.
func setOnceTimePropertyWithParams(field *time.Time, value string, params map[string]string, propertyName string, componentType string) error {
	parsedTime, err := icaldur.ParseIcalTimeOrDate(value, params[model.ParamValue])
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

// appendTimePropertyWithParams appends a parsed time, honoring VALUE=DATE when present.
func appendTimePropertyWithParams(field *[]time.Time, value string, params map[string]string, propertyName string, componentType string) error {
	parsedTime, err := icaldur.ParseIcalTimeOrDate(value, params[model.ParamValue])
	if err != nil {
		return fmt.Errorf("%w: %s property %s in iCal", icalerr.ErrParseErrorInComponent, componentType, propertyName)
	}
	*field = append(*field, parsedTime)
	return nil
}
