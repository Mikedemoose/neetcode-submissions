func checkInclusion(s1 string, s2 string) bool {
	if len(s2) < len(s1) {
		return false
	}
	s1LetterCount := make(map[byte]int)
	l1 := len(s1)

	for i:= range l1 {
		s1LetterCount[s1[i]]++
	}

	s2CurrCountMap := make(map[byte]int)
	left, right := 0, l1-1
	for i:=range l1 {
		s2CurrCountMap[s2[i]]++
	}

	for {
		if checkIfWordLetterCountsMatch(s1LetterCount, s2CurrCountMap) {
			return true
		}
		s2CurrCountMap[s2[left]]--
		left++
		right++
		if right>=len(s2) {
			break
		}
		s2CurrCountMap[s2[right]]++
	}

	return false
}

func checkIfWordLetterCountsMatch(m1, m2 map[byte]int) bool {
	// assumes both m1 and m2 contains the same number of total letters
	fmt.Println(m1, m2)
	for k, v := range m1 {
		if val, ok := m2[k]; !ok || val != v {
			return false
		}
	}
	return true
}
