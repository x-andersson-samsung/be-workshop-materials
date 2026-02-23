package exercise6

import (
	"fmt"
	"testing"
)

func TestDivide(t *testing.T) {
	type testCase struct {
		dividend int
		divisor  int
		wantQuot int
		wantErr  bool
	}

	cases := []testCase{
		{
			dividend: 10,
			divisor:  2,
			wantQuot: 5,
			wantErr:  false,
		},
		{
			dividend: 7,
			divisor:  3,
			wantQuot: 2,
			wantErr:  false,
		},
		{
			dividend: 10,
			divisor:  0,
			wantQuot: 0,
			wantErr:  true,
		},
		{
			dividend: -10,
			divisor:  2,
			wantQuot: -5,
			wantErr:  false,
		},
		{
			dividend: 0,
			divisor:  5,
			wantQuot: 0,
			wantErr:  false,
		},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("Divide(%d, %d)", tc.dividend, tc.divisor), func(t *testing.T) {
			gotQuot, gotErr := Divide(tc.dividend, tc.divisor)

			// Check if error state matches expectation
			if (gotErr != nil) != tc.wantErr {
				t.Errorf("Divide(%d, %d) error = %v, wantErr %t", tc.dividend, tc.divisor, gotErr, tc.wantErr)
				return
			}

			// Check quotient if no error expected
			if !tc.wantErr && gotQuot != tc.wantQuot {
				t.Errorf("Divide(%d, %d) = %d, want %d", tc.dividend, tc.divisor, gotQuot, tc.wantQuot)
			}
		})
	}
}
