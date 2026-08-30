func orangesRotting(grid [][]int) int {
    // find all the rotten fruits
	rottenFruitIndices := make([][2]int, 0)
	m, n := len(grid), len(grid[0])

	numOfFruits := 0
	for i := range m {
		for j := range n {
			if grid[i][j] == 2 {
				rottenFruitIndices = append(rottenFruitIndices, [2]int{i, j})
			} else if grid[i][j] == 1 {
				numOfFruits++
			}
		}
	}
	maxTime := 0

	directions := [][2]int{
		{-1, 0},
		{0, -1},
		{1, 0},
		{0, 1},
	}

	for len(rottenFruitIndices) > 0 {
		newIndices := make([][2]int, 0)
		for _, ind := range rottenFruitIndices {
			for _, dir := range directions {
				new_i, new_j := ind[0]+dir[0], ind[1]+dir[1]
				if new_i >= 0 && new_j >= 0 && new_i < m && new_j < n && grid[new_i][new_j] == 1 {
					grid[new_i][new_j] = 2
					newIndices = append(newIndices, [2]int{new_i, new_j})
				}
			}
		}
		if len(newIndices) == 0 {
			break
		}
		numOfFruits -= len(newIndices)
		rottenFruitIndices = newIndices
		maxTime++
	}

	if numOfFruits == 0 {
		return maxTime
	}
	return -1
}
