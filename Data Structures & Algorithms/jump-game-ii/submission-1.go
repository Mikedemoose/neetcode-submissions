func jump(nums []int) int {
	var recurse func(int)int
	memo := make(map[int]int)

	recurse = func (start int) int {
		if start >= len(nums)-1 {
			return 0
		}

		if val, ok := memo[start]; ok {
			return val
		}

		var res int
		currVal := nums[start]
		if start + currVal >= len(nums)-1 {
			res = 1
		} else if currVal == 0 {
			res = -1
		} else {
			val1 := -1
			i := 1
			for val1 == -1 && i <= currVal{
				val1 = recurse(start+i) 
				i++
			}

			for i <= currVal {
				temp := recurse(start+i)
				if temp != -1 {
					val1 = min(temp, val1)
				}
				i++
			}
			res = val1+1
		}
		memo[start] = res
		return res
	}

    return recurse(0)
}
