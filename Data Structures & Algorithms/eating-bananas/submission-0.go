func canFinish(piles []int, h int, k int) bool {
	curr := 0
	for _, p := range piles {
		hours_for_p := (p + k-1)/k
		curr += hours_for_p
	}

	if curr > h {
		return false
	}

	return true
}

func minEatingSpeed(piles []int, h int) int {
	
	// find the max pile
	maxVal := piles[0]
	for _, p := range piles {
		if p > maxVal {
			maxVal = p
		}
	}

	// binary search range of k to find the min k satisfying canFinish
	left, right := 1, maxVal
	var mid int
	for left < right {
		mid = (left+right)/2

		if canFinish(piles, h, mid) {
			right = mid
		} else {
			left = mid+1
		}
	} 

	return right
}
