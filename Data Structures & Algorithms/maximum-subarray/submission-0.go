func maxSubArray(nums []int) int {
	currentMax := nums[0]
    
	// save the max sum ending at i
	for i:=1; i<len(nums); i++ {
		if nums[i] + nums[i-1] > nums[i] {
			nums[i] = nums[i] + nums[i-1]
		}

		if nums[i] > currentMax {
			currentMax = nums[i]
		}
	}

	return currentMax
}
