func combinationSum2(candidates []int, target int) [][]int {
	sort.Slice(candidates, func(i,j int)bool{
		return candidates[i] < candidates[j]
	})

	res, _ := recurse(candidates, target)
	return res
}


func recurse(candidates []int, target int) ([][]int, bool) {

	res := make([][]int, 0)


	if target == 0 {
		return [][]int{{}}, true
	} else if len(candidates) == 0 || candidates[0] > target {
		return res, false
	}

	currElement := candidates[0]
	var ok1, ok2 bool
	withRes, withoutRes := make([][]int, 0), make([][]int, 0)
	// check with currElement
	if withRes, ok1 = recurse(candidates[1:], target-currElement); ok1 {
		for _, item := range withRes {
			res = append(res, append(item, currElement))
		}
	}
	
	// check without currElement
	ind := 0
	for _, val := range candidates {
		if val != currElement {
			break
		}
		ind++
	}
	if withoutRes, ok2 = recurse(candidates[ind:], target); ok2 {
		res = append(res, withoutRes...)
	}

	return res, ok1 || ok2

}
