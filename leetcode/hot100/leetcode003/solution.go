package leetcode003

func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]int)
	r := 0
	n := len(s)
	maxLen := 0
	for i := 0; i < n; i++ {
		if i > 0 {
			m[s[i-1]] = 0
		}
		for r < n && m[s[r]] == 0 {
			m[s[r]]++
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
