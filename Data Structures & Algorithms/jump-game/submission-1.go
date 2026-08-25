func canJump(nums []int) bool {
    maxReachable := 0

	for i, num := range nums {
		if maxReachable < i {
			return false
		}
		currReachable := i+num
		if currReachable > maxReachable {
			maxReachable = currReachable
		}

		if maxReachable >= len(nums)-1 {
			return true
		}
	}

	return false
}
