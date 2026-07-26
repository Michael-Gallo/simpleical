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
func setOnceIntProperty(field *int, value, propertyName string, componentType string) error {
	intValue, err := strconv.Atoi(value)
	if err != nil {
		return fmt.Errorf("%w: %s property %s in iCal", icalerr.ErrParseErrorInComponent, componentType, propertyName)
	}
	return setOnceProperty(field, intValue, propertyName, componentType)
}

// setOnceIntPtrProperty sets a *int field, distinguishing absent from explicitly zero.
func setOnceIntPtrProperty(field **int, value, propertyName string, componentType string) error {
	if *field != nil {
		return fmt.Errorf(icalerr.ErrDuplicatePropertyInComponentFormat, icalerr.ErrDuplicatePropertyInComponent, propertyName, componentType)
	}
	intValue, err := strconv.Atoi(value)
	if err != nil {
		return fmt.Errorf("%w: %s property %s in iCal", icalerr.ErrParseErrorInComponent, componentType, propertyName)
	}
	*field = &intValue
	return nil
}

// setOnceBoundedInt parses an int, enforces [lo, hi], then setOnceProperty.
func setOnceBoundedInt(field *int, value string, lo, hi int, rangeErr error, propertyName, componentType string) error {
	intValue, err := strconv.Atoi(value)
	if err != nil {
		return fmt.Errorf("%w: %s property %s in iCal", icalerr.ErrParseErrorInComponent, componentType, propertyName)
	}
	if intValue < lo || intValue > hi {
		return rangeErr
	}
	return setOnceProperty(field, intValue, propertyName, componentType)
}

const recurrenceIDRangeThisAndFuture = "THISANDFUTURE"

// parseRecurrenceIDRange validates the RANGE parameter on RECURRENCE-ID.
// Only an omitted value or THISANDFUTURE is allowed (RFC 5545 §3.2.13).
func parseRecurrenceIDRange(params map[string]string) (string, error) {
	r := params[model.ParamRange]
	if r == "" || r == recurrenceIDRangeThisAndFuture {
		return r, nil
	}
	return "", fmt.Errorf("%w: RANGE %s", icalerr.ErrInvalidEnumValue, r)
}

// dateValueAllowed reports whether propertyName may carry VALUE=DATE.
func dateValueAllowed(propertyName string) bool {
	switch propertyName {
	case model.PropDTStart,
		model.PropDTEnd,
		model.PropDue,
		model.PropRecurrenceID,
		model.PropExDate,
		model.PropRDate:
		return true
	default:
		return false
	}
}

// timeValueType returns the VALUE parameter for date-capable properties, else "".
func timeValueType(propertyName string, params map[string]string) string {
	if !dateValueAllowed(propertyName) {
		return ""
	}
	return params[model.ParamValue]
}

// parseDateTimeValue parses a DATE or DATE-TIME value and applies TZID when present.
func parseDateTimeValue(value string, params map[string]string, propertyName string) (model.DateTime, error) {
	valueType := timeValueType(propertyName, params)
	temporal, err := icaldur.ParseTemporal(value, valueType)
	if err != nil {
		return model.DateTime{}, err
	}
	tzid := params[model.ParamTZID]
	if tzid != "" && (temporal.Form == icaldur.FormUTC || temporal.Form == icaldur.FormDate) {
		// RFC 5545 3.2.19: TZID MUST NOT be applied to DATE or UTC DATE-TIME.
		return model.DateTime{}, icaldur.ErrTZIDWithUTC
	}
	switch temporal.Form {
	case icaldur.FormDate:
		return model.DateTime{Form: model.DateTimeFormDate, Time: temporal.Time}, nil
	case icaldur.FormUTC:
		return model.DateTime{Form: model.DateTimeFormUTC, Time: temporal.Time}, nil
	case icaldur.FormFloating:
		if tzid != "" {
			return model.DateTime{Form: model.DateTimeFormLocalTZ, Time: temporal.Time, TZID: tzid}, nil
		}
		return model.DateTime{Form: model.DateTimeFormFloating, Time: temporal.Time}, nil
	default:
		return model.DateTime{}, icalerr.ErrParseErrorInComponent
	}
}

// parseUTCDateTimeValue parses a DATE-TIME and requires UTC form (trailing Z).
func parseUTCDateTimeValue(value string, params map[string]string, propertyName string) (model.DateTime, error) {
	dt, err := parseDateTimeValue(value, params, propertyName)
	if err != nil {
		return model.DateTime{}, err
	}
	if dt.Form != model.DateTimeFormUTC {
		return model.DateTime{}, icalerr.ErrUTCValueRequired
	}
	return dt, nil
}

// setOnceTimePropertyWithParams sets a DateTime field only if it hasn't been set before.
func setOnceTimePropertyWithParams(field *model.DateTime, value string, params map[string]string, propertyName string, componentType string) error {
	parsed, err := parseDateTimeValue(value, params, propertyName)
	if err != nil {
		return fmt.Errorf("%w: %s property %s in iCal: %w", icalerr.ErrParseErrorInComponent, componentType, propertyName, err)
	}
	return setOnceProperty(field, parsed, propertyName, componentType)
}

// setOnceUTCTimePropertyWithParams sets a UTC-only DateTime field.
func setOnceUTCTimePropertyWithParams(field *model.DateTime, value string, params map[string]string, propertyName string, componentType string) error {
	parsed, err := parseUTCDateTimeValue(value, params, propertyName)
	if err != nil {
		return fmt.Errorf("%w: %s property %s in iCal: %w", icalerr.ErrParseErrorInComponent, componentType, propertyName, err)
	}
	return setOnceProperty(field, parsed, propertyName, componentType)
}

