package ical

import (
	"errors"
	"fmt"
	"net/url"
	"strings"

	"github.com/michael-gallo/simpleical/icaldur"
	"github.com/michael-gallo/simpleical/internal/icalerr"
	"github.com/michael-gallo/simpleical/model"
	"github.com/michael-gallo/simpleical/rrule"
)

const timezoneLocation = "TimeZone"

// parseTimezoneProperty parses a single property line and adds it to the provided timezone.
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
	switch propertyName {
	case model.PropTZID:
		return setOnceProperty(&timezone.TimeZoneID, value, propertyName, timezoneLocation)
	case model.PropLastModified:
		return setOnceUTCTimePropertyWithParams(&timezone.LastMod, value, params, propertyName, timezoneLocation)
	case model.PropTZURL:
		parsedURL, err := url.Parse(value)
		if err != nil {
			return err
		}
		return setOnceProperty(&timezone.TimeZoneURL, parsedURL, propertyName, timezoneLocation)
	default:
		appendExtensionProperty(&timezone.XProp, &timezone.IANAProp, propertyName, value, params)
		return nil
	}
}

// parseTimeZonePropertySubComponent parses a single property line for STANDARD or DAYLIGHT sub-components.
func parseTimeZonePropertySubComponent(propertyName string, value string, params map[string]string, tzProp *model.TimeZoneProperty) error {
	switch propertyName {
	case model.PropTZOffsetFrom:
		offset, err := icaldur.ParseUTCOffset(value)
		if err != nil {
			return fmt.Errorf("%w: %s property %s in iCal: %w", icalerr.ErrParseErrorInComponent, timezoneLocation, propertyName, err)
		}
		return setOnceProperty(&tzProp.TimeZoneOffsetFrom, model.UTCOffset(offset), propertyName, timezoneLocation)
	case model.PropTZOffsetTo:
		offset, err := icaldur.ParseUTCOffset(value)
		if err != nil {
			return fmt.Errorf("%w: %s property %s in iCal: %w", icalerr.ErrParseErrorInComponent, timezoneLocation, propertyName, err)
		}
		return setOnceProperty(&tzProp.TimeZoneOffsetTo, model.UTCOffset(offset), propertyName, timezoneLocation)
	case model.PropDTStart:
		parsedTime, err := parseTimezoneLocalTime(value, propertyName)
		if err != nil {
			return err
		}
		return setOnceProperty(&tzProp.DTStart, parsedTime, propertyName, timezoneLocation)
	case model.PropComment:
		return appendTextProperty(&tzProp.Comment, value, params)
	case model.PropRDate:
		for part := range strings.SplitSeq(value, ",") {
			parsedTime, err := parseTimezoneLocalTime(part, propertyName)
			if err != nil {
				return err
			}
			tzProp.Rdate = append(tzProp.Rdate, parsedTime)
		}
	case model.PropTZName:
		tzProp.TimeZoneName = append(tzProp.TimeZoneName, value)
	case model.PropRRule:
		rule, err := rrule.ParseRRule(value)
		if err != nil {
			return fmt.Errorf("%w: %s", icalerr.ErrInvalidTimezoneProperty, err.Error())
		}
		return setOnceProperty(&tzProp.RRule, rule, propertyName, timezoneLocation)
	default:
		appendExtensionProperty(&tzProp.XProp, &tzProp.IANAProp, propertyName, value, params)
	}
	return nil
}

// parseTimezoneLocalTime parses DTSTART/RDATE values for STANDARD/DAYLIGHT.
// RFC 5545 requires these to be local wall-time values (no trailing "Z").
func parseTimezoneLocalTime(value, propertyName string) (model.DateTime, error) {
	temporal, err := icaldur.ParseTemporalDateTime(value)
	if err != nil {
		if errors.Is(err, icaldur.ErrLocalTimeRequired) {
			return model.DateTime{}, icalerr.ErrTimezoneLocalTimeRequired
		}
		return model.DateTime{}, fmt.Errorf("%w: %s property %s in iCal: %w", icalerr.ErrParseErrorInComponent, timezoneLocation, propertyName, err)
	}
	if temporal.Form == icaldur.FormUTC {
		return model.DateTime{}, icalerr.ErrTimezoneLocalTimeRequired
	}
	return model.DateTime{Form: model.DateTimeFormFloating, Time: temporal.Time}, nil
}

// validateTimeZone ensures that all required values are present for a timezone.
func validateTimeZone(timezone *model.TimeZone) error {
	if timezone.TimeZoneID == "" {
		return icalerr.ErrMissingTimezoneTZIDProperty
	}
	if len(timezone.Standard) == 0 && len(timezone.Daylight) == 0 {
		return icalerr.ErrMissingTimezoneObservance
	}
	return nil
}

// validateObservance ensures required STANDARD/DAYLIGHT properties are present.
func validateObservance(tzProp *model.TimeZoneProperty) error {
	if tzProp.DTStart.IsZero() {
		return icalerr.ErrMissingObservanceDTStart
	}
	if tzProp.TimeZoneOffsetFrom == "" {
		return icalerr.ErrMissingObservanceOffsetFrom
	}
	if tzProp.TimeZoneOffsetTo == "" {
		return icalerr.ErrMissingObservanceOffsetTo
	}
	return nil
}
