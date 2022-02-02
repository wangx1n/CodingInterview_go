package _02201

var moveTo = [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

func isEscapePossible(blocked [][]int, source []int, target []int) bool {
	m, n := len(blocked), len(blocked[0])
	return dfs(m, n, source[0], source[1], blocked, target)
}

func dfs(length, width, nextX, nextY int, blocked [][]int, target []int) bool {
	if nextX < 0 || nextX >= length || nextY < 0 || nextY >= width {
		return false
	}
	if nextX == target[0] && nextY == target[1] {
		return true
	}
	return dfs(length, width, nextX + moveTo[0][0], nextY, blocked, target) ||
		dfs(length, width, nextX + moveTo[1][0], nextY, blocked, target) ||
		dfs(length, width, nextX, nextY + moveTo[2][1], blocked, target) ||
		dfs(length, width, nextX, nextY + moveTo[3][1], blocked, target)
}