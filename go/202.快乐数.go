func isHappy(n int) bool {
	if n == 1 {
		return true
	}

	dict := make(map[int]struct{})
	dict[n] = struct{}{}
	for n != 1 {
		n = getSum(n)
		if _, ok := dict[n]; ok {
			return false
		}
		dict[n] = struct{}{}
	}

	return true
}

func getSum(m int) int {
	sum := 0
	for m > 0 {
		sum += (m % 10) * (m % 10)
		m = m / 10
	}
	return sum
}