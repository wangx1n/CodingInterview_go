package leetcode76

import (
	"math"
	"strings"
)

func minWindow(s string, t string) string {
	maxLenSubString := ""
	maxLen := math.MaxInt
	for i := 0; i < len(s); i++ {
		if !contains(rune(s[i]), t) {
			continue
		}
		m := make(map[rune]int, len(t))
		for _, tt := range t {
			m[tt]++
		}
		for r := i; r < len(s); r++ {
			if len(s[i:r+1]) > maxLen {
				break
			}
			if contains(rune(s[r]), t) {
				m[rune(s[r])]--
				if m[rune(s[r])] <= 0 {
					delete(m, rune(s[r]))
				}
			}
			if len(m) == 0 {
				if maxLen > len(s[i:r+1]) {
					maxLenSubString = s[i : r+1]
					maxLen = len(s[i : r+1])
					break
				}
			}
		}
	}
	return maxLenSubString
}

func contains(ss rune, t string) bool {
	return strings.ContainsRune(t, ss)
}
