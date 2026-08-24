func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func rob(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	n_2, n_1 := nums[0], max(nums[0], nums[1])

	var temp int
	for i:=2; i<n; i++ {
		temp = n_1
		n_1 = max(n_2 + nums[i], n_1)
		n_2 = temp
	}

	return n_1
}
