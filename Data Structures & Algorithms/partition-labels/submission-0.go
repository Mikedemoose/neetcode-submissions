func partitionLabels(s string) []int {
	subStringCount := make([]int, 0)
	letterMap := make(map[byte]SubString)
	for i := range len(s) {
		l := s[i]
		if val, ok := letterMap[l]; ok {
			val.end = i
			letterMap[l] = val
		} else {
			letterMap[l] = SubString {
				start: i,
				end: i,
			}
		}
	}
	letterMapArray := make([]SubString, 0)
	for _, v := range letterMap {
		letterMapArray = append(letterMapArray, v)
	}

	sort.Slice(letterMapArray, func(i, j int)bool {
		return letterMapArray[i].start < letterMapArray[j].start
	})


	curr := letterMapArray[0]
	for i := 1; i < len(letterMapArray); i++ {
		if letterMapArray[i].start < curr.end {
			curr.end = max(curr.end, letterMapArray[i].end)
		} else {
			subStringCount = append(subStringCount, curr.end-curr.start+1)
			curr = letterMapArray[i]
		}
	}
	subStringCount = append(subStringCount, curr.end-curr.start+1)

	return subStringCount
}

type SubString struct {
	start int
	end int
}
