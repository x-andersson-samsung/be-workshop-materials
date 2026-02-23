package exercise8

import (
	"fmt"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	type testCase struct {
		input []int
		want  []int
	}

	cases := []testCase{
		{
			input: []int{1, 2, 2, 3, 3, 3, 4},
			want:  []int{1, 2, 3, 4},
		},
		{
			input: []int{1, 1, 1},
			want:  []int{1},
		},
		{
			input: []int{1, 2, 3, 4, 5},
			want:  []int{1, 2, 3, 4, 5},
		},
		{
			input: []int{},
			want:  []int{},
		},
		{
			input: []int{0, -1, -1, 0, 2, 2},
			want:  []int{0, -1, 2},
		},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("RemoveDuplicates(%#v)", tc.input), func(t *testing.T) {
			got := RemoveDuplicates(tc.input)
			
			// Check if slices have the same length
			if len(got) != len(tc.want) {
				t.Errorf("RemoveDuplicates(%#v) = %#v, want %#v", tc.input, got, tc.want)
				return
			}
			
			// Check if all elements match
			for i, v := range tc.want {
				if got[i] != v {
					t.Errorf("RemoveDuplicates(%#v) = %#v, want %#v", tc.input, got, tc.want)
					return
				}
			}
		})
	}
}
