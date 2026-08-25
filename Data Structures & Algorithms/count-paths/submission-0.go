func uniquePaths(m int, n int) int {
	if m==1 || n==1 {
		return 1
	}

    res := make([][]int, m)
	for i := range m {
		res[i] = make([]int, n)
	}

	f := func(i, j int) int {
		if i == 0 || j == 0 {
			return 1
		}
		return res[i][j]
	}

	for i:=1; i<m; i++ {
		for j:=1; j<n; j++ {
			res[i][j] = f(i-1,j) + f(i,j-1)
		}
	}

	return res[m-1][n-1]
}
