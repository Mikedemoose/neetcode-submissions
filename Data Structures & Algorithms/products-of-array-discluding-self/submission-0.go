func productExceptSelf(nums []int) []int {
	prefixProduct := make([]int, len(nums))
	suffixProduct := make([]int, len(nums))

	prefixProduct[0] = 1
	suffixProduct[len(nums)-1] = 1
	for i := 1; i < len(nums); i++ {
		prefixProduct[i] = prefixProduct[i-1] * nums[i-1]
	}
	for i := len(nums)-2; i >= 0; i-- {
		suffixProduct[i] = suffixProduct[i+1] * nums[i+1]
	}

	finalResult := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		finalResult[i] = prefixProduct[i] * suffixProduct[i]
	}

	return finalResult
}
