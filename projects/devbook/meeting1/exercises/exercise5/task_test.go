package exercise5

import (
	"fmt"
	"testing"
)

func TestFilterEven(t *testing.T) {
	type testCase struct {
		input []int
		want  []int
	}

	cases := []testCase{
		{
			input: []int{1, 2, 3, 4, 5, 6},
			want:  []int{2, 4, 6},
		},
		{
			input: []int{1, 3, 5},
			want:  []int{},
		},
		{
			input: []int{2, 4, 6},
			want:  []int{2, 4, 6},
		},
		{
			input: []int{},
			want:  []int{},
		},
		{
			input: []int{0, -1, -2, -3, -4},
			want:  []int{0, -2, -4},
		},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("FilterEven(%#v)", tc.input), func(t *testing.T) {
			got := FilterEven(tc.input)
			
			// Check if slices have the same length
			if len(got) != len(tc.want) {
				t.Errorf("FilterEven(%#v) = %#v, want %#v", tc.input, got, tc.want)
				return
			}
			
			// Check if all elements match
			for i, v := range tc.want {
				if got[i] != v {
					t.Errorf("FilterEven(%#v) = %#v, want %#v", tc.input, got, tc.want)
					return
				}
			}
		})
	}
}
