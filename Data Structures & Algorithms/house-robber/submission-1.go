func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	maxLootTillHouse := make([]int, len(nums))
	maxLootTillHouse[0], maxLootTillHouse[1] = nums[0], max(nums[0], nums[1])

	for i:=2; i<len(nums); i++ {
		maxLootTillHouse[i] = max(maxLootTillHouse[i-2] + nums[i], maxLootTillHouse[i-1])
	}

	return maxLootTillHouse[len(nums)-1]
}
