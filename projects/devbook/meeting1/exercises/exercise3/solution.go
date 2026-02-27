package exercise3

func Solution_ArrStats(array []int) (int, int, float64) {
	if len(array) == 0 {
		return 0, 0, 0
	}

	var (
		max   = array[0]
		min   = array[0]
		count int
		sum   int
	)

	for _, v := range array {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}

		sum += v
		count++
	}

	return min, max, float64(sum) / float64(count)
}
