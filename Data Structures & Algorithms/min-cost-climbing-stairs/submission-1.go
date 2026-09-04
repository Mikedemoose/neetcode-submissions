func minCostClimbingStairs(cost []int) int {
	return dfs(cost, make(map[int]int), 0)
}

func dfs(cost []int, memo map[int]int, head int) int {
	if head >= len(cost) {
		return 0
	}

	if val, ok := memo[head]; ok {
		return val
	}

	min1 := cost[head] + dfs(cost, memo, head+1)
	min2 := 0
	if head+1 < len(cost) {
		min2 = cost[head+1] + dfs(cost, memo, head+2)
	}


	minVal := min(min1, min2)
	memo[head] = minVal

	return minVal
}
