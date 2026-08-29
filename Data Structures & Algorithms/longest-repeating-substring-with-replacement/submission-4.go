func characterReplacement(s string, k int) int {
	lenMap := make(map[byte]int)
	for i:=range k+1 {
		lenMap[s[i]]++
	}
	maxStringLen := k+1

	findMaxLen := func() (maxLen int) {
		for _, val := range lenMap {
			maxLen = max(maxLen, val)
		}
		return
	}

	left, right := 0, k // next element

	for left <= right && right < len(s) {
		// fmt.Println("left:", left, "right:", right)
		// fmt.Println("Current string:", s[left:right+1])
		// fmt.Println("CUrrent map: ", lenMap)

		currLen := right-left+1

		if currLen-findMaxLen() > k {
			lenMap[s[left]]--
			left++
		} else {
			maxStringLen = max(maxStringLen, currLen)
			right++
			if right < len(s) {
				lenMap[s[right]]++
			}
		}
	}

	return maxStringLen
}
