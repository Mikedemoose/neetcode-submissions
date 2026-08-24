func countSubstrings(s string) int {
	totalCount := 0

	for i, _ := range s {
		left, right := i, i
		for left >= 0 && right < len(s) {
			if s[left] != s[right] {
				break
			}
			totalCount++
			left--
			right++
		}
	}

	for i:=1; i<len(s); i++ {
		if s[i] != s[i-1] {
			continue
		}

		left, right := i-1, i
		for left >= 0 && right < len(s) {
			if s[left] != s[right] {
				break
			}
			totalCount++
			left--
			right++
		}
	}

	return totalCount
}
