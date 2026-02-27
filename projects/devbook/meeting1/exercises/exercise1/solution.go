package exercise1

func Solution_Fibonacci(n int) int {
	a, b := 0, 1
	for range n {
		a, b = b, a+b
	}

	return a
}
