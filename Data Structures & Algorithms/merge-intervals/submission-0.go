func merge(intervals [][]int) [][]int {
    sort.Slice(intervals, func(i, j int) bool{
		return intervals[i][0] < intervals[j][0]
	})


	finalList := make([][]int, 0)

	currStart, currEnd := intervals[0][0], intervals[0][1]
	i:= 1
	for i<len(intervals) {

		if intervals[i][0] <= currEnd {
			currStart = min(intervals[i][0], currStart)
			currEnd = max(intervals[i][1], currEnd)
		} else {
			finalList = append(finalList, []int{currStart, currEnd})
			currStart, currEnd = intervals[i][0], intervals[i][1]
		}
		i++
	}

	finalList = append(finalList, []int{currStart, currEnd})
	return finalList
}
