func eraseOverlapIntervals(intervals [][]int) int {
    sort.Slice(intervals, func(i, j int)bool {
		return intervals[i][0] < intervals[j][0]
	})

	currEnd := intervals[0][1]

	i:=1
	minRemovals := 0

	for i<len(intervals) {

		if intervals[i][0] < currEnd {
			currEnd = min(currEnd, intervals[i][1])
			minRemovals++
		} else {
			currEnd = intervals[i][1]
		}

		i++

	}

	return minRemovals
}
