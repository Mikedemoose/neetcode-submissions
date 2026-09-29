func solveNQueens(n int) [][]string {
	
	res := make([][]string, 0)
	var recurse func([]int, int)

	recurse = func(currArr []int, n int) {
		if len(currArr) == n {
			// fmt.Println("reached terminal state. Adding", currArr, "to res")
			newRes := make([]string, 0)
			for _, num := range currArr {
				newRes = append(newRes, convertNumToString(num, n))
			}
			res = append(res, newRes)
			// fmt.Println("curr res:", res)
			return
		}
		for i := range n {
			isValid := true
			for u, v := range currArr {
				// check if column is occupied
				if v == i {
					isValid = false
					break
				}
				// check if diagonal is occupied
				if sqr(v-i) == sqr(len(currArr)-u) {
					isValid = false
					break
				}
			}
			if isValid {
				newarr := append(append([]int{}, currArr...), i)
				// fmt.Println("isValid is true. Current state:", newarr)
				recurse(newarr, n)
			}
		}
	}

	recurse(make([]int, 0), n)

	return res
}

func sqr(a int) int {
	return a*a
}

func convertNumToString(num, totalLen int) string {
	bytArr := make([]byte, totalLen)
	for i := range totalLen {
		if num == i {
			bytArr[i] = 'Q'
		} else {
			bytArr[i] = '.'
		}
	}
	return string(bytArr)
}
