func hammingWeight(n int) int {
	total := 0
	for n > 0 {
		total += n%2
		n = n>>1
	}

	return total
}
