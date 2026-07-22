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