// setOncePositiveDurationProperty requires a positive duration.
func setOncePositiveDurationProperty(field *time.Duration, value, propertyName string, componentType string) error {
	duration, err := icaldur.ParseICalDuration(value)
	if err != nil {
		return fmt.Errorf("%w: %s property %s in iCal", icalerr.ErrParseErrorInComponent, componentType, propertyName)
	}
	if duration <= 0 {
		return icalerr.ErrPositiveDurationRequired
	}
	return setOnceProperty(field, duration, propertyName, componentType)
}

// appendTimePropertyWithParams parses one DATE/DATE-TIME and appends it to field.
func appendTimePropertyWithParams(field *[]model.DateTime, value string, params map[string]string, propertyName string, componentType string) error {
	parsed, err := parseDateTimeValue(value, params, propertyName)
	if err != nil {
		return fmt.Errorf("%w: %s property %s in iCal: %w", icalerr.ErrParseErrorInComponent, componentType, propertyName, err)
	}
	*field = append(*field, parsed)
	return nil
}

// appendCommaSeparatedTimePropertyWithParams appends each comma-separated DATE/DATE-TIME value.
func appendCommaSeparatedTimePropertyWithParams(field *[]model.DateTime, value string, params map[string]string, propertyName string, componentType string) error {
	for part := range strings.SplitSeq(value, ",") {
		if err := appendTimePropertyWithParams(field, part, params, propertyName, componentType); err != nil {
			return err
		}
	}
	return nil
}

// parseTextValue unescapes TEXT and captures LANGUAGE/ALTREP parameters.
func parseTextValue(value string, params map[string]string) (model.TextValue, error) {
	unescaped, err := unescapeText(value)
	if err != nil {
		return model.TextValue{}, err
	}
	return model.TextValue{
		Value:    unescaped,
		Language: params[model.ParamLanguage],
		Altrep:   params[model.ParamAltrep],
	}, nil
}

// setOnceTextProperty parses TEXT and sets field only if it has not been set.
func setOnceTextProperty(field *model.TextValue, value string, params map[string]string, propertyName string, componentType string) error {
	tv, err := parseTextValue(value, params)
	if err != nil {
		return err
	}
	return setOnceProperty(field, tv, propertyName, componentType)
}

// appendTextProperty parses TEXT and appends it to a multi-valued property.
func appendTextProperty(field *[]model.TextValue, value string, params map[string]string) error {
	tv, err := parseTextValue(value, params)
	if err != nil {
		return err
	}
	*field = append(*field, tv)
	return nil
}

// appendTextListProperty splits an unescaped-comma TEXT list and appends the parts.
func appendTextListProperty(field *[]string, value string) error {
	parts, err := splitUnescapedComma(value)
	if err != nil {
		return err
	}
	*field = append(*field, parts...)
	return nil
}

// appendRelatedToProperty parses RELATED-TO and appends value plus RELTYPE.
func appendRelatedToProperty(field *[]model.RelatedToValue, value string, params map[string]string) error {
	unescaped, err := unescapeText(value)
	if err != nil {
		return err
	}
	*field = append(*field, model.RelatedToValue{
		Value:   unescaped,
		RelType: params[model.ParamRelType],
	})
	return nil
}

// appendRDateProperty appends RDATE DATE/DATE-TIME or PERIOD values to field.
func appendRDateProperty(field *[]model.RecurrenceDate, value string, params map[string]string, propertyName string, componentType string) error {
	if strings.EqualFold(params[model.ParamValue], "PERIOD") {
		for part := range strings.SplitSeq(value, ",") {
			period, err := parsePeriod(part, params, propertyName)
			if err != nil {
				return fmt.Errorf("%w: %s property %s in iCal: %w", icalerr.ErrParseErrorInComponent, componentType, propertyName, err)
			}
			p := period
			*field = append(*field, model.RecurrenceDate{Period: &p})
		}
		return nil
	}
	for part := range strings.SplitSeq(value, ",") {
		dt, err := parseDateTimeValue(part, params, propertyName)
		if err != nil {
			return fmt.Errorf("%w: %s property %s in iCal: %w", icalerr.ErrParseErrorInComponent, componentType, propertyName, err)
		}
		d := dt
		*field = append(*field, model.RecurrenceDate{DateTime: &d})
	}
	return nil
}

// parsePeriod parses a PERIOD value as start/end or start/duration. Endpoints
// honor the property's TZID, since RDATE permits a non-UTC period; callers that
// require UTC (FREEBUSY) must check the returned forms themselves.
func parsePeriod(value string, params map[string]string, propertyName string) (model.Period, error) {
	startStr, rest, found := strings.Cut(value, "/")
	if !found || rest == "" {
		return model.Period{}, icalerr.ErrInvalidFreeBusyFormat
	}
	// Both endpoints are DATE-TIME and share the property's TZID, so they are
	// resolved the same way any other DATE-TIME on this property would be.
	startDT, err := parseDateTimeValue(startStr, params, propertyName)
	if err != nil {
		return model.Period{}, err
	}

	if strings.HasPrefix(rest, "P") || strings.HasPrefix(rest, "-P") || strings.HasPrefix(rest, "+P") {
		dur, err := icaldur.ParseICalDuration(rest)
		if err != nil {
			return model.Period{}, err
		}
		if dur <= 0 {
			return model.Period{}, icalerr.ErrPositiveDurationRequired
		}
		return model.Period{Start: startDT, Duration: dur}, nil
	}
	endDT, err := parseDateTimeValue(rest, params, propertyName)
	if err != nil {
		return model.Period{}, err
	}
	// RFC 5545 3.3.9: explicit period start MUST be before end.
	if !endDT.Time.After(startDT.Time) {
		return model.Period{}, icalerr.ErrPeriodEndNotAfterStart
	}
	return model.Period{Start: startDT, End: endDT}, nil
}
