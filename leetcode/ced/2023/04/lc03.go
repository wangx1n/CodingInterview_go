package _4

// 1. 右指针向右移动，直到遇到重复字符
// 2. 左指针向右移动，直到没有重复字符
func lengthOfLongestSubstring(s string) int {
	n := len(s)
	r := 0
	m := map[byte]int{}
	maxLen := 0
	for i := 0; i < n; i++ {
		if i > 0 {
			m[s[i-1]] = 0
		}
		for r < n && m[s[r]] < 1 {
			m[s[r+1]]++
			r++
		}
		maxLen = max(maxLen, r-i)
	}
	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
