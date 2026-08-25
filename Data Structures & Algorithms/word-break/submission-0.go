func wordBreak(s string, wordDict []string) bool {
    wordMap := make(map[string]struct{})
	for _, word := range wordDict {
		wordMap[word] = struct{}{}
	}

	queueMap := make(map[int]struct{})
	queue := []int{-1}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		for i:=curr+2; i<=len(s); i++ {
			if _, ok := wordMap[s[curr+1:i]]; ok {
				if i == len(s) {
					return true
				}
				if _, ok1:= queueMap[i-1]; !ok1 {
					queue = append(queue, i-1)
					queueMap[i-1] = struct{}{}
				}
			}
		}
	}

	return false


}
