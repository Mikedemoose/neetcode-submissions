func checkValidString(s string) bool {
    charCount := make(map[byte]int)
	for i := range len(s) {
		charCount[s[i]]++
	}

	return isValid(s, 0, 0, 0, charCount['('], charCount[')'], charCount['*'])

}

func isValid(s string, index int, currOpen, currClosed int, totalOpen, totalClosed int, availWildCardCount int) bool {

	if currOpen == currClosed && totalOpen == currOpen && totalClosed == currClosed {
		return true
	}

	if index == len(s) || currOpen < currClosed || totalOpen > totalClosed+availWildCardCount || totalClosed > totalOpen + availWildCardCount {
		return false
	}

	if s[index] == '(' {
		return isValid(s, index+1, currOpen+1, currClosed, totalOpen, totalClosed, availWildCardCount)
	} else if s[index] == ')' {
		return isValid(s, index+1, currOpen, currClosed+1, totalOpen, totalClosed, availWildCardCount)
	}

	return isValid(s, index+1, currOpen+1, currClosed, totalOpen+1, totalClosed, availWildCardCount-1) ||
		isValid(s, index+1, currOpen, currClosed+1, totalOpen, totalClosed+1, availWildCardCount-1) ||
		isValid(s, index+1, currOpen, currClosed, totalOpen, totalClosed, availWildCardCount-1)
}
