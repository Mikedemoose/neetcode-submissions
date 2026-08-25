func longestConsecutive(nums []int) int {
	maxLength := 0

	valMap := make(map[int]struct{})
	for _, num := range nums {
		valMap[num] = struct{}{}
	}

	for _, num := range nums {
		if _, ok := valMap[num+1]; ok {
			continue
		}
		
		maxSubseqLength := 1
		prev := num-1
		
		for {
			if _, ok1:= valMap[prev]; !ok1 {
				break
			}
			maxSubseqLength++
			prev--
		}

		if maxSubseqLength > maxLength {
			maxLength = maxSubseqLength
		}
	}

	return maxLength
}
