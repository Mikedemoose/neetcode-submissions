func permute(nums []int) [][]int {
	if len(nums) <= 1 {
		res := [][]int{nums}
		return res
	}


	currElement := nums[0]

	permNext := permute(nums[1:])

	finalList := make([][]int, 0)

	for _, permList := range permNext {
		for j := range len(permList)+1 {
			currList := make([]int, 0)	
			currList = append(currList, permList[:j]...)
			currList = append(currList, currElement)
			currList = append(currList, permList[j:]...)
			finalList = append(finalList, currList)
		}
	}


	return finalList
}
