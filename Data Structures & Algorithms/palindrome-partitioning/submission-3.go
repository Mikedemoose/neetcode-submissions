func partition(s string) [][]string {
	if len(s) == 1 {
		return [][]string{{s}}
	}

	finalRes := make([][]string, 0)

	for i:=0; i<len(s); i++ {
		if checkIfPalindrome(s[:i+1]) {
			restPartitions := partition(s[i+1:])
			for _, item := range restPartitions {
				finalRes = append(finalRes, append([]string{s[:i+1]}, item...))
			}
		}
	}
	if len(s) > 0 && checkIfPalindrome(s) {
		finalRes = append(finalRes, []string{s})
	}
	return finalRes

}

func checkIfPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left <= right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}