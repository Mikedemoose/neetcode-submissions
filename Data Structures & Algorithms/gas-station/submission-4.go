func canCompleteCircuit(gas []int, cost []int) int {
	index := -1
	currSum, totalSum := 0, 0
	var item int

	for i := range len(cost) {
		item = cost[i]-gas[i]

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
