package leetcode022

func generateParenthesis(n int) []string {
	var ret []string
	dfs(&ret, "", 0, 0, n)
	return ret
}

func dfs(ret *[]string, ans string, open, close int, max int) {
	if len(ans) == max*2 {
		*ret = append(*ret, ans)
		return
	}
	if open < max {
		ans = ans + "("
		dfs(ret, ans, open+1, close, max)
		ans = ans[:len(ans)-1]
	}
	if close < open {
		ans = ans + ")"
		dfs(ret, ans, open, close+1, max)
		ans = ans[:len(ans)-1]
	}
}
