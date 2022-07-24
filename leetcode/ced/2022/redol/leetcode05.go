package redol

func longestPalindrome(s string) string {
	length, maxLen := len(s), 1
	left := 0
	n := make([][]bool, length)
	for i := 0; i < length; i++ {
		n[i] = make([]bool, length)
	}
	for i := 0; i < length; i++ {
		n[i][i] = true
	}
	for j := 1; j < length; j++ {
		for i := 0; i < j; i++ {
			if s[i] == s[j] {
				if j-i < 3 {
					n[i][j] = true
				} else {
					n[i][j] = n[i+1][j-1]
				}
			}
			if n[i][j] {
				if j-i+1 > maxLen {
					left = i
					maxLen = j - i + 1
				}
			}
		}
	}
	res := s[left : left+maxLen]
	return res
}
