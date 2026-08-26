
func markVisited(i, j int, i_max, j_max int, grid [][]byte, visited *[][]bool) {
	if i<0 || j<0 || i>i_max || j>j_max || (*visited)[i][j]{
		return
	}
	(*visited)[i][j] = true

	if i>0 && grid[i-1][j] == '1' {
		markVisited(i-1, j, i_max, j_max, grid, visited)
	}
	if j>0 && grid[i][j-1] == '1' {
		markVisited(i, j-1, i_max, j_max, grid, visited)
	}
	if i<i_max && grid[i+1][j] == '1' {
		markVisited(i+1, j, i_max, j_max, grid, visited)
	}
	if j<j_max && grid[i][j+1] == '1' {
		markVisited(i, j+1, i_max, j_max, grid, visited)
	}
}

func numIslands(grid [][]byte) int {
    visited := make([][]bool, len(grid))
	for i := range len(grid) {
		visited[i] = make([]bool, len(grid[0]))
	}

	numIslands := 0

	for i:=0; i<len(grid); i++ {
		for j:=0; j<len(grid[0]); j++ {
			if grid[i][j] == '1' && !visited[i][j] {
				numIslands++
				markVisited(i,j, len(grid)-1, len(grid[0])-1, grid, &visited)
			}
		}
	}

	return numIslands
}
