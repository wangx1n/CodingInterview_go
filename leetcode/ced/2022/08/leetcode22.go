package _8

func generateParenthesis(n int) []string {
	ans := []byte{}
	var res_ []string
	backTrack(&res_, ans, 0, 0, n)
	return res_
}

func backTrack(res_ *[]string, ans []byte, open, close, max int) {
	if len(ans) == max*2 {
		*res_ = append(*res_, string(ans))
		return
	}

	if open < max {
		ans = append(ans, '(')
		backTrack(res_, ans, open+1, close, max)
		ans = ans[:len(ans)-1]
	}

	if close < open {
		ans = append(ans, ')')
		backTrack(res_, ans, open, close+1, max)
		ans = ans[:len(ans)-1]
	}
}
