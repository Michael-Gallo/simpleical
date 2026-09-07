package ical

import (
	"testing"

	"github.com/michael-gallo/simpleical/internal/icalerr"
	"github.com/michael-gallo/simpleical/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseClassCanonicalizesAndAcceptsExtensions(t *testing.T) {
	t.Parallel()
	got, err := parseClass("public")
	require.NoError(t, err)
	assert.Equal(t, model.ClassPublic, got)

	got, err = parseClass("X-CORPORATE")
	require.NoError(t, err)
	assert.Equal(t, model.Class("X-CORPORATE"), got)

	got, err = parseClass("CUSTOM-CLASS")
	require.NoError(t, err)
	assert.Equal(t, model.Class("CUSTOM-CLASS"), got)

	_, err = parseClass("")
	require.ErrorIs(t, err, icalerr.ErrInvalidEnumValue)
	_, err = parseClass("not valid")
	require.ErrorIs(t, err, icalerr.ErrInvalidEnumValue)
}

func TestParseAlarmActionCanonicalizesAndRejectsMalformed(t *testing.T) {
	t.Parallel()
	got, err := parseAlarmAction("display")
	require.NoError(t, err)
	assert.Equal(t, model.AlarmActionDisplay, got)

	got, err = parseAlarmAction("X-SMS")
	require.NoError(t, err)
	assert.Equal(t, model.AlarmAction("X-SMS"), got)

	_, err = parseAlarmAction("")
	require.ErrorIs(t, err, icalerr.ErrInvalidEnumValue)
	_, err = parseAlarmAction("???")
	require.ErrorIs(t, err, icalerr.ErrInvalidEnumValue)
}

func TestParseFreeBusyStatusCanonicalizes(t *testing.T) {
	t.Parallel()
	got, err := parseFreeBusyStatus("busy-tentative")
	require.NoError(t, err)
	assert.Equal(t, model.FreeBusyStatusBusyTentative, got)

	got, err = parseFreeBusyStatus("X-MOVIE-NIGHT")
	require.NoError(t, err)
	assert.Equal(t, model.FreeBusyStatus("X-MOVIE-NIGHT"), got)

	_, err = parseFreeBusyStatus("BUSY NOW")
	require.ErrorIs(t, err, icalerr.ErrInvalidEnumValue)
}

func TestClosedEnumsAreCaseInsensitive(t *testing.T) {
	t.Parallel()
	status, err := parseEventStatus("confirmed")
	require.NoError(t, err)
	assert.Equal(t, model.EventStatusConfirmed, status)

	transp, err := parseTransp("transparent")
	require.NoError(t, err)
	assert.Equal(t, model.TranspTransparent, transp)

	todoStatus, err := parseTodoStatus("needs-action")
	require.NoError(t, err)
	assert.Equal(t, model.TodoStatusNeedsAction, todoStatus)

	journalStatus, err := parseJournalStatus("draft")
	require.NoError(t, err)
	assert.Equal(t, model.JournalStatusDraft, journalStatus)
}
