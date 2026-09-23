func reverse(x int) int {
	isNegative := x < 0

	if isNegative {
		x = -x
	}

	finalVal := 0

	for x > 0 {
		finalVal *= 10
		finalVal += x%10
		x/=10
	}

	if isNegative {
		finalVal = -finalVal
	}

	if finalVal > math.MaxInt32 || finalVal < math.MinInt32 {
		return 0
	}

	return finalVal

}
