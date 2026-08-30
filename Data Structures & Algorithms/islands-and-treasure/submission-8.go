func islandsAndTreasure(grid [][]int) {
    m, n := len(grid), len(grid[0])
	INF := 2147483647
	queue := make([][2]int, 0)

	for i := range m {
		for j := range n {
			if grid[i][j] == 0 {
				queue = append(queue, [2]int{i, j})
			}
		}
	}

	dirs := [][2] int {
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	}

	currDist := 0

	for len(queue) > 0 {
		newQueue := make([][2]int, 0)
		for _, item := range queue {
			for _, dir := range dirs {
				u, v := item[0]+dir[0], item[1]+dir[1]

				if u>=0 && v>=0 && u<m && v<n && grid[u][v] == INF {
					grid[u][v] = currDist+1
					newQueue = append(newQueue, [2]int{u, v})
				}
			}
		}
		currDist++
		queue = newQueue
	}
}
