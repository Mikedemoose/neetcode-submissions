func combinationSum(nums []int, target int) [][]int {
    result := make([][]int, 0)

	sumMap := make(map[int][][]int)
	for _, num := range nums {
		sumMap[num] = [][]int{{num}}
	}

	for i := range target+1 {
		for _, num := range nums {
			if num < i {
				if vals, ok := sumMap[i-num]; ok {
					for _, val := range vals {
						newVal := append([]int{}, val...)
						if num>=newVal[len(newVal)-1] {
							sumMap[i] = append(sumMap[i], append(newVal, num))
						}
					}
				}
			}
		}
	}

	if vals, ok := sumMap[target]; ok {
		return vals
	}
	return result
}
