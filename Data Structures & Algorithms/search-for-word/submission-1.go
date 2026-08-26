func exist(board [][]byte, word string) bool {
	l, w := len(board), len(board[0])

	var dfs func(int, int, string) bool

	dfs = func(i, j int, word string) bool {
		if len(word) == 0 {
			return true
		}
		if i > l-1 || j > w-1 || i < 0 || j < 0 {
			return false
		}

		letter := word[0]
		word = word[1:]

		isMatch := false

		if board[i][j] == letter {

			board[i][j] = '-'

			isMatch = dfs(i-1, j, word) || dfs(i+1, j, word) || dfs(i, j-1, word) || dfs(i, j+1, word)

			board[i][j] = letter

		}

		return isMatch

	}

	for i:=0; i<l; i++ {
		for j:=0; j<w; j++ {
			if dfs(i, j, word) {
				return true
			}
		}
	}

	return false

}
