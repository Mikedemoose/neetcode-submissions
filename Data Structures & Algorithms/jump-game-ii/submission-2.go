func jump(nums []int) int {
    numJumps := 0
	currLimit := 0
	var i int

	for currLimit < len(nums)-1 {
		tempLimit := currLimit
		for i <= tempLimit {
			if i + nums[i] > currLimit {
				currLimit = i + nums[i]
			}
			i++
		}
		numJumps++
	}
	return numJumps
}
