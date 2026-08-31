func searchMatrix(matrix [][]int, target int) bool {

	m, n := len(matrix), len(matrix[0])
	rowNum := -1

	for i:= range m {
		if matrix[i][0] == target {
			return true
		} else if matrix[i][0] > target {
			rowNum = i-1
			break
		}
		rowNum = i
	}

	if rowNum < 0 {
		return false
	}

	left, right := 1, n-1
	for left <= right {
		mid := (left+right)/2
		if matrix[rowNum][mid] == target {
			return true
		} else if matrix[rowNum][mid] < target {
			left = mid+1
		} else {
			right = mid-1
		}
	}

	return false

}
