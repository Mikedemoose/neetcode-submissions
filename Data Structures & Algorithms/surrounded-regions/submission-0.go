func solve(board [][]byte) {
	m, n := len(board), len(board[0])
	if m == 1 || n == 1 {
		return
	}
    
	edgeOQueue := make([][2]int, 0)

	for j := range n {
		if board[0][j] == 'O' {
			board[0][j] = 'Y'
			edgeOQueue = append(edgeOQueue, [2]int{0, j})
		}
		if board[m-1][j] == 'O' {
			board[m-1][j] = 'Y'
			edgeOQueue = append(edgeOQueue, [2]int{m-1, j})
		}
	}

	for i:=1; i<m-1; i++ {
		if board[i][0] == 'O' {
			board[i][0] = 'Y'
			edgeOQueue = append(edgeOQueue, [2]int{i, 0})
		}
		if board[i][n-1] == 'O' {
			board[i][n-1] = 'Y'
			edgeOQueue = append(edgeOQueue, [2]int{i, n-1})
		}
	}

	dirs := [][2]int{
		{1, 0},
		{0, 1},
		{-1, 0},
		{0, -1},
	}

	for len(edgeOQueue) > 0 {
		newQueue := make([][2]int, 0)
		for _, item := range edgeOQueue {
			for _, dir := range dirs {
				new_i, new_j := item[0]+dir[0], item[1]+dir[1]
				if new_i >= 0 && new_j >= 0 && new_i < m && new_j < n && board[new_i][new_j] == 'O' {
					board[new_i][new_j] = 'Y'
					newQueue = append(newQueue, [2]int{new_i, new_j})
				}
			}
		}
		edgeOQueue = newQueue
	}


	for i := range m {
		for j := range n {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			} else if board[i][j] == 'Y' {
				board[i][j] = 'O'
			}
		}
	}
}
