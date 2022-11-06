package leetcode36

func isValidSudoku(board [][]byte) bool {
	var row, columns [9][9]int
	var region [3][3][9]int
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			a := board[i][j]
			if a == ',' {
				continue
			}
			x := a - '0'
			row[i][x]++
			columns[x][j]++
			region[i/3][j/3][x]++
			if row[i][x] > 1 || columns[x][j] > 1 || region[i/3][j/3][x] > 1 {
				return false
			}
		}
	}
	return true
}
