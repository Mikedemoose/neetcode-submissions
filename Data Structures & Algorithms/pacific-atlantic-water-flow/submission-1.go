func pacificAtlantic(heights [][]int) [][]int {
    m, n := len(heights), len(heights[0])

	visited := make([][]byte, m)
	visited1 := make([][]byte, m)
	for i := range m {
		visited[i] = make([]byte, n)
		visited1[i] = make([]byte, n)
	}

	for j:=0; j<n; j++ {
		pacificFill(heights, visited, 0, j, m, n)
		atlanticFill(heights, visited1, m-1, j, m, n)
	}
	for i := range m {
		pacificFill(heights, visited, i, 0, m, n)
		atlanticFill(heights, visited1, i, n-1, m, n)
	}

	finalList := make([][]int, 0)
	for i:= range m {
		for j:= range n {
			if visited[i][j] == 'p' && visited1[i][j] == 'p' {
				finalList = append(finalList, []int{i, j})
			}
		}
	}

	return finalList
}


func pacificFill(heights [][]int, visited [][]byte, i, j int, m, n int) {
	if i < 0 || j < 0 || i >= m || j >= n || visited[i][j] == 'p' {
		return
	}

	visited[i][j] = 'p'

	if i < m-1 && heights[i][j] <= heights[i+1][j]{
		pacificFill(heights, visited, i+1, j, m, n)
	} 
	if j < n-1 && heights[i][j] <= heights[i][j+1] {
		pacificFill(heights, visited, i, j+1, m, n)
	} 
	if i > 0 && heights[i][j] <= heights[i-1][j] {
		pacificFill(heights, visited, i-1, j, m, n)
	} 
	if j > 0 && heights[i][j] <= heights[i][j-1] {
		pacificFill(heights, visited, i, j-1, m, n)
	} 

}


func atlanticFill(heights [][]int, visited [][]byte, i, j int, m, n int) {
	if i < 0 || j < 0 || i >= m || j >= n || visited[i][j] == 'p' {
		return
	}

	visited[i][j] = 'p'

	if i < m-1 && heights[i][j] <= heights[i+1][j] {
		atlanticFill(heights, visited, i+1, j, m, n)
	} 
	if j < n-1 && heights[i][j] <= heights[i][j+1] {
		atlanticFill(heights, visited, i, j+1, m, n)
	} 
	if i > 0 && heights[i][j] <= heights[i-1][j] {
		atlanticFill(heights, visited, i-1, j, m, n)
	} 
	if j > 0 && heights[i][j] <= heights[i][j-1] {
		atlanticFill(heights, visited, i, j-1, m, n)
	} 

}