package ical

import (
	"testing"
	"time"

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

	var start time.Time
	require.NoError(t, setOnceTimePropertyWithParams(&start, "20070628", params, "DTSTART", "Event"))
	assert.Equal(t, time.Date(2007, 6, 28, 0, 0, 0, 0, time.UTC), start)

	var stamp time.Time
	err := setOnceTimePropertyWithParams(&stamp, "20070628", params, "DTSTAMP", "Event")
	require.Error(t, err)
	assert.True(t, stamp.IsZero())
}
