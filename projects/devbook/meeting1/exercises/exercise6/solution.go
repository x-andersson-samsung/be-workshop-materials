package exercise6

import (
	"fmt"
)

func Solution_Divide(dividend, divisor int) (int, error) {
	if divisor == 0 {
		return 0, fmt.Errorf("dividing by zero")
	}
	return dividend / divisor, nil
}
