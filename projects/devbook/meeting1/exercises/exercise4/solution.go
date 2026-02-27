package exercise4

func Solution_MergeMap(map1, map2 map[string]string) map[string]string {
	result := make(map[string]string)
	for k, v := range map2 {
		result[k] = v
	}
	for k, v := range map1 {
		result[k] = v
	}
	return result
}
