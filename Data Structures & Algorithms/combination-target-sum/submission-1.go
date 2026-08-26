func dfs(nums []int, target int, start int) [][]int {

	currRes := make([][]int, 0)

	for i:=start; i<len(nums); i++ {
		num := nums[i]
		if num == target {
			currRes = append(currRes, []int{num})
		} else if num < target {
			nextRes := dfs(nums, target-num, i)
			for _, item := range nextRes {
				currRes = append(currRes, append(item, num))
			}
		}
	}

	return currRes


}

func combinationSum(nums []int, target int) [][]int {
	return dfs(nums, target, 0)
}
