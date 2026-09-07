package ical

import (
	"fmt"
	"strings"

	"github.com/michael-gallo/simpleical/internal/icalerr"
	"github.com/michael-gallo/simpleical/model"
)

// isIANAToken reports whether s matches RFC 5545 iana-token: 1*(ALPHA / DIGIT / "-").
func isIANAToken(s string) bool {
	if s == "" {
		return false
	}
	for i := 0; i < len(s); i++ {
		c := s[i]
		if (c < 'A' || c > 'Z') && (c < 'a' || c > 'z') && (c < '0' || c > '9') && c != '-' {
			return false
		}
	}
	return true
}

// parseClass validates CLASS: known values are canonicalized case-insensitively;
// other valid iana-token / x-name values are retained verbatim.
func parseClass(value string) (model.Class, error) {
	switch strings.ToUpper(value) {
	case string(model.ClassPublic):
		return model.ClassPublic, nil
	case string(model.ClassPrivate):
		return model.ClassPrivate, nil
	case string(model.ClassConfidential):
		return model.ClassConfidential, nil
	}
	if !isIANAToken(value) {
		return "", fmt.Errorf("%w: CLASS %s", icalerr.ErrInvalidEnumValue, value)
	}
	return model.Class(value), nil
}

// parseFreeBusyStatus validates FBTYPE: known values are canonicalized;
// other valid iana-token / x-name values are retained verbatim.
func parseFreeBusyStatus(value string) (model.FreeBusyStatus, error) {
	switch strings.ToUpper(value) {
	case string(model.FreeBusyStatusFree):
		return model.FreeBusyStatusFree, nil
	case string(model.FreeBusyStatusBusy):
		return model.FreeBusyStatusBusy, nil
	case string(model.FreeBusyStatusBusyTentative):
		return model.FreeBusyStatusBusyTentative, nil
	case string(model.FreeBusyStatusBusyUnavailable):
		return model.FreeBusyStatusBusyUnavailable, nil
	}
	if !isIANAToken(value) {
		return "", fmt.Errorf("%w: FBTYPE %s", icalerr.ErrInvalidEnumValue, value)
	}
	return model.FreeBusyStatus(value), nil
}

// parseAlarmAction validates ACTION: known values are canonicalized;
// other valid iana-token / x-name values are retained so END:VALARM can drop them.
func parseAlarmAction(value string) (model.AlarmAction, error) {
	switch strings.ToUpper(value) {
	case string(model.AlarmActionAudio):
		return model.AlarmActionAudio, nil
	case string(model.AlarmActionDisplay):
		return model.AlarmActionDisplay, nil
	case string(model.AlarmActionEmail):
		return model.AlarmActionEmail, nil
	}
	if !isIANAToken(value) {
		return "", fmt.Errorf("%w: ACTION %s", icalerr.ErrInvalidEnumValue, value)
	}
	return model.AlarmAction(value), nil
}

// parseEventStatus validates a VEVENT STATUS property value.
func parseEventStatus(value string) (model.EventStatus, error) {
	u := strings.ToUpper(value)
	switch model.EventStatus(u) {
	case model.EventStatusConfirmed, model.EventStatusTentative, model.EventStatusCancelled:
		return model.EventStatus(u), nil
	default:
		return "", fmt.Errorf("%w: STATUS %s", icalerr.ErrInvalidEnumValue, value)
	}
}

// parseTodoStatus validates a VTODO STATUS property value.
func parseTodoStatus(value string) (model.TodoStatus, error) {
	u := strings.ToUpper(value)
	switch model.TodoStatus(u) {
	case model.TodoStatusNeedsAction, model.TodoStatusCompleted, model.TodoStatusInProcess, model.TodoStatusCancelled:
		return model.TodoStatus(u), nil
	default:
		return "", fmt.Errorf("%w: STATUS %s", icalerr.ErrInvalidEnumValue, value)
	}
}

// parseJournalStatus validates a VJOURNAL STATUS property value.
func parseJournalStatus(value string) (model.JournalStatus, error) {
	u := strings.ToUpper(value)
	switch model.JournalStatus(u) {
	case model.JournalStatusDraft, model.JournalStatusFinal, model.JournalStatusCancelled:
		return model.JournalStatus(u), nil
	default:
		return "", fmt.Errorf("%w: STATUS %s", icalerr.ErrInvalidEnumValue, value)
	}
}

// parseTransp validates a TRANSP property value.
func parseTransp(value string) (model.Transp, error) {
	u := strings.ToUpper(value)
	switch model.Transp(u) {
	case model.TranspTransparent, model.TranspOpaque:
		return model.Transp(u), nil
	default:
		return "", fmt.Errorf("%w: TRANSP %s", icalerr.ErrInvalidEnumValue, value)
	}
}

// isKnownAlarmAction reports whether action is one this parser can validate.
// Alarms with any other action are discarded, since RFC 5545 3.8.6.1 requires
// applications to ignore alarms whose action they do not recognize.
func isKnownAlarmAction(action model.AlarmAction) bool {
	switch action {
	case model.AlarmActionAudio, model.AlarmActionDisplay, model.AlarmActionEmail:
		return true
	default:
		return false
	}
}
