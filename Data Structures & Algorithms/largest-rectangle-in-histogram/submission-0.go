func largestRectangleArea(heights []int) int {
	sizeMap := make([]map[int]int, 0)
	maxRectSize := heights[0]

	var compareAndUpdate func(int)

	compareAndUpdate = func(val int) {
		// fmt.Println("comparing for", val)
		if val > maxRectSize {
			// fmt.Println(val, "is greater than", maxRectSize)
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
				// fmt.Println("setting", v+1, "for", newKey)
				compareAndUpdate(newKey * (v+1))
			}
		}

		if sizeMap[i][heights[i]] == 0 {
			sizeMap[i][heights[i]] = 1
			compareAndUpdate(heights[i])
		}

		// fmt.Println("At index", i, "map is", sizeMap[i])
	}

	return maxRectSize
}
