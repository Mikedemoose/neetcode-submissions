func canCompleteCircuit(gas []int, cost []int) int {
    need := make([]int, len(cost))
	for i := range len(cost) {
		need[i] = cost[i]-gas[i]
	}

	index := -1
	currSum, totalSum := 0, 0

	for i, item := range need {
		currSum += item
		totalSum += item
		if index == -1 {
			if item <= 0 {
				index = i
			}
		} else {
			if currSum > 0 {
				index = -1
				currSum = 0
			}
		}
	}

	if currSum > 0 || totalSum > 0 {
		index = -1
	}

	return index
}
