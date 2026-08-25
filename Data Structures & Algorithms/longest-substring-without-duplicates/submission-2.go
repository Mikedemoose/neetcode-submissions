func lengthOfLongestSubstring(s string) int {
	charFreqMap := make(map[byte]struct{})
	maxLength := 0
	currLength := 0

	left, right := 0, 0

	for left <= right && right < len(s) {
		if _, ok := charFreqMap[s[right]]; ok {
			if currLength > maxLength {
				maxLength = currLength
			}
			delete(charFreqMap, s[left])
			left++
			currLength--
		} else {
			charFreqMap[s[right]] = struct{}{}
			currLength++
			right++
		}
	}
	if currLength > maxLength {
		maxLength = currLength
	}
	return maxLength
}
