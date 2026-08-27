func insert(intervals [][]int, newInterval []int) [][]int {
    if len(intervals) == 0 {
		return [][]int{newInterval}
	}

	finalList := make([][]int, 0)

	// find the highest interval which is less than new interval
	maxIntervalLessThanNew := 0
	for i:=0; i<len(intervals); i++ {
		maxIntervalLessThanNew = i
		if intervals[i][1] >= newInterval[0] {
			maxIntervalLessThanNew = i-1
			break
		} else {
			// fmt.Println("adding smaller interval", intervals[i])
			finalList = append(finalList, intervals[i])
		}

	}

	// merge new interval with all the overlapping intervals
	currIndex := maxIntervalLessThanNew
	for i:=maxIntervalLessThanNew+1; i<len(intervals); i++ {
		currIndex = i
		if intervals[i][0] <= newInterval[1] {
			newInterval[0] = min(intervals[i][0], newInterval[0])
			newInterval[1] = max(intervals[i][1], newInterval[1])
			// fmt.Println("merging with overlapping interval to form", newInterval)
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
