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
		return setOncePositiveDurationProperty(&alarm.Duration, value, propertyName, alarmLocation)
	case model.PropDescription:
		unescaped, err := unescapeText(value)
		if err != nil {
			return err
		}
		return setOnceProperty(&alarm.Description, unescaped, propertyName, alarmLocation)
	case model.PropRepeat:
		return setOnceIntPtrProperty(&alarm.Repeat, value, propertyName, alarmLocation)
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
		if isRFCDefinedProperty(propertyName) {
			return fmt.Errorf("%w: %s", icalerr.ErrPropertyNotAllowedInAlarm, propertyName)
		}
		appendExtensionProperty(&alarm.XProp, &alarm.IANAProp, propertyName, value, params)
	}
	return nil
}

// isRFCDefinedProperty reports whether name is a property defined by RFC 5545.
// Such names must not be accepted as iana-prop inside VALARM unless the alarm
// productions list them.
func isRFCDefinedProperty(name string) bool {
	switch name {
	case model.PropAction,
		model.PropAttach,
		model.PropAttendee,
		model.PropCategories,
		model.PropClass,
		model.PropComment,
		model.PropCompleted,
		model.PropContact,
		model.PropCreated,
		model.PropDescription,
		model.PropDTEnd,
		model.PropDTStamp,
		model.PropDTStart,
		model.PropDue,
		model.PropDuration,
		model.PropExDate,
		model.PropFreeBusy,
		model.PropGeo,
		model.PropLastModified,
		model.PropLocation,
		model.PropOrganizer,
		model.PropPercentComplete,
		model.PropPriority,
		model.PropRDate,
		model.PropRecurrenceID,
		model.PropRelatedTo,
		model.PropRepeat,
		model.PropRequestStatus,
		model.PropResources,
		model.PropRRule,
		model.PropSequence,
		model.PropStatus,
		model.PropSummary,
		model.PropTransp,
		model.PropTrigger,
		model.PropTZID,
		model.PropTZName,
		model.PropTZOffsetFrom,
		model.PropTZOffsetTo,
		model.PropTZURL,
		model.PropUID,
		model.PropURL:
		return true
	default:
		return false
	}
}

// parseTrigger parses a TRIGGER as absolute UTC DATE-TIME or relative DURATION.
// RELATED is only valid for duration triggers; absolute triggers reject non-default RELATED.
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

// triggerLooksLikeDateTime reports whether value looks like a DATE-TIME rather than a DURATION.
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
// Callers must have established that the action is one of AUDIO, DISPLAY, or
// EMAIL; alarms with any other action are discarded before reaching here.
//
//   - ACTION and TRIGGER are always required
//   - DURATION and REPEAT must both be present or both absent
//   - DISPLAY requires DESCRIPTION; forbids ATTACH, ATTENDEE, SUMMARY
//   - EMAIL requires DESCRIPTION, SUMMARY, and at least one ATTENDEE
//   - AUDIO may include at most one ATTACH; forbids DESCRIPTION, SUMMARY, ATTENDEE
func validateAlarm(alarm *model.Alarm) error {
	if alarm.Action == "" {
		return icalerr.ErrMissingAlarmActionProperty
	}
	if alarm.Trigger.Duration == nil && alarm.Trigger.Absolute == nil {
		return icalerr.ErrMissingAlarmTriggerProperty
	}

	hasDuration := alarm.Duration != 0
	hasRepeat := alarm.Repeat != nil
	if hasDuration != hasRepeat {
		return icalerr.ErrAlarmDurationRepeatCoupling
	}

	switch alarm.Action {
	case model.AlarmActionDisplay:
		if alarm.Description == "" {
			return icalerr.ErrMissingAlarmDescriptionForDisplay
		}
		if len(alarm.Attach) > 0 {
			return icalerr.ErrAlarmAttachNotAllowedForDisplay
		}
		if len(alarm.Attendees) > 0 {
			return fmt.Errorf("%w: ATTENDEE", icalerr.ErrAlarmPropertyNotAllowed)
		}
		if alarm.Summary != "" {
			return fmt.Errorf("%w: SUMMARY", icalerr.ErrAlarmPropertyNotAllowed)
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
		if alarm.Description != "" {
			return fmt.Errorf("%w: DESCRIPTION", icalerr.ErrAlarmPropertyNotAllowed)
		}
		if alarm.Summary != "" {
			return fmt.Errorf("%w: SUMMARY", icalerr.ErrAlarmPropertyNotAllowed)
		}
		if len(alarm.Attendees) > 0 {
			return fmt.Errorf("%w: ATTENDEE", icalerr.ErrAlarmPropertyNotAllowed)
		}
	}

	return nil
}
