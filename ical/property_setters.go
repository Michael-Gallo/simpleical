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

func timeValueType(propertyName string, params map[string]string) string {
	if !dateValueAllowed(propertyName) {
		return ""
	}
	return params[model.ParamValue]
}

func parseDateTimeValue(value string, params map[string]string, propertyName string) (model.DateTime, error) {
	valueType := timeValueType(propertyName, params)
	temporal, err := icaldur.ParseTemporal(value, valueType)
	if err != nil {
		return model.DateTime{}, err
	}
	tzid := params[model.ParamTZID]
	if tzid != "" && temporal.Form == icaldur.FormUTC {
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
		return fmt.Errorf("%w: %s property %s in iCal", icalerr.ErrParseErrorInComponent, componentType, propertyName)
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

// setOnceDurationProperty sets a duration field only if it hasn't been set before.
func setOnceDurationProperty(field *time.Duration, value, propertyName string, componentType string) error {
	duration, err := icaldur.ParseICalDuration(value)
	if err != nil {
		return fmt.Errorf("%w: %s property %s in iCal", icalerr.ErrParseErrorInComponent, componentType, propertyName)
	}
	return setOnceProperty(field, duration, propertyName, componentType)
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

func appendTimePropertyWithParams(field *[]model.DateTime, value string, params map[string]string, propertyName string, componentType string) error {
	parsed, err := parseDateTimeValue(value, params, propertyName)
	if err != nil {
		return fmt.Errorf("%w: %s property %s in iCal", icalerr.ErrParseErrorInComponent, componentType, propertyName)
	}
	*field = append(*field, parsed)
	return nil
}

func appendCommaSeparatedTimePropertyWithParams(field *[]model.DateTime, value string, params map[string]string, propertyName string, componentType string) error {
	for part := range strings.SplitSeq(value, ",") {
		if err := appendTimePropertyWithParams(field, part, params, propertyName, componentType); err != nil {
			return err
		}
	}
	return nil
}

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

func setOnceTextProperty(field *model.TextValue, value string, params map[string]string, propertyName string, componentType string) error {
	tv, err := parseTextValue(value, params)
	if err != nil {
		return err
	}
	if field.Value != "" {
		return fmt.Errorf(icalerr.ErrDuplicatePropertyInComponentFormat, icalerr.ErrDuplicatePropertyInComponent, propertyName, componentType)
	}
	*field = tv
	return nil
}

func appendTextProperty(field *[]model.TextValue, value string, params map[string]string) error {
	tv, err := parseTextValue(value, params)
	if err != nil {
		return err
	}
	*field = append(*field, tv)
	return nil
}

func appendTextListProperty(field *[]string, value string) error {
	parts, err := splitUnescapedComma(value)
	if err != nil {
		return err
	}
	*field = append(*field, parts...)
	return nil
}

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

func appendRDateProperty(field *[]model.RecurrenceDate, value string, params map[string]string, propertyName string, componentType string) error {
	if strings.EqualFold(params[model.ParamValue], "PERIOD") {
		for part := range strings.SplitSeq(value, ",") {
			period, err := parsePeriod(part)
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
			return fmt.Errorf("%w: %s property %s in iCal", icalerr.ErrParseErrorInComponent, componentType, propertyName)
		}
		d := dt
		*field = append(*field, model.RecurrenceDate{DateTime: &d})
	}
	return nil
}

// parsePeriod parses a UTC PERIOD value as start/end or start/duration.
func parsePeriod(value string) (model.Period, error) {
	startStr, rest, found := strings.Cut(value, "/")
	if !found || rest == "" {
		return model.Period{}, icalerr.ErrInvalidFreeBusyFormat
	}
	start, err := icaldur.ParseTemporalDateTime(startStr)
	if err != nil {
		return model.Period{}, err
	}
	if start.Form != icaldur.FormUTC {
		return model.Period{}, icalerr.ErrUTCValueRequired
	}
	startDT := model.DateTime{Form: model.DateTimeFormUTC, Time: start.Time}

	if strings.HasPrefix(rest, "P") || strings.HasPrefix(rest, "-P") || strings.HasPrefix(rest, "+P") {
		dur, err := icaldur.ParseICalDuration(rest)
		if err != nil {
			return model.Period{}, err
		}
		return model.Period{Start: startDT, Duration: dur}, nil
	}
	end, err := icaldur.ParseTemporalDateTime(rest)
	if err != nil {
		return model.Period{}, err
	}
	if end.Form != icaldur.FormUTC {
		return model.Period{}, icalerr.ErrUTCValueRequired
	}
	return model.Period{
		Start: startDT,
		End:   model.DateTime{Form: model.DateTimeFormUTC, Time: end.Time},
	}, nil
}
