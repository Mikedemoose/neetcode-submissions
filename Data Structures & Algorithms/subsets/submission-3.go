func subsets(nums []int) [][]int {
	if len(nums) == 0 {
		return make([][]int, 0)
	} else if len(nums) == 1 {
		return [][]int{{}, {nums[0]}}
	}

	currElement := nums[0]
	nextSubsets := subsets(nums[1:])
	nextSubsetsLength := len(nextSubsets)

	for i:=range nextSubsetsLength {
		newItem := make([]int, len(nextSubsets[i]))
		copy(newItem, nextSubsets[i])
		nextSubsets = append(nextSubsets, append(newItem, currElement))
	}

	return nextSubsets
}
