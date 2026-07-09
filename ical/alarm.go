package ical

import (
	"fmt"

	"github.com/michael-gallo/simpleical/internal/icalerr"
	"github.com/michael-gallo/simpleical/model"
)

const alarmLocation = "Alarm"

// parseAlarmProperty parses a single property line and adds it to the provided alarm.
func parseAlarmProperty(propertyName string, value string, params map[string]string, alarm *model.Alarm) error {
	switch model.AlarmToken(propertyName) {
	case model.AlarmTokenAction:
		return setOnceProperty(&alarm.Action, model.AlarmAction(value), propertyName, alarmLocation)
	case model.AlarmTokenTrigger:
		return setOnceProperty(&alarm.Trigger, value, propertyName, alarmLocation)
	case model.AlarmTokenAttach:
		attachment, err := parseAttachment(value, params)
		if err != nil {
			return err
		}
		return setOnceProperty(&alarm.Attach, attachment, propertyName, alarmLocation)
	case model.AlarmTokenDuration:
		return setOnceDurationProperty(&alarm.Duration, value, propertyName, alarmLocation)
	case model.AlarmTokenDescription:
		return setOnceProperty(&alarm.Description, value, propertyName, alarmLocation)
	case model.AlarmTokenRepeat:
		return setOnceIntProperty(&alarm.Repeat, value, propertyName, alarmLocation)
	case model.AlarmTokenSummary:
		return setOnceProperty(&alarm.Summary, value, propertyName, alarmLocation)
	case model.AlarmTokenAttendee:
		attendee, err := parseAttendee(value, params)
		if err != nil {
			return err
		}
		alarm.Attendees = append(alarm.Attendees, *attendee)
	default:
		return fmt.Errorf("%w: %s", icalerr.ErrInvalidAlarmProperty, propertyName)
	}
	return nil
}

// validateAlarm ensures that all required values are present for an alarm.
// An alarm has the following requirements:
//
//   - The "VALARM" calendar component MUST include the "ACTION" and "TRIGGER" properties
//
// This requirement is modified based on the action in the following way.
//   - The "DISPLAY" action MUST include the "DESCRIPTION" property
//   - The "EMAIL" action MUST include the "DESCRIPTION" and "SUMMARY" properties
//   - The "EMAIL" action MUST include at least one "ATTENDEE" property
//   - The "AUDIO" action MUST include a single "ATTACH" property
//   - The "PROCEDURE" action does not have any additional requirements
func validateAlarm(alarm *model.Alarm) error {
	if alarm.Action == "" {
		return icalerr.ErrMissingAlarmActionProperty
	}
	if alarm.Trigger == "" {
		return icalerr.ErrMissingAlarmTriggerProperty
	}

	// Validate action-specific requirements
	switch alarm.Action {
	case model.AlarmActionDisplay:
		if len(alarm.Description) == 0 {
			return icalerr.ErrMissingAlarmDescriptionForDisplay
		}
	case model.AlarmActionEmail:
		if len(alarm.Description) == 0 {
			return icalerr.ErrMissingAlarmDescriptionForEmail
		}
		if alarm.Summary == "" {
			return icalerr.ErrMissingAlarmSummaryForEmail
		}
		if len(alarm.Attendees) == 0 {
			return icalerr.ErrMissingAlarmAttendeesForEmail
		}
	case model.AlarmActionProcedure:
	case model.AlarmActionAudio:
		if alarm.Attach == nil {
			return icalerr.ErrMissingAlarmAttachForAudio
		}
	default:
		return fmt.Errorf("%w: %s", icalerr.ErrUnknownAlarmAction, alarm.Action)
	}

	return nil
}
