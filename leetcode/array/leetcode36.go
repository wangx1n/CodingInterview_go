package main

func isValidSudoku(board [][]byte) bool {
	var row, col [9][9]int
	var rigion [3][3][9]int
	for i, r := range board {
		for j, c := range r {
			if c == '.' {
				continue
			}
			btoi := c - '1'
			row[i][btoi] += 1
			col[j][btoi] += 1
			rigion[i/3][j/3][btoi] += 1
			if row[i][btoi] > 1 || col[j][btoi] > 1 || rigion[i/3][j/3][btoi] > 1 {
				return false
			}
		}
	}
	return true
}
