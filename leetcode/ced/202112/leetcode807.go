package _02112

func maxIncreaseKeepingSkyline(grid [][]int) int {
	ans := 0
	n := len(grid)
	maxRow := make([]int, n)
	maxCol := make([]int, n)
	for i, row := range grid {
		for j, h := range row {
			maxRow[i] = max(maxRow[i], h)
			maxCol[j] = max(maxCol[j], h)
		}
	}
	for i, row := range grid {
		for j, h := range row {
			ans += min(maxRow[i], maxCol[j]) - h
		}
	}
	return ans
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func min(i, j int) int {
	if i > j {
		return j
	}
	return i
}