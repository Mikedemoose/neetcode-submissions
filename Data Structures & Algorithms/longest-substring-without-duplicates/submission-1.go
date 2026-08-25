func lengthOfLongestSubstring(s string) int {
	charFreqMap := make(map[byte]int)
	maxLength := 0
	currLength := 0

	left, right := 0, 0

	for left <= right && right < len(s) {
		if val, ok := charFreqMap[s[right]]; ok {
			if val == 0 {
				charFreqMap[s[right]] = 1
				currLength++
				right++
			} else {
				if currLength > maxLength {
					maxLength = currLength
				}
				charFreqMap[s[left]]=0
				left++
				currLength--
			}
		} else {
			charFreqMap[s[right]] = 1
			currLength++
			right++
		}
	}
	if currLength > maxLength {
		maxLength = currLength
	}
	return maxLength
}
