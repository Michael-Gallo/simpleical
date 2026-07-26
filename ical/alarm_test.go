package ical

import (
	"testing"

	"github.com/michael-gallo/simpleical/internal/icalerr"
	"github.com/michael-gallo/simpleical/model"
	"github.com/stretchr/testify/require"
)

func TestIsKnownAlarmAction(t *testing.T) {
	for _, action := range []model.AlarmAction{model.AlarmActionAudio, model.AlarmActionDisplay, model.AlarmActionEmail} {
		require.True(t, isKnownAlarmAction(action))
	}
	require.False(t, isKnownAlarmAction(model.AlarmAction("X-SMS")))
	require.False(t, isKnownAlarmAction(""))
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
