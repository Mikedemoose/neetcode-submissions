func groupAnagrams(strs []string) [][]string {
	result := make([][]string, 0)

	sortLetters := func(s string) string {
		res := []rune(s)
		sort.Slice(res, func(i, j int) bool {
			return res[i] < res[j]
		})
		return string(res)
	}

	strMap := make(map[string][]string)

	for _, str := range strs {
		key := sortLetters(str)
		strMap[key] = append(strMap[key], str)
	}

	for _, val := range strMap {
		result = append(result, val)
	}

	return result
}