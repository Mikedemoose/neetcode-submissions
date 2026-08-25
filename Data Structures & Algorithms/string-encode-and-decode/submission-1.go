type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	delimiter := "#"
	finalString := ""

	for _, str := range strs {
		finalString = fmt.Sprintf("%s%d%s%s", finalString, len(str), delimiter, str)
	}

	return finalString
}

func (s *Solution) Decode(encoded string) []string {
	delimiter := "#"
	strs := make([]string, 0)

	left, right := 0, 0
	for right<len(encoded) {
		if string(encoded[right]) == delimiter {
			numChars, _ := strconv.Atoi(encoded[left:right])
			strs = append(strs, encoded[right+1:right+1+numChars])

			left, right = right+1+numChars, right+1+numChars
		} else {
			right++
		}
	}

	return strs
}
