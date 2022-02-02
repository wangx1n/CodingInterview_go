package dynamicProgramming

func longestPalindrome(s string) string {
	var dp [][]bool
	n := len(s)
	maxLen, begin := 1, 0
	for i := 0; i < n; i++ {
		dp = append(dp, make([]bool, len(s)))
	}
	for i := 0; i < n; i++ {
		dp[i][i] = true
	}
	for j := 0; j < n; j++ {
		for i := 0; i < j; i++ {
			if s[i] == s[j] {
				if j-i < 3 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			}else {
				dp[i][j] = false
			}

			if dp[i][j] && j - i + 1 > maxLen {
				maxLen = j - i + 1
				begin = i
			}
		}
	}
	return s[begin:begin + maxLen]
}
