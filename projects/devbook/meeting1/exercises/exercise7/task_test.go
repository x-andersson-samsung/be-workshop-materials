package exercise7

import (
	"fmt"
	"testing"
)

func TestGetArea(t *testing.T) {
	type testCase struct {
		shape Shape
		want  float64
		name  string
	}

	cases := []testCase{
		{
			shape: Rectangle{Width: 2, Height: 3},
			want:  6,
			name:  "Rectangle 2x3",
		},
		{
			shape: Rectangle{Width: 5, Height: 4},
			want:  20,
			name:  "Rectangle 5x4",
		},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("GetArea(%s)", tc.name), func(t *testing.T) {
			got := tc.shape.GetArea()
			// Allow for small floating point differences
			diff := got - tc.want
			if diff < -0.0001 || diff > 0.0001 {
				t.Errorf("GetArea(%s) = %f, want %f", tc.name, got, tc.want)
			}
		})
	}
}
