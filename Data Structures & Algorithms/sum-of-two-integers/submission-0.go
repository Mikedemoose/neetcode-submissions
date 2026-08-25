func getSum(a int, b int) int {
	s1 := a ^ b
	c1 := (a & b) << 1

	for c1 != 0 {
		s1, c1 = s1 ^ c1, (s1 & c1) << 1
	}

	return s1
}
