func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	valMap := make(map[int]struct{})
	for _, num := range nums {
		valMap[num] = struct{}{}
	}

	visitedMap := make(map[int]int)
	maxLength := 1

	for _, num := range nums {
		if _, ok := visitedMap[num]; ok {
			continue
		}
		maxSubseqLength := 1
		prev := num-1
		for {
			if _, ok1:= valMap[prev]; !ok1 {
				break
			}
			if prevSum, ok2 := visitedMap[prev]; ok2 {
				maxSubseqLength += prevSum
				break
			}
			maxSubseqLength++
			prev--
		}
		visitedMap[num] = maxSubseqLength
		if maxSubseqLength > maxLength {
			maxLength = maxSubseqLength
		}
	}

	return maxLength
}
