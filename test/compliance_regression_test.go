package test

import (
	_ "embed"
	"strings"
	"testing"

	"github.com/michael-gallo/simpleical/ical"
	"github.com/michael-gallo/simpleical/icaldur"
	"github.com/michael-gallo/simpleical/internal/icalerr"
	"github.com/michael-gallo/simpleical/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Spec regressions locked in while aligning the parser with RFC 5545.
// Compatibility tolerances (trailing whitespace on BEGIN/END, late calendar
// properties) live in calendar_test.go instead.

var (
	//go:embed test_data/spec_gaps/valid_class_x_name.ical
	testGapClassXNameInput string
	//go:embed test_data/spec_gaps/valid_alarm_unknown_action.ical
	testGapAlarmUnknownActionInput string
	//go:embed test_data/spec_gaps/valid_freebusy_fbtype_x_name.ical
	testGapFreeBusyFBTypeXNameInput string
	//go:embed test_data/spec_gaps/valid_rdate_period_tzid.ical
	testGapRDatePeriodTZIDInput string
	//go:embed test_data/spec_gaps/invalid_global_tzid_without_vtimezone.ical
	testGapGlobalTZIDWithoutVTimezoneInput string
	//go:embed test_data/spec_gaps/valid_alarm_repeat_zero_with_duration.ical
	testGapAlarmRepeatZeroWithDurationInput string

	//go:embed test_data/spec_gaps/invalid_display_alarm_with_attach.ical
	testGapDisplayAlarmWithAttachInput string
	//go:embed test_data/spec_gaps/invalid_tzid_on_date_value.ical
	testGapTZIDOnDateValueInput string
)

func TestClassAcceptsXName(t *testing.T) {
	cal, err := ical.ReadSingle(strings.NewReader(testGapClassXNameInput))
	require.NoError(t, err)
	assert.Equal(t, model.Class("X-CORPORATE"), cal.Events[0].Class)
}

func TestUnrecognizedAlarmActionDropsAlarm(t *testing.T) {
	cal, err := ical.ReadSingle(strings.NewReader(testGapAlarmUnknownActionInput))
	require.NoError(t, err)
	require.Len(t, cal.Events, 1)
	assert.Empty(t, cal.Events[0].Alarms, "the alarm is ignored but the event it sits in still parses")
}

func TestFreeBusyTypeAcceptsXName(t *testing.T) {
	cal, err := ical.ReadSingle(strings.NewReader(testGapFreeBusyFBTypeXNameInput))
	require.NoError(t, err)
	require.Len(t, cal.FreeBusys[0].FreeBusy, 1)
	assert.Equal(t, model.FreeBusyStatus("X-MOVIE-NIGHT"), cal.FreeBusys[0].FreeBusy[0].FBType)
}

func TestRDatePeriodAcceptsTZID(t *testing.T) {
	cal, err := ical.ReadSingle(strings.NewReader(testGapRDatePeriodTZIDInput))
	require.NoError(t, err)
	require.Len(t, cal.Events[0].Rdate, 1)
	period := cal.Events[0].Rdate[0].Period
	require.NotNil(t, period)
	assert.Equal(t, model.DateTimeFormLocalTZ, period.Start.Form)
	assert.Equal(t, "America/New_York", period.Start.TZID)
	assert.Equal(t, "America/New_York", period.End.TZID)
}

func TestGlobalTZIDRequiresVTimezone(t *testing.T) {
	_, err := ical.ReadSingle(strings.NewReader(testGapGlobalTZIDWithoutVTimezoneInput))
	require.ErrorIs(t, err, icalerr.ErrUnknownTZID)
}

func TestAlarmRepeatZeroWithDuration(t *testing.T) {
	cal, err := ical.ReadSingle(strings.NewReader(testGapAlarmRepeatZeroWithDurationInput))
	require.NoError(t, err)
	require.Len(t, cal.Events[0].Alarms, 1)
	require.NotNil(t, cal.Events[0].Alarms[0].Repeat)
	assert.Equal(t, 0, *cal.Events[0].Alarms[0].Repeat)
}

func TestAttachRejectedOnDisplayAlarm(t *testing.T) {
	_, err := ical.ReadSingle(strings.NewReader(testGapDisplayAlarmWithAttachInput))
	require.ErrorIs(t, err, icalerr.ErrAlarmAttachNotAllowedForDisplay)
}

func TestTZIDRejectedOnDateValue(t *testing.T) {
	_, err := ical.ReadSingle(strings.NewReader(testGapTZIDOnDateValueInput))
	require.ErrorIs(t, err, icaldur.ErrTZIDWithUTC)
}
