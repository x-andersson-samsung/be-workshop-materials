package exercise1

import (
	"fmt"
	"testing"
)

func TestFibonacci(t *testing.T) {
	type testCase struct {
		input int
		want  int
	}
	cases := []testCase{
		{input: 0, want: 0},
		{input: 1, want: 1},
		{input: 2, want: 1},
		{input: 3, want: 2},
		{input: 4, want: 3},
		{input: 5, want: 5},
		{input: 10, want: 55},
		{input: 15, want: 610},
		{input: 20, want: 6765},
	}
	for _, tc := range cases {
		t.Run(fmt.Sprintf("Fibonacci(%d)", tc.input), func(t *testing.T) {
			if got := Fibonacci(tc.input); got != tc.want {
				t.Errorf("Fibonacci(%d) = %d, want %d", tc.input, got, tc.want)
			}
		})

	}
}
