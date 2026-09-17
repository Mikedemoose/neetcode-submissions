func canPartition(nums []int) bool {
    sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%2 == 1 || len(nums)==1 {
		return false
	}
	target := sum/2

	memo := make(map[int]map[int]bool)

	return getSubsetWithSum(nums, target, 0, memo)
}

func getSubsetWithSum(nums []int, target int, index int, memo map[int]map[int]bool) bool {
	if len(nums) == 0 || index >= len(nums){
		return false
	}
	if m, ok := memo[target]; ok {
		if val, ok1 := m[index]; ok1 {
			return val
		}
	}
	currElement := nums[index]
	var res bool
	// check with currElement
	if currElement == target {
		res = true
	} else if currElement < target {
		if ok1 := getSubsetWithSum(nums, target-currElement, index+1, memo); ok1 {
			res = true
		}
	}
	// check without currElement
	if !res {
		if ok2 := getSubsetWithSum(nums, target, index+1, memo); ok2 {
			res = true
		}
	}

	if _, ok := memo[target]; ok {
		memo[target][index] = res
	} else {
		memo[target] = map[int]bool {index: res}
	}
	return res
}
