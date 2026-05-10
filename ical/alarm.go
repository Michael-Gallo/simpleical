package ical

import (
	"fmt"
	"net/url"

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
		alarm.Attach = append(alarm.Attach, *attachment)
		return nil
	case model.AlarmTokenDuration:
		return setOnceDurationProperty(&alarm.Duration, value, propertyName, alarmLocation)
	case model.AlarmTokenDescription:
		return setOnceProperty(&alarm.Description, value, propertyName, alarmLocation)
	case model.AlarmTokenRepeat:
		return setOnceIntProperty(&alarm.Repeat, value, propertyName, alarmLocation)
	case model.AlarmTokenSummary:
		return setOnceProperty(&alarm.Summary, value, propertyName, alarmLocation)
	case model.AlarmTokenAttendee:
		parsedURL, err := url.Parse(value)
		if err != nil {
			return err
		}
		alarm.Attendees = append(alarm.Attendees, *parsedURL)
	default:
		return fmt.Errorf("%w: %s", icalerr.ErrInvalidAlarmProperty, propertyName)
	}
	return nil
}

// validateAlarm ensures that all required values are present for an alarm.
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
	default:
		return fmt.Errorf("%w: %s", icalerr.ErrUnknownAlarmAction, alarm.Action)
	}

	return nil
}
