package ical

import (
	"fmt"
	"strings"

	"github.com/michael-gallo/simpleical/icaldur"
	"github.com/michael-gallo/simpleical/internal/icalerr"
	"github.com/michael-gallo/simpleical/model"
)

const alarmLocation = "Alarm"

// parseAlarmProperty parses a single property line and adds it to the provided alarm.
func parseAlarmProperty(propertyName string, value string, params map[string]string, alarm *model.Alarm) error {
	switch propertyName {
	case model.PropAction:
		action, err := parseAlarmAction(value)
		if err != nil {
			return err
		}
		return setOnceProperty(&alarm.Action, action, propertyName, alarmLocation)
	case model.PropTrigger:
		if alarm.Trigger.Duration != nil || alarm.Trigger.Absolute != nil {
			return fmt.Errorf(icalerr.ErrDuplicatePropertyInComponentFormat, icalerr.ErrDuplicatePropertyInComponent, propertyName, alarmLocation)
		}
		trigger, err := parseTrigger(value, params)
		if err != nil {
			return err
		}
		alarm.Trigger = trigger
		return nil
	case model.PropAttach:
		attachment, err := parseAttachment(value, params)
		if err != nil {
			return err
		}
		alarm.Attach = append(alarm.Attach, *attachment)
		return nil
	case model.PropDuration:
		return setOnceDurationProperty(&alarm.Duration, value, propertyName, alarmLocation)
	case model.PropDescription:
		unescaped, err := unescapeText(value)
		if err != nil {
			return err
		}
		return setOnceProperty(&alarm.Description, unescaped, propertyName, alarmLocation)
	case model.PropRepeat:
		return setOnceIntProperty(&alarm.Repeat, value, propertyName, alarmLocation)
	case model.PropSummary:
		unescaped, err := unescapeText(value)
		if err != nil {
			return err
		}
		return setOnceProperty(&alarm.Summary, unescaped, propertyName, alarmLocation)
	case model.PropAttendee:
		attendee, err := parseAttendee(value, params)
		if err != nil {
			return err
		}
		alarm.Attendees = append(alarm.Attendees, *attendee)
	default:
		appendExtensionProperty(&alarm.XProp, &alarm.IANAProp, propertyName, value, params)
	}
	return nil
}

func parseTrigger(value string, params map[string]string) (model.Trigger, error) {
	related := model.TriggerRelatedStart
	if raw := params[model.ParamRelated]; raw != "" {
		switch model.TriggerRelated(strings.ToUpper(raw)) {
		case model.TriggerRelatedStart:
			related = model.TriggerRelatedStart
		case model.TriggerRelatedEnd:
			related = model.TriggerRelatedEnd
		default:
			return model.Trigger{}, icalerr.ErrInvalidAlarmTrigger
		}
	}

	valueType := strings.ToUpper(params[model.ParamValue])
	asDateTime := valueType == "DATE-TIME" || (valueType != "DURATION" && triggerLooksLikeDateTime(value))
	if asDateTime {
		// RELATED is only valid for duration triggers (default START).
		if related != model.TriggerRelatedStart {
			return model.Trigger{}, icalerr.ErrInvalidAlarmTrigger
		}
		dt, err := parseUTCDateTimeValue(value, params, model.PropTrigger)
		if err != nil {
			return model.Trigger{}, fmt.Errorf("%w: %w", icalerr.ErrInvalidAlarmTrigger, err)
		}
		return model.Trigger{Absolute: &dt, Related: related}, nil
	}

	dur, err := icaldur.ParseICalDuration(value)
	if err != nil {
		return model.Trigger{}, fmt.Errorf("%w: %w", icalerr.ErrInvalidAlarmTrigger, err)
	}
	return model.Trigger{Duration: &dur, Related: related}, nil
}

func triggerLooksLikeDateTime(value string) bool {
	if value == "" {
		return false
	}
	s := value
	if s[0] == '+' || s[0] == '-' {
		if len(s) > 1 && (s[1] == 'P' || s[1] == 'p') {
			return false
		}
	}
	if s[0] == 'P' || s[0] == 'p' {
		return false
	}
	return true
}

// validateAlarm ensures that all required values are present for an alarm.
//
//   - ACTION and TRIGGER are always required
//   - DISPLAY requires DESCRIPTION
//   - EMAIL requires DESCRIPTION, SUMMARY, and at least one ATTENDEE
//   - AUDIO may include at most one ATTACH
//   - DURATION and REPEAT must both be present or both absent
func validateAlarm(alarm *model.Alarm) error {
	if alarm.Action == "" {
		return icalerr.ErrMissingAlarmActionProperty
	}
	if alarm.Trigger.Duration == nil && alarm.Trigger.Absolute == nil {
		return icalerr.ErrMissingAlarmTriggerProperty
	}

	hasDuration := alarm.Duration != 0
	hasRepeat := alarm.Repeat != 0
	if hasDuration != hasRepeat {
		return icalerr.ErrAlarmDurationRepeatCoupling
	}

	switch alarm.Action {
	case model.AlarmActionDisplay:
		if alarm.Description == "" {
			return icalerr.ErrMissingAlarmDescriptionForDisplay
		}
	case model.AlarmActionEmail:
		if alarm.Description == "" {
			return icalerr.ErrMissingAlarmDescriptionForEmail
		}
		if alarm.Summary == "" {
			return icalerr.ErrMissingAlarmSummaryForEmail
		}
		if len(alarm.Attendees) == 0 {
			return icalerr.ErrMissingAlarmAttendeesForEmail
		}
	case model.AlarmActionAudio:
		if len(alarm.Attach) > 1 {
			return icalerr.ErrAlarmAttachTooManyForAudio
		}
	default:
		return fmt.Errorf("%w: %s", icalerr.ErrUnknownAlarmAction, alarm.Action)
	}

	return nil
}
