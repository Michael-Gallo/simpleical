package ical

import (
	"fmt"
	"net/url"

	"github.com/michael-gallo/simpleical/icaldur"
	"github.com/michael-gallo/simpleical/internal/icalerr"
	"github.com/michael-gallo/simpleical/model"
	"github.com/michael-gallo/simpleical/rrule"
)

const timezoneLocation = "TimeZone"

// parseTimezoneProperty parses a single property line and adds it to the provided timezone.
// TODO: support X-PROP and IANA-PROP
func parseTimezoneProperty(propertyName string, value string, params map[string]string, currentState parserState, timezone *model.TimeZone) error {
	// Handle sub-components (STANDARD and DAYLIGHT)
	if currentState == stateStandard || currentState == stateDaylight {
		var tzProp *model.TimeZoneProperty
		if currentState == stateStandard {
			tzProp = &timezone.Standard[len(timezone.Standard)-1]
		} else {
			tzProp = &timezone.Daylight[len(timezone.Daylight)-1]
		}
		return parseTimeZonePropertySubComponent(propertyName, value, params, tzProp)
	}

	// Handle timezone-level properties
	switch model.TimezoneToken(propertyName) {
	case model.TimezoneTokenTimeZoneID:
		return setOnceProperty(&timezone.TimeZoneID, value, propertyName, timezoneLocation)
	case model.TimezoneTokenLastMod:
		return setOnceTimeProperty(&timezone.LastMod, value, propertyName, timezoneLocation)
	case model.TimezoneTokenTimeZoneURL:
		parsedURL, err := url.Parse(value)
		if err != nil {
			return err
		}
		return setOnceProperty(&timezone.TimeZoneURL, parsedURL, propertyName, timezoneLocation)
	default:
		return fmt.Errorf("%w: %s", icalerr.ErrInvalidTimezoneProperty, propertyName)
	}
}

// parseTimeZonePropertySubComponent parses a single property line for STANDARD or DAYLIGHT sub-components.
// TODO: support X-PROP and IANA-PROP
func parseTimeZonePropertySubComponent(propertyName string, value string, _ map[string]string, tzProp *model.TimeZoneProperty) error {
	switch model.TimezonePropertyToken(propertyName) {
	case model.TimezonePropertyTokenTimeZoneOffsetFrom:
		return setOnceProperty(&tzProp.TimeZoneOffsetFrom, value, propertyName, timezoneLocation)
	case model.TimezonePropertyTokenTimeZoneOffsetTo:
		return setOnceProperty(&tzProp.TimeZoneOffsetTo, value, propertyName, timezoneLocation)
	case model.TimezonePropertyTokenDTStart:
		return setOnceTimeProperty(&tzProp.DTStart, value, propertyName, timezoneLocation)
	case model.TimezonePropertyComment:
		tzProp.Comment = append(tzProp.Comment, value)
	case model.TimezonePropertyRdate:
		parsedTime, err := icaldur.ParseIcalTime(value)
		if err != nil {
			return fmt.Errorf("%w: %s", icalerr.ErrInvalidTimezoneProperty, err.Error())
		}
		tzProp.Rdate = append(tzProp.Rdate, parsedTime)
	case model.TimezonePropertyTimeZoneName:
		tzProp.TimeZoneName = append(tzProp.TimeZoneName, value)
	case model.TimezonePropertyRRule:
		rule, err := rrule.ParseRRule(value)
		if err != nil {
			return fmt.Errorf("%w: %s", icalerr.ErrInvalidTimezoneProperty, err.Error())
		}
		return setOnceProperty(&tzProp.RRule, rule, propertyName, timezoneLocation)
	default:
		return fmt.Errorf("%w: %s", icalerr.ErrInvalidTimezoneProperty, propertyName)
	}
	return nil
}

// validateTimeZone ensures that all required values are present for a timezone.
func validateTimeZone(timezone *model.TimeZone) error {
	if timezone.TimeZoneID == "" {
		return icalerr.ErrMissingTimezoneTZIDProperty
	}
	return nil
}
