package exercise4

import (
	"fmt"
	"testing"
)

func TestMergeMap(t *testing.T) {
	type testCase struct {
		input1 map[string]string
		input2 map[string]string
		want   map[string]string
	}

	cases := []testCase{
		{
			input1: map[string]string{"a": "1", "b": "2"},
			input2: map[string]string{"c": "3", "d": "4"},
			want:   map[string]string{"a": "1", "b": "2", "c": "3", "d": "4"},
		},
		{
			input1: map[string]string{"a": "1"},
			input2: map[string]string{"a": "2", "b": "3"},
			want:   map[string]string{"a": "1", "b": "3"},
		},
		{
			input1: map[string]string{},
			input2: map[string]string{"a": "1"},
			want:   map[string]string{"a": "1"},
		},
		{
			input1: map[string]string{"a": "1"},
			input2: map[string]string{},
			want:   map[string]string{"a": "1"},
		},
		{
			input1: map[string]string{},
			input2: map[string]string{},
			want:   map[string]string{},
		},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("MergeMap(%#v, %#v)", tc.input1, tc.input2), func(t *testing.T) {
			input1Len := len(tc.input1)
			input2Len := len(tc.input2)

			got := MergeMap(tc.input1, tc.input2)

			// Check if maps have the same length
			if len(got) != len(tc.want) {
				t.Errorf("MergeMap(%#v, %#v) = %#v, want %#v", tc.input1, tc.input2, got, tc.want)
				return
			}

			// Check if all key-value pairs match
			for k, v := range tc.want {
				if got[k] != v {
					t.Errorf("MergeMap(%#v, %#v) = %#v, want %#v", tc.input1, tc.input2, got, tc.want)
					return
				}
			}

			if input1Len != len(tc.input1) {
				t.Errorf("MergeMap(%#v, %#v) = map1 was modified", tc.input1, tc.input2)
				return
			}

			if input2Len != len(tc.input2) {
				t.Errorf("MergeMap(%#v, %#v) = map2 was modified", tc.input1, tc.input2)
				return
			}
		})
	}
}
