package calculator

import (
	"testing"
)

func TestCalculator_Add(t *testing.T) {
	c := Calculator{}

	want := 2
	a, b := 1, 1
	if got := c.Add(1, 1); got != want {
		t.Errorf("Add(%d, %d) = %d, expected %d\n", a, b, got, want)
	}
}
