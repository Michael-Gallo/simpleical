package ical

import (
	"fmt"

	"github.com/michael-gallo/simpleical/internal/icalerr"
	"github.com/michael-gallo/simpleical/model"
)

// parseEnum accepts a string value when it matches one of the allowed constants.
// label is included in the error (e.g. "CLASS"); pass "" to omit it (ALARM ACTION).
func parseEnum[T ~string](value string, invalid error, label string, allowed ...T) (T, error) {
	v := T(value)
	for _, a := range allowed {
		if v == a {
			return v, nil
		}
	}
	if label == "" {
		return "", fmt.Errorf("%w: %s", invalid, value)
	}
	return "", fmt.Errorf("%w: %s %s", invalid, label, value)
}

func parseClass(value string) (model.Class, error) {
	return parseEnum(value, icalerr.ErrInvalidEnumValue, "CLASS",
		model.ClassPublic, model.ClassPrivate, model.ClassConfidential)
}

func parseEventStatus(value string) (model.EventStatus, error) {
	return parseEnum(value, icalerr.ErrInvalidEnumValue, "STATUS",
		model.EventStatusConfirmed, model.EventStatusTentative, model.EventStatusCancelled)
}

func parseTodoStatus(value string) (model.TodoStatus, error) {
	return parseEnum(value, icalerr.ErrInvalidEnumValue, "STATUS",
		model.TodoStatusNeedsAction, model.TodoStatusCompleted, model.TodoStatusInProcess, model.TodoStatusCancelled)
}

func parseJournalStatus(value string) (model.JournalStatus, error) {
	return parseEnum(value, icalerr.ErrInvalidEnumValue, "STATUS",
		model.JournalStatusDraft, model.JournalStatusFinal, model.JournalStatusCancelled)
}

func parseTransp(value string) (model.Transp, error) {
	return parseEnum(value, icalerr.ErrInvalidEnumValue, "TRANSP",
		model.TranspTransparent, model.TranspOpaque)
}

func parseFreeBusyStatus(value string) (model.FreeBusyStatus, error) {
	return parseEnum(value, icalerr.ErrInvalidEnumValue, "FBTYPE",
		model.FreeBusyStatusFree, model.FreeBusyStatusBusy, model.FreeBusyStatusBusyTentative, model.FreeBusyStatusBusyUnavailable)
}

func parseAlarmAction(value string) (model.AlarmAction, error) {
	return parseEnum(value, icalerr.ErrUnknownAlarmAction, "",
		model.AlarmActionAudio, model.AlarmActionDisplay, model.AlarmActionEmail)
}
