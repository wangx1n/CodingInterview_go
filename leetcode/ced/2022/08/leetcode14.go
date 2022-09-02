package _8

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	var findMid func(start, end int) string
	findMid = func(start, end int) string {
		if start == end {
			return strs[start]
		}
		mid := (start + end) / 2
		left, right := findMid(start, mid), findMid(mid+1, end)
		if len(left) == 0 || len(right) == 0 {
			return ""
		}
		minLen := min(len(left), len(right))
		for i := 0; i < minLen; i++ {
			if left[i] != right[i] {
				return left[:i]
			}
		}
		return left[:minLen]
	}

	return findMid(0, len(strs)-1)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
