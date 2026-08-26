func missingNumber(nums []int) int {
	max_n := len(nums)

	expected_sum := (max_n)*(max_n + 1)/2

	for _, num := range nums {
		expected_sum -= num
	}

	return expected_sum
}
