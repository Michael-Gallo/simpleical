package ical

import (
	"testing"
	"time"

	"github.com/michael-gallo/simpleical/internal/icalerr"
	"github.com/michael-gallo/simpleical/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDateValueAllowed(t *testing.T) {
	t.Parallel()
	for _, name := range []string{"DTSTART", "DTEND", "DUE", "RECURRENCE-ID", "EXDATE", "RDATE"} {
		assert.True(t, dateValueAllowed(name), name)
	}
	for _, name := range []string{"DTSTAMP", "CREATED", "LAST-MODIFIED", "COMPLETED"} {
		assert.False(t, dateValueAllowed(name), name)
	}
}

func TestSetOnceTimePropertyWithParamsDateGating(t *testing.T) {
	t.Parallel()
	params := map[string]string{"VALUE": "DATE"}

	var start model.DateTime
	require.NoError(t, setOnceTimePropertyWithParams(&start, "20070628", params, "DTSTART", "Event"))
	assert.Equal(t, model.NewDate(time.Date(2007, 6, 28, 0, 0, 0, 0, time.UTC)), start)

	var stamp model.DateTime
	err := setOnceTimePropertyWithParams(&stamp, "20070628", params, "DTSTAMP", "Event")
	require.Error(t, err)
	assert.True(t, stamp.IsZero())
}

func TestAppendCommaSeparatedTimePropertyWithParams(t *testing.T) {
	t.Parallel()
	var dates []model.DateTime
	require.NoError(t, appendCommaSeparatedTimePropertyWithParams(
		&dates,
		"19960402T010000Z,19960403T010000Z,19960404T010000Z",
		nil,
		"EXDATE",
		"Event",
	))
	assert.Equal(t, []model.DateTime{
		model.NewUTCDateTime(time.Date(1996, 4, 2, 1, 0, 0, 0, time.UTC)),
		model.NewUTCDateTime(time.Date(1996, 4, 3, 1, 0, 0, 0, time.UTC)),
		model.NewUTCDateTime(time.Date(1996, 4, 4, 1, 0, 0, 0, time.UTC)),
	}, dates)
}

func TestParseRecurrenceIDRange(t *testing.T) {
	t.Parallel()
	got, err := parseRecurrenceIDRange(nil)
	require.NoError(t, err)
	assert.Empty(t, got)

	got, err = parseRecurrenceIDRange(map[string]string{model.ParamRange: "THISANDFUTURE"})
	require.NoError(t, err)
	assert.Equal(t, "THISANDFUTURE", got)

	_, err = parseRecurrenceIDRange(map[string]string{model.ParamRange: "THISANDPRIOR"})
	require.ErrorIs(t, err, icalerr.ErrInvalidEnumValue)
}

func TestParsePeriodRejectsNonPositiveDuration(t *testing.T) {
	t.Parallel()
	_, err := parsePeriod("19970101T180000Z/PT0S")
	require.ErrorIs(t, err, icalerr.ErrPositiveDurationRequired)

	_, err = parsePeriod("19970101T180000Z/-PT1H")
	require.ErrorIs(t, err, icalerr.ErrPositiveDurationRequired)

	p, err := parsePeriod("19970101T180000Z/PT1H")
	require.NoError(t, err)
	assert.Equal(t, time.Hour, p.Duration)
}

func TestValidateFreeBusyRequiresUTCBoundaries(t *testing.T) {
	t.Parallel()
	fb := &model.FreeBusy{
		UID:     "fb@example.com",
		DTStamp: model.NewUTCDateTime(time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)),
		DTStart: model.NewFloatingDateTime(time.Date(2024, 1, 1, 9, 0, 0, 0, time.UTC)),
		DTEnd:   model.NewUTCDateTime(time.Date(2024, 1, 1, 17, 0, 0, 0, time.UTC)),
	}
	require.ErrorIs(t, validateFreeBusy(fb), icalerr.ErrUTCValueRequired)

	fb.DTStart = model.NewUTCDateTime(time.Date(2024, 1, 1, 9, 0, 0, 0, time.UTC))
	require.NoError(t, validateFreeBusy(fb))
}
