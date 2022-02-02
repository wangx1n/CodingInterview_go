package _02201

func construct2DArray(original []int, m int, n int) [][]int {
	length := len(original)
	if length != m*n {
		return [][]int{}
	}
	result := [][]int{}
	tmp := 0
	for i := 0; i < m; i++ {
		result = append(result, original[tmp:tmp+n])
		tmp += n
	}
	return result
}
