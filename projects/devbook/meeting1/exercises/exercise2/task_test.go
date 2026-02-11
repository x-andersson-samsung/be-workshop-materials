package exercise2

import (
	"fmt"
	"testing"
)

func TestAtoI(t *testing.T) {
	type testCase struct {
		input string
		want  int
	}
	cases := []testCase{
		{input: "0", want: 0},
		{input: "-0", want: 0},
		{input: "123", want: 123},
		{input: "-123", want: -123},
	}
	for _, tc := range cases {
		t.Run(fmt.Sprintf("AtoI(\"%s\")", tc.input), func(t *testing.T) {
			if got := AtoI(tc.input); got != tc.want {
				t.Errorf("AtoI(\"%s\") = %d, want %d", tc.input, got, tc.want)
			}
		})

	}
}
