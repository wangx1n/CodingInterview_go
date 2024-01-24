package leetcode003

func lengthOfLongestSubstring_1002(s string) int {
	m := make(map[byte]bool, len(s))

	r := -1
	maxLen := 0
	for i := 0; i < len(s); i++ {
		if i > 0 {
			m[s[i-1]] = false
		}
		for r < len(s) {
			r++
			if m[s[r]] {
				break
			}
			m[s[r]] = true
			tmpMax := r - i + 1
			if tmpMax > maxLen {
				maxLen = tmpMax
			}
		}
	}
	return maxLen
}
