func maxAreaOfIsland(grid [][]int) int {
    m, n := len(grid), len(grid[0])
	maxArea := 0
	for i := range m {
		for j := range n {
			if grid[i][j] == 1 {
				maxArea = max(maxArea, findIslandArea(grid, i, j, m, n))
			}
		}
	}

	return maxArea
}


func findIslandArea(grid [][]int, i, j int, m, n int) int {

	if i<0 || j <0 || i >= m || j >= n || grid[i][j] != 1 {
		return 0
	}
	grid[i][j] = 0
	
	return 1 + findIslandArea(grid, i+1, j, m, n) + findIslandArea(grid, i, j+1, m, n) + findIslandArea(grid, i-1, j, m, n) + findIslandArea(grid, i, j-1, m, n)

}
