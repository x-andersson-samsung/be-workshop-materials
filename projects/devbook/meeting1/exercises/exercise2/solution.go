package exercise2

func Solution_AtoI(s string) int {
	isNegative := false
	if s[0] == '-' {
		isNegative = true
		s = s[1:]
	}

	number := 0
	for _, c := range s {
		number = number*10 + int(c-'0')
	}

	if isNegative {
		return -number
	}
	return number
}
