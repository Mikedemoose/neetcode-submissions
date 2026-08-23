func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxArea(heights []int) int {
	left, right := 0, len(heights)-1

	currMax := min(heights[left], heights[right]) * 1

	for left < right {
		currMax = max(currMax, (right-left)*min(heights[left], heights[right]))

		if heights[left] < heights[right] {
			left++
		} else {
			right--
		}
	}

	return currMax
	
}
