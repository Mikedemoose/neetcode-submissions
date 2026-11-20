func lengthOfLongestSubstring(s string) int {
	countArray := make([]int, 128)

	left, right := 0, 0
	maxLength, currLength := 0, 0

	getIndex := func(a byte) int {
		return int(a)
	}

	for right < len(s) {
		for right < len(s) {
			if countArray[getIndex(s[right])] > 0 {
				break
			}
			countArray[getIndex(s[right])] = 1
			right++
		}
		currLength = right - left
		maxLength = max(maxLength, currLength)

		// remove letters till condition is satisfied
		for left < right && right < len(s) {
			if s[left] == s[right] {
				left++
				break
			}
			countArray[getIndex(s[left])] = 0
			left++
		}
		right++
	}

	return maxLength
}