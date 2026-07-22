package ical

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUnescapeText(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    string
		wantErr bool
	}{
		{name: "newline", input: `Line1\nLine2`, want: "Line1\nLine2"},
		{name: "uppercase N", input: `Line1\NLine2`, want: "Line1\nLine2"},
		{name: "comma", input: `A\,B`, want: "A,B"},
		{name: "semicolon", input: `A\;B`, want: "A;B"},
		{name: "backslash", input: `A\\B`, want: `A\B`},
		{name: "trailing escape", input: `A\`, wantErr: true},
		{name: "invalid escape", input: `\x`, wantErr: true},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := unescapeText(tc.input)
			if tc.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestSplitUnescapedComma(t *testing.T) {
	parts, err := splitUnescapedComma(`A\,B,C`)
	require.NoError(t, err)
	assert.Equal(t, []string{"A,B", "C"}, parts)
}
