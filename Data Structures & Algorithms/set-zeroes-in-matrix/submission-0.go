func setZeroes(matrix [][]int) {
    m, n := len(matrix), len(matrix[0])
	indexRow := -1
	for i:=0; i<m; i++ {
		isZeroPresent := false
		for j:=0; j<n; j++ {
			if matrix[i][j] == 0 {
				isZeroPresent = true
				break
			}
		}
		if isZeroPresent {
			if indexRow == -1 {
				indexRow = i
			}
			for j:=0; j<n; j++ {
				if matrix[i][j] == 0 {
					matrix[indexRow][j] = 1
				} else {
					matrix[i][j] = 0
				}
			}
		}
	}


	if indexRow == -1 {
		return
	}


	for j:=0; j<n; j++ {
		if matrix[indexRow][j] == 1 {
			for i:=0; i<m; i++ {
				matrix[i][j] = 0
			}
		}
	}
}
