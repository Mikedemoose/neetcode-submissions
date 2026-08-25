func canJump(nums []int) bool {
    maxReachable := 0

	for i, num := range nums {
		if maxReachable < i {
			return false
		}
		if i+num > maxReachable {
			maxReachable = i+num
		}

		if maxReachable >= len(nums)-1 {
			return true
		}
	}

	return false
}
