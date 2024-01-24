package leetcode079

type pair struct{ x, y int }

var direction = []pair{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func exist(board [][]byte, word string) bool {
	vis := make([][]bool, len(board))
	for i := range vis {
		vis[i] = make([]bool, len(board[0]))
	}
	for i, row := range board {
		for j := range row {
			if dfs(i, j, 0, vis, board, word) {
				return true
			}
		}
	}
	return false
}

func dfs(i, j int, index int, visited [][]bool, board [][]byte, word string) bool {
	if board[i][j] != word[index] {
		return false
	}
	if index == len(word)-1 {
		return true
	}

	visited[i][j] = true
	defer func() { visited[i][j] = false }()

	for _, dir := range direction {
		if newI, newJ := i+dir.x, j+dir.y; newI >= 0 && newI < len(board) && newJ >= 0 && newJ < len(board[0]) {
			if visited[newI][newJ] {
				continue
			}
			if dfs(newI, newJ, index+1, visited, board, word) {
				return true
			}
		}
	}
	return false
}
