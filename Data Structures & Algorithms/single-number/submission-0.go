func singleNumber(nums []int) int {
	finalVal := nums[0]
	for i:=1; i<len(nums); i++ {
		finalVal ^= nums[i]
	}

	return finalVal
}
