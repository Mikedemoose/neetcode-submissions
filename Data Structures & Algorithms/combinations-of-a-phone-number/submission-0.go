func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return make([]string, 0)
	}

	queue := []string {""}

	for i := range len(digits) {

		currLetter := digits[i]
		chars := getCharsForDigit(currLetter)
		newQueue := make([]string, 0)

		for _, char := range chars {
			for _, item := range queue {
				newQueue = append(newQueue, item + string(char))
			}
		}

		queue = newQueue

	}

	return queue
	
}

func getCharsForDigit(digit byte) []byte {
	digitMap := map[byte] []byte {
		'2': []byte{'a', 'b', 'c'},
		'3': []byte{'d', 'e', 'f'},
		'4': []byte{'g', 'h', 'i'},
		'5': []byte{'j', 'k', 'l'},
		'6': []byte{'m', 'n', 'o'},
		'7': []byte{'p', 'q', 'r', 's'},
		'8': []byte{'t', 'u', 'v'},
		'9': []byte{'w', 'x', 'y', 'z'},
	}

	return digitMap[digit]
}
