func isValidSudoku(board [][]byte) bool {
	// check for each row
	for i:=0; i<9; i++ {
		rowMap := make(map[byte]struct{})
		for j := 0; j<9; j++ {
			if board[i][j] == '.' {
				continue
			}
			if _, ok := rowMap[board[i][j]]; ok {
				return false
			}
			rowMap[board[i][j]] = struct{}{}
		}
	}
	// check for each column
	for i:=0; i<9; i++ {
		colMap := make(map[byte]struct{})
		for j := 0; j<9; j++ {
			if board[j][i] == '.' {
				continue
			}
			if _, ok := colMap[board[j][i]]; ok {
				return false
			}
			colMap[board[j][i]] = struct{}{}
		}
	}

	// check for each 9x9 grid
	for i:= range 3 {
		for j:= range 3 {

			gridMap := make(map[byte]struct{})

			row_start := 3*i
			col_start := 3*j

			for i_new := range 3 {

				for j_new := range 3 {

					u, v := row_start + i_new, col_start+j_new
					if board[u][v] == '.' {
						continue
					}

					if _, ok := gridMap[board[u][v]]; ok {
						return false
					}
					gridMap[board[u][v]] = struct{}{}
				}
			}
		}
	}

	return true
}
