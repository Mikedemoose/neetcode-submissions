type A struct {
	value int
	index int
}

func twoSum(nums []int, target int) []int {
	newVals := make([]A, len(nums))
	for i, val := range nums {
		newVals[i] = A {
			value: val,
			index: i,
		}
	}
    left, right := 0, len(nums)-1
	sort.Slice(newVals, func(i, j int) bool {
		return newVals[i].value < newVals[j].value
	})

	for left < right {
		sum := newVals[left].value + newVals[right].value

		if sum == target {
			break
		} else if sum < target {
			left++
		} else {
			right--
		}
	}

	if newVals[left].index < newVals[right].index {
		return []int{newVals[left].index, newVals[right].index}
	}
	return []int{newVals[right].index, newVals[left].index}
}
