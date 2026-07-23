package ical

import (
	"fmt"

	"github.com/michael-gallo/simpleical/internal/icalerr"
	"github.com/michael-gallo/simpleical/model"
)

// parseClass validates a CLASS property value.
func parseClass(value string) (model.Class, error) {
	switch model.Class(value) {
	case model.ClassPublic, model.ClassPrivate, model.ClassConfidential:
		return model.Class(value), nil
	default:
		return "", fmt.Errorf("%w: CLASS %s", icalerr.ErrInvalidEnumValue, value)
	}
}

// parseEventStatus validates a VEVENT STATUS property value.
func parseEventStatus(value string) (model.EventStatus, error) {
	switch model.EventStatus(value) {
	case model.EventStatusConfirmed, model.EventStatusTentative, model.EventStatusCancelled:
		return model.EventStatus(value), nil
	default:
		return "", fmt.Errorf("%w: STATUS %s", icalerr.ErrInvalidEnumValue, value)
	}
}

// parseTodoStatus validates a VTODO STATUS property value.
func parseTodoStatus(value string) (model.TodoStatus, error) {
	switch model.TodoStatus(value) {
	case model.TodoStatusNeedsAction, model.TodoStatusCompleted, model.TodoStatusInProcess, model.TodoStatusCancelled:
		return model.TodoStatus(value), nil
	default:
		return "", fmt.Errorf("%w: STATUS %s", icalerr.ErrInvalidEnumValue, value)
	}
}

// parseJournalStatus validates a VJOURNAL STATUS property value.
func parseJournalStatus(value string) (model.JournalStatus, error) {
	switch model.JournalStatus(value) {
	case model.JournalStatusDraft, model.JournalStatusFinal, model.JournalStatusCancelled:
		return model.JournalStatus(value), nil
	default:
		return "", fmt.Errorf("%w: STATUS %s", icalerr.ErrInvalidEnumValue, value)
	}
}

// parseTransp validates a TRANSP property value.
func parseTransp(value string) (model.Transp, error) {
	switch model.Transp(value) {
	case model.TranspTransparent, model.TranspOpaque:
		return model.Transp(value), nil
	default:
		return "", fmt.Errorf("%w: TRANSP %s", icalerr.ErrInvalidEnumValue, value)
	}
}

// parseFreeBusyStatus validates an FBTYPE parameter value.
func parseFreeBusyStatus(value string) (model.FreeBusyStatus, error) {
	switch model.FreeBusyStatus(value) {
	case model.FreeBusyStatusFree, model.FreeBusyStatusBusy, model.FreeBusyStatusBusyTentative, model.FreeBusyStatusBusyUnavailable:
		return model.FreeBusyStatus(value), nil
	default:
		return "", fmt.Errorf("%w: FBTYPE %s", icalerr.ErrInvalidEnumValue, value)
	}
}

// parseAlarmAction validates a VALARM ACTION property value.
func parseAlarmAction(value string) (model.AlarmAction, error) {
	switch model.AlarmAction(value) {
	case model.AlarmActionAudio, model.AlarmActionDisplay, model.AlarmActionEmail:
		return model.AlarmAction(value), nil
	default:
		return "", fmt.Errorf("%w: %s", icalerr.ErrUnknownAlarmAction, value)
	}
}
