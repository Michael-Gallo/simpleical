package icaldur

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseICalDuration(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		want        time.Duration
		expectError error
	}{
		{name: "hours only", input: "PT1H", want: time.Hour},
		{name: "minutes only", input: "PT1M", want: time.Minute},
		{name: "seconds only", input: "PT1S", want: time.Second},
		{name: "hours and minutes", input: "PT1H30M", want: time.Hour + time.Minute*30},
		{name: "hours minutes and seconds", input: "PT1H30M1S", want: time.Hour + time.Minute*30 + time.Second},
		{name: "days and time", input: "P15DT5H0M20S", want: time.Hour*24*15 + time.Hour*5 + time.Minute*0 + time.Second*20},
		{name: "positive duration prefix", input: "+P15DT5H0M20S", want: time.Hour*24*15 + time.Hour*5 + time.Minute*0 + time.Second*20},
		{name: "negative duration prefix", input: "-P15DT5H0M20S", want: -(time.Hour*24*15 + time.Hour*5 + time.Minute*0 + time.Second*20)},
		{name: "empty input", input: "", want: 0, expectError: errEmpty},
		{name: "bad prefix", input: "+Q15DT5H0M20S", expectError: errBadPrefix},
		{name: "unexpected character", input: "+P15DT5H0M20G", expectError: errUnexpectedChar},
		{name: "missing unit", input: "+P15DT5H0M20", expectError: errMissingUnit},
		{name: "duplicate seconds unit", input: "+P15DT5H0M20S20S", expectError: errDuplicateUnit},
		{name: "duplicate day unit", input: "P1D2D", expectError: errDuplicateUnit},
		{name: "no components after prefix", input: "P", expectError: errNoComponents},
		{name: "no time components after T", input: "PT", expectError: errNoComponents},
		{name: "T without time part", input: "P1DT", expectError: errTWithoutTimePart},
		{name: "duplicate T in time-only duration", input: "PTT1H", expectError: errDuplicateT},
		{name: "duplicate T after date part", input: "P1DTT1H", expectError: errDuplicateT},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			got, err := ParseICalDuration(test.input)
			if test.expectError != nil {
				assert.ErrorIs(t, err, test.expectError)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, test.want, got)
		})
	}
}

func BenchmarkParseICalDuration(b *testing.B) {
	for b.Loop() {
		_, err := ParseICalDuration("P15DT5H0M20S")
		if err != nil {
			b.Fatal(err)
		}
	}
}
