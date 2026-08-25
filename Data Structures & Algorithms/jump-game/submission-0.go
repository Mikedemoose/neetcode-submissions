func canJump(nums []int) bool {
    visitable := make([]bool, len(nums))
	visitable[0] = true

	for i:=0; i<len(nums)-1; i++ {
		if visitable[i] {
			num := nums[i]

			for j := 1; j<=num; j++ {
				if i+j >= len(nums) {
					break
				}
				visitable[i+j] = true
			}
		}
	}

	return visitable[len(nums)-1]
}
