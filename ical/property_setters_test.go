package ical

import (
	"testing"
	"time"

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
