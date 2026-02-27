package exercise5

func Solution_FilterEven(arr []int) []int {
	var out []int
	for _, v := range arr {
		if v%2 == 0 {
			out = append(out, v)
		}
	}
	return out
}
