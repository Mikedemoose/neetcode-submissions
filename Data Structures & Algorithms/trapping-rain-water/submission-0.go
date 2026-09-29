func trap(height []int) int {

	left := 0
	total := 0

	for left < len(height)-1 {
		// find the max height which is less than/equal to left
		right := left + 1
		maxheight := -1

		for i := left+1; i<len(height); i++ {
			if height[i] >= height[left] {
				maxheight = height[i]
				right = i
				break
			}
			if height[i] > maxheight {
				maxheight = height[i]
				right = i
			}
		}

		maxheight = min(height[left], maxheight)

		for i := left+1; i <= right-1; i++ {
			total += max(0, maxheight-height[i])
		}

		left = right 
	}

	return total
}
