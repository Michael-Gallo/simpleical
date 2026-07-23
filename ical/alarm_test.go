package ical

import (
	"testing"
	"time"

	"github.com/michael-gallo/simpleical/internal/icalerr"
	"github.com/michael-gallo/simpleical/model"
	"github.com/stretchr/testify/require"
)

func TestValidateAlarmRejectsUnknownAction(t *testing.T) {
	dur := -15 * time.Minute
	alarm := &model.Alarm{
		Action:  model.AlarmAction("BOGUS"),
		Trigger: model.Trigger{Duration: &dur, Related: model.TriggerRelatedStart},
	}

	err := validateAlarm(alarm)

	require.ErrorIs(t, err, icalerr.ErrUnknownAlarmAction)
	require.ErrorContains(t, err, "BOGUS")
}

func TestParseTriggerRejectsRelatedOnAbsolute(t *testing.T) {
	_, err := parseTrigger("19980101T050000Z", map[string]string{
		model.ParamRelated: string(model.TriggerRelatedEnd),
	})
	require.ErrorIs(t, err, icalerr.ErrInvalidAlarmTrigger)
}

func TestParseTriggerAllowsRelatedEndOnDuration(t *testing.T) {
	trig, err := parseTrigger("-PT15M", map[string]string{
		model.ParamRelated: string(model.TriggerRelatedEnd),
	})
	require.NoError(t, err)
	require.NotNil(t, trig.Duration)
	require.Equal(t, model.TriggerRelatedEnd, trig.Related)
}
