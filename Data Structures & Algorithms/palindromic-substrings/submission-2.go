func countSubstrings(s string) int {
	totalCount := 0

	for i, _ := range s {
		for j:=0; j<=i && i+j<len(s); j++ {
			if s[i-j] != s[i+j] {
				break
			}
			totalCount++
		}
		
		if i<1 || s[i] != s[i-1] {
			continue
		}

		for j:=0; j<=i-1 && i+j<len(s); j++ {
			if s[i-j-1] != s[i+j] {
				break
			}
			totalCount++
		}
	}

	return totalCount
}
