func minWindow(s string, t string) string {
    tMap := make(map[byte]int)
	numMatchesReq := len(t)

	for i := range numMatchesReq {
		tMap[t[i]]++
	}

	minVal := ""

	if len(t) == 0 {
		return minVal
	}

	sMap := make(map[byte]int)

	left, right := 0, 0
	for left < len(s)-len(t)+1 {
		// find left
		for left < len(s)-len(t)+1 {
			// letter not matching
			if _, ok := tMap[s[left]]; !ok {
				left++
				continue
			}

			// letter is matching
			break
		}

		// fmt.Println("Found left: ", left)
		if left >= len(s)-len(t)+1 {
			break
		}

		
		// find right
		right = max(right, left)
		for numMatchesReq > 0 && right < len(s) {
			// letter not matching
			if _, ok := tMap[s[right]]; !ok {
				right++
				continue
			}

			// letter matching
			// fmt.Println("found match in right at", right)
			sMap[s[right]]++
			if tMap[s[right]] >= sMap[s[right]] {
				numMatchesReq--
			}
			right++
		}

		if numMatchesReq == 0 && (len(minVal) == 0 || len(minVal) > right-left) {
			minVal = s[left:right]
			// fmt.Println("setting minVal to", minVal)
		}

		sMap[s[left]]--
		if sMap[s[left]] < tMap[s[left]] {
			numMatchesReq++
			// fmt.Println("increasing numMatches required to", numMatchesReq)
		}
		left++
	}

	return minVal

}
