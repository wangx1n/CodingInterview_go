package Matrix

// GetCol 获取矩阵的第j列
func GetCol (matrix [][]int, j int) (col []int) {
	for row := 0; row < len(matrix); row++ {
		col = append(col, matrix[row][j: j + 1][0])
	}
	return
}
