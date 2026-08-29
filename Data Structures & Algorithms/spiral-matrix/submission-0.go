func spiralOrder(matrix [][]int) []int {
    
	m, n := len(matrix), len(matrix[0])

	top, bottom := 0, m-1
	left, right := 0, n-1

	finalList := make([]int, 0)


	for left <= right && top <= bottom {
		for j:=left; j<=right; j++ {
			finalList = append(finalList, matrix[top][j])
		}
		top++
		if top>bottom {
			break
		}

		for i:=top; i<=bottom; i++ {
			finalList = append(finalList, matrix[i][right])
		}
		right--
		if left>right {
			break
		}

		for j:=right; j>=left; j-- {
			finalList = append(finalList, matrix[bottom][j])
		}
		bottom--
		if top>bottom {
			break
		}

		for i:=bottom; i>=top; i-- {
			finalList = append(finalList, matrix[i][left])
		}
		left++
		if left>right{
			break
		}
	}

	return finalList

}
