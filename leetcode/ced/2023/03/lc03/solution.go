package lc03

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	m := map[byte]int{}
	left := 0
	right := 0
	maxLen := 1
	m[s[left]] = 1
	for right < len(s) {
		if m[s[right]] == 0 {
			m[s[right]]++
			if (right - left + 1) > maxLen {
				maxLen = right - left + 1
			}
			right++
		} else {
			m[s[left]] = 1
			left = left + 1
		}
	}
	return maxLen
}
