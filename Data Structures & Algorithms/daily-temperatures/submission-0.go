func dailyTemperatures(temperatures []int) []int {
	stack := make([]int, 0)
	stack = append(stack, 0)
	finalArray := make([]int, len(temperatures))

	for i:=1; i<len(temperatures); i++ {
		if temperatures[stack[len(stack)-1]] >= temperatures[i] {
			stack = append(stack, i)
		} else {
			// pop till curr temp value is smaller
			currTempIndex := stack[len(stack)-1]
			for len(stack)>0 && temperatures[currTempIndex] < temperatures[i] {
				finalArray[currTempIndex] = i-currTempIndex
				stack = stack[:len(stack)-1]
				if len(stack) > 0 {
					currTempIndex = stack[len(stack)-1]
				}
			}
			stack = append(stack, i)
		}
	}

	return finalArray

}
