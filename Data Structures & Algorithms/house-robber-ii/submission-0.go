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

	n := len(nums)
	var n_2, n_1, temp int

    // case 1: pick the first house, omit the last house
	n_2, n_1 = nums[0], max(nums[0], nums[1])

	for i:=2; i<n-1; i++ {
		temp = n_1
		n_1 = max(n_2+nums[i], n_1)
		n_2 = temp
	}

	case_1_max := n_1
	if len(nums) == 2 {
		return case_1_max
	}

	// case 2: omit the first house, pick the last house
	if len(nums) == 3 {
		return max(case_1_max, nums[2])
	}

	n_2, n_1 = nums[1], max(nums[1], nums[2])
	for i:=3; i<n; i++ {
		temp = n_1
		n_1 = max(n_2+nums[i], n_1)
		n_2 = temp
	}

	return max(case_1_max, n_1)
}
