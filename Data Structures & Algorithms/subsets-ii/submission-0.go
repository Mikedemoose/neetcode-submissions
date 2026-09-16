func subsetsWithDup(nums []int) [][]int {
	sort.Slice(nums, func(i, j int)bool{
		return nums[i] < nums[j]
	})

	return recurse(nums)
}

func recurse(nums []int) [][]int {
	res := [][]int{{}}
	if len(nums) == 0 {
		return res
	}


	currElement := nums[0]
	// check subsets with the element
	withElement := recurse(nums[1:])
	for _, item := range withElement {
		res = append(res, append([]int{currElement}, item...))
	}


	// check subsets without the element
	newNums := append([]int{}, nums...)
	for len(newNums) > 0 && newNums[0]==currElement {
		newNums = newNums[1:]
	}
	withoutElement := recurse(newNums)
	for _, item := range withoutElement {
		if len(item) > 0 {
			res = append(res, append([]int{}, item...))
		}
	}

	return res

}
