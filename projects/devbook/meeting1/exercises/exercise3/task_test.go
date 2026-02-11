package exercise3

import (
	"fmt"
	"testing"
)

func TestArrStats(t *testing.T) {
	type testCase struct {
		input   []int
		wantMin int
		wantMax int
		wantAvg float64
	}
	cases := []testCase{
		{input: []int{}, wantMin: 0, wantMax: 0, wantAvg: 0},
		{input: []int{1, 2, 3}, wantMin: 1, wantMax: 3, wantAvg: 2},
		{input: []int{3, 2, 1}, wantMin: 1, wantMax: 3, wantAvg: 2},
		{input: []int{5, 5, 5, 5}, wantMin: 5, wantMax: 5, wantAvg: 5},
		{input: []int{0, 1}, wantMin: 0, wantMax: 1, wantAvg: 0.5},
	}
	for _, tc := range cases {
		t.Run(fmt.Sprintf("ArrStats(\"%#v\")", tc.input), func(t *testing.T) {
			gotMin, gotMax, gotAvg := ArrStats(tc.input)
			if gotMin != tc.wantMin || gotMax != tc.wantMax || gotAvg != tc.wantAvg {
				t.Errorf(
					"ArrStats(\"%#v\") = (%d, %d, %f), want (%d, %d, %f)",
					tc.input,
					gotMin, gotMax, gotAvg,
					tc.wantMin, tc.wantMax, tc.wantAvg,
				)
			}
		})

	}
}
