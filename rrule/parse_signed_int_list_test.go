package rrule

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseSignedIntListBounded(t *testing.T) {
	t.Run("int8", func(t *testing.T) {
		tests := []struct {
			name        string
			input       string
			want        []int8
			expectError error
		}{
			{
				name:  "single positive value",
				input: "20",
				want:  []int8{20},
			},
			{
				name:  "explicit positive sign",
				input: "+1",
				want:  []int8{1},
			},
			{
				name:  "multiple values with explicit positive sign",
				input: "+1,-2,+53",
				want:  []int8{1, -2, 53},
			},
			{
				name:  "single negative value",
				input: "-2",
				want:  []int8{-2},
			},
			{
				name:  "multiple comma-separated values",
				input: "1,-2,53",
				want:  []int8{1, -2, 53},
			},
			{
				name:  "minimum allowed negative value",
				input: "-53",
				want:  []int8{-53},
			},
			{
				name:        "zero is rejected",
				input:       "0",
				expectError: errInvalidWeekno,
			},
			{
				name:        "zero in list is rejected",
				input:       "1,0,2",
				expectError: errInvalidWeekno,
			},
			{
				name:        "above max is rejected",
				input:       "54",
				expectError: errInvalidWeekno,
			},
			{
				name:        "overflowing int8 magnitude is rejected",
				input:       "130",
				expectError: errInvalidWeekno,
			},
			{
				name:        "below min is rejected",
				input:       "-54",
				expectError: errInvalidWeekno,
			},
			{
				name:        "non-numeric token is rejected",
				input:       "abc",
				expectError: strconv.ErrSyntax,
			},
			{
				name:        "empty input is rejected",
				input:       "",
				expectError: strconv.ErrSyntax,
			},
			{
				name:        "empty list element is rejected",
				input:       "1,,2",
				expectError: strconv.ErrSyntax,
			},
			{
				name:        "bare minus is rejected",
				input:       "-",
				expectError: strconv.ErrSyntax,
			},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				got, err := parseSignedIntListBounded(test.input, int8(53), errInvalidWeekno)
				if test.expectError != nil {
					require.Error(t, err)
					assert.ErrorContains(t, err, test.expectError.Error())
					assert.Nil(t, got)
					return
				}

				require.NoError(t, err)
				assert.Equal(t, test.want, got)
			})
		}
	})

	t.Run("int16", func(t *testing.T) {
		tests := []struct {
			name        string
			input       string
			want        []int16
			expectError error
		}{
			{
				name:  "single positive value",
				input: "3",
				want:  []int16{3},
			},
			{
				name:  "explicit positive sign",
				input: "+366",
				want:  []int16{366},
			},
			{
				name:  "multiple comma-separated values",
				input: "1,-1,366",
				want:  []int16{1, -1, 366},
			},
			{
				name:  "minimum allowed negative value",
				input: "-366",
				want:  []int16{-366},
			},
			{
				name:        "zero is rejected",
				input:       "0",
				expectError: errInvalidBySetPos,
			},
			{
				name:        "above max is rejected",
				input:       "367",
				expectError: errInvalidBySetPos,
			},
			{
				name:        "below min is rejected",
				input:       "-367",
				expectError: errInvalidBySetPos,
			},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				got, err := parseSignedIntListBounded(test.input, int16(366), errInvalidBySetPos)
				if test.expectError != nil {
					require.Error(t, err)
					assert.ErrorContains(t, err, test.expectError.Error())
					assert.Nil(t, got)
					return
				}

				require.NoError(t, err)
				assert.Equal(t, test.want, got)
			})
		}
	})
}
