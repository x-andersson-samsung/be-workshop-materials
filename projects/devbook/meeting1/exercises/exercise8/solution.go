package exercise8

// Write a function called RemoveDuplicates accepting a []int and returning a new slice without duplicates.
func RemoveDuplicates(arr []int) []int {
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
