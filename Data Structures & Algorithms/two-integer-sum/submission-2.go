func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)

	for i, val := range nums {
		if ind, ok := numMap[target-val]; ok {
			return []int {ind, i}
		}
		numMap[val] = i
	}

	return nil
}
