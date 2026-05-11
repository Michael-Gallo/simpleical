package ical

import (
	"testing"

	"github.com/michael-gallo/simpleical/internal/icalerr"
	"github.com/michael-gallo/simpleical/model"
	"github.com/stretchr/testify/require"
)

func TestValidateAlarmRejectsUnknownAction(t *testing.T) {
	alarm := &model.Alarm{
		Action:  model.AlarmAction("BOGUS"),
		Trigger: "-PT15M",
	}

	err := validateAlarm(alarm)

	require.ErrorIs(t, err, icalerr.ErrUnknownAlarmAction)
	require.ErrorContains(t, err, "BOGUS")
}
