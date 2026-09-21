func findTargetSumWays(nums []int, target int) int {
	memo := make(map[MemoKey]int)
	finalVal := recurse(nums, target, 0, memo)
	return finalVal
}

func recurse(nums []int, target int, start int, memo map[MemoKey]int) int {

	if start == len(nums) {
		if target == 0 {
			return 1
		} 
		return 0
	}

	if val, ok := memo[MemoKey{start: start, target: target}]; ok {
		return val
	}

	currVal := nums[start]
	finalVal := recurse(nums, target-currVal, start+1, memo) + recurse(nums, target+currVal, start+1, memo)

	memo[MemoKey{start: start, target: target}] = finalVal 

	return finalVal

}

type MemoKey struct {
	start int
	target int
}