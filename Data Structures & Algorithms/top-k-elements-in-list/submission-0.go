func topKFrequent(nums []int, k int) []int {
	freqMap := make(map[int]int) 

	for _, num := range nums {
		if freq, ok := freqMap[num]; ok {
			freqMap[num] = freq+1
		} else {
			freqMap[num] = 1
		}
	}

	items := make([][]int, 0)

	for k, v := range freqMap {
		items = append(items, []int{k, v})
	}

	sort.Slice(items, func(i, j int)bool {
		return items[i][1] > items[j][1]
	})

	result := make([]int, k)

	for i := range k {
		result[i] = items[i][0]
	}

	return result
}
