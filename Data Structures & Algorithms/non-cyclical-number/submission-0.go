func isHappy(n int) bool {
	numMap := make(map[int]struct{})
	numMap[n] = struct{}{}

	for n != 1 {
		n = getSumOfDigitSquares(n)
		if _, ok := numMap[n]; ok {
			return false
		}
		numMap[n] = struct{}{}
	}

	return true
}

func getSumOfDigitSquares(n int) int {
	sum := 0

	for n > 0 {
		sum += (n%10)*(n%10)
		n /= 10
	}

	return sum
}
