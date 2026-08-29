func rotate(matrix [][]int)  {

	n := len(matrix)
	
	rotateOne := func(i, j int) {
		temp := matrix[i][j]
		matrix[i][j] = matrix[n-1-j][i]
		matrix[n-1-j][i] = matrix[n-1-i][n-1-j]
		matrix[n-1-i][n-1-j] = matrix[j][n-1-i]
		matrix[j][n-1-i] = temp
	}


	left, right := 0, n-1

	for left < right {
		numRotations := right-left

		for i:=0; i<numRotations; i++ {
			rotateOne(left, left+i)
		}
		left++
		right--
	}

}
