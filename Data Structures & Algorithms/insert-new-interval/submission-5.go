func insert(intervals [][]int, newInterval []int) [][]int {
    if len(intervals) == 0 {
		return [][]int{newInterval}
	}

	finalList := make([][]int, 0)

	// find the highest interval which is less than new interval
	var i int
	for i<len(intervals) {
		if intervals[i][1] >= newInterval[0] {
			break
		} 
		finalList = append(finalList, intervals[i])
		i++
	}

	// merge new interval with all the overlapping intervals
	for i<len(intervals) {
		if intervals[i][0] <= newInterval[1] {
			newInterval[0] = min(intervals[i][0], newInterval[0])
			newInterval[1] = max(intervals[i][1], newInterval[1])
		} else {
			break
		}
		i++
	}

	// add new interval and then the rest of the intervals
	finalList = append(finalList, newInterval)
	finalList = append(finalList, intervals[i:]...)

	return finalList
}
