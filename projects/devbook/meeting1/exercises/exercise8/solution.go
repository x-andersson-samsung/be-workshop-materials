package exercise8

// Write a function called RemoveDuplicates accepting a []int and returning a new slice without duplicates.
func Solution_RemoveDuplicates(arr []int) []int {
	out := make([]int, 0, len(arr))
	for _, val := range arr {
		found := false
		for _, v2 := range out {
			if val == v2 {
				found = true
				break
			}
		}
		if !found {
			out = append(out, val)
		}
	}
	return out
}

func Solution_RemoveDuplicates_Map(arr []int) []int {
	check := make(map[int]bool)
	out := make([]int, 0, len(arr))

	for _, val := range arr {
		if _, ok := check[val]; ok {
			continue
		}
		out = append(out, val)
		check[val] = true
	}
	return out
}
