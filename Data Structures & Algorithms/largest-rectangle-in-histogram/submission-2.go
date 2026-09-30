func largestRectangleArea(heights []int) int {
	sizeMap := make([]map[int]int, 0)
	maxRectSize := heights[0]

	var compareAndUpdate func(int)

	compareAndUpdate = func(val int) {
		if val > maxRectSize {
			maxRectSize = val
		}
	}

	for range len(heights) {
		sizeMap = append(sizeMap, make(map[int]int))
	}

	sizeMap[0][heights[0]] = 1

	for i:=1; i<len(heights); i++ {
		prevMap := sizeMap[i-1]

		for k, v := range prevMap {
			newKey := min(k, heights[i])
			if newKey != 0 && sizeMap[i][newKey] == 0 {
				sizeMap[i][newKey] = v+1
				compareAndUpdate(newKey * (v+1))
			}
		}

		if sizeMap[i][heights[i]] == 0 {
			sizeMap[i][heights[i]] = 1
			compareAndUpdate(heights[i])
		}

	}

	return maxRectSize
}
