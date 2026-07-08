package rrule

import (
	"strconv"
	"strings"
	"testing"
)

func BenchmarkParseByWeekNo(b *testing.B) {
	const value = "1,2,3,4,5,6,7,8,-1,-2,-3,-4,-5,-6,-7,-8"

	b.Run("custom_int8", func(b *testing.B) {
		b.ReportAllocs()
		for b.Loop() {
			if _, err := parseSignedIntListBounded(value, int8(53), errInvalidWeekno); err != nil {
				b.Fatal(err)
			}
		}
	})

	b.Run("split_plus_Atoi", func(b *testing.B) {
		b.ReportAllocs()
		for b.Loop() {
			if err := parseSignedIntListAtoi(value); err != nil {
				b.Fatal(err)
			}
		}
	})
}

func BenchmarkParseBySetPos(b *testing.B) {
	const value = "1,2,3,4,5,100,200,300,366,-1,-2,-3,-4,-100,-200,-300,-366"

	b.Run("custom_int16", func(b *testing.B) {
		b.ReportAllocs()
		for b.Loop() {
			if _, err := parseSignedIntListBounded(value, int16(366), errInvalidBySetPos); err != nil {
				b.Fatal(err)
			}
		}
	})

	b.Run("split_plus_Atoi", func(b *testing.B) {
		b.ReportAllocs()
		for b.Loop() {
			if err := parseSignedIntListAtoi(value); err != nil {
				b.Fatal(err)
			}
		}
	})
}

func parseSignedIntListAtoi(value string) error {
	values := strings.Split(value, ",")
	for _, part := range values {
		if _, err := strconv.Atoi(part); err != nil {
			return err
		}
	}
	return nil
}
