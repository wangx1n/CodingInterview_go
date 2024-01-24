package leetcode212

type pair struct{ x, y int }

var direction = []pair{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func findWords(board [][]byte, words []string) []string {
	vis := make([][]bool, len(board))
	for i := range vis {
		vis[i] = make([]bool, len(board[0]))
	}

	var ans []string

	var check func(word string) bool

	check = func(word string) bool {
		for i, row := range board {
			for j := range row {
				if dfs(board, word, i, j, 0, vis) {
					return true
				}
			}
		}
		return false
	}

	for _, word := range words {
		if check(word) {
			ans = append(ans, word)
		}
	}
	return ans
}

func dfs(board [][]byte, word string, i, j, index int, vis [][]bool) bool {
	if word[index] != board[i][j] {
		return false
	}
	if index == len(word)-1 {
		return true
	}

	vis[i][j] = true
	defer func() { vis[i][j] = false }()

	for _, dir := range direction {
		newI, newJ := i+dir.x, j+dir.y
		if newI < 0 || newI >= len(board) || newJ < 0 || newJ >= len(board[0]) || vis[newI][newJ] {
			continue
		}
		if dfs(board, word, newI, newJ, index+1, vis) {
			return true
		}
	}
	return false
}
