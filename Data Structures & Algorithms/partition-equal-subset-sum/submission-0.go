func canPartition(nums []int) bool {
    sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%2 == 1 || len(nums)==1 {
		return false
	}
	target := sum/2

	return getSubsetWithSum(nums, target)
}

func getSubsetWithSum(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
	}
	currElement := nums[0]
	// check with currElement
	if currElement == target {
		return true
	} else if currElement < target {
		if ok1 := getSubsetWithSum(nums[1:], target-currElement); ok1 {
			return true
		}
	}
	// check without currElement
	if ok2 := getSubsetWithSum(nums[1:], target); ok2 {
		return true
	}

	return false
}
