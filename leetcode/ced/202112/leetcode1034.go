package _02112

func colorBorder(grid [][]int, row int, col int, color int) [][]int {
	var rot = []struct{x ,y int}{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}
	m, n := len(grid), len(grid[0])
	startPositionColor := grid[row][col]
	type point struct{ x, y int }
	borders := []point{}
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	var dfs func(x, y int)
	dfs = func(x, y int) {
		visited[x][y] = true
		isBoard := false
		for _, v := range rot {
			nx, ny := x + v.x, y + v.y
			if ! (nx >= 0 && nx < m && ny >= 0 && ny < n && grid[nx][ny] == startPositionColor) {
				isBoard = true
			} else if !visited[nx][ny] {
				dfs(nx, ny)
			}
		}
		if isBoard {
			borders = append(borders, point{x, y})
		}
	}

	dfs(row, col)
	for _, v := range borders {
		x, y := v.x, v.y
		grid[x][y] = color
	}
	return grid
}