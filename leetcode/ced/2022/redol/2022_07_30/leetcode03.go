package _022_07_30

func lengthOfLongestSubstring(s string) int {
	n, r := len(s), -1
	maxLen := 0
	m := map[byte]int{}
	for i := 0; i < n; i++ {
		if i != 0 {
			delete(m, s[i-1])
		}
		for r+1 < n && m[s[r+1]] == 0 {
			m[s[r+1]]++
			r++
		}
		if r-i+1 > maxLen {
			maxLen = r - i + 1
		}
	}
	return maxLen
}
