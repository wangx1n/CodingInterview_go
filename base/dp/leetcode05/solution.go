package leetcode05

func longestPalindrome(s string) string {
	m := len(s)
	dp := make([][]bool, m)
	for i := range dp {
		dp[i] = make([]bool, m)
	}

	for i := 0; i < m; i++ {
		dp[i][i] = true
	}

	begin := 0
	maxLen := 1
	for j := 1; j < m; j++ {
		for i := 0; i < j; i++ {
			if s[i] == s[j] {
				if j-1 == i {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}

			}
			if dp[i][j] && j-i+1 > maxLen {
				maxLen = j - i + 1
				begin = i
			}
		}
	}
	return s[begin : begin+maxLen]
}
