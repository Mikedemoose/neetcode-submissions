func wordBreak(s string, wordDict []string) bool {
    wordMap := make(map[string]struct{})
	for _, word := range wordDict {
		wordMap[word] = struct{}{}
	}

	state := make([]bool, len(s)+1)
	state[0] = true

	for i:=0; i<=len(s); i++ {
		for j:=0; j<i; j++ {
			if state[j] {
				if _, ok:= wordMap[s[j:i]]; ok {
					state[i] = true
					break
				}
			}
		}
	}

	return state[len(s)]
}
