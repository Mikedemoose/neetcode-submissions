func insert(intervals [][]int, newInterval []int) [][]int {
    if len(intervals) == 0 {
		return [][]int{newInterval}
	}

	finalList := make([][]int, 0)

	// find the highest interval which is less than new interval
	currIndex := 0
	for i:=0; i<len(intervals); i++ {
		currIndex = i
		if intervals[i][1] >= newInterval[0] {
			currIndex = i-1
			break
		} else {
			finalList = append(finalList, intervals[i])
		}

	}

	// merge new interval with all the overlapping intervals
	for i:=currIndex+1; i<len(intervals); i++ {
		currIndex = i
		if intervals[i][0] <= newInterval[1] {
			newInterval[0] = min(intervals[i][0], newInterval[0])
			newInterval[1] = max(intervals[i][1], newInterval[1])
		} else {
			currIndex = i-1
			break
		}
	}

	// add new interval and then the rest of the intervals
	finalList = append(finalList, newInterval)
	finalList = append(finalList, intervals[currIndex+1:]...)

	return finalList
}
