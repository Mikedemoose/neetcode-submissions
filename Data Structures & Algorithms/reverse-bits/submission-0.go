func reverseBits(n int) int {
	finalVal := 0

	for _ = range 32 {
		finalVal <<= 1
		finalVal += n%2

		n >>= 1
	}

	return finalVal
}
