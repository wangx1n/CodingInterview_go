package _1

func generateParenthesis(n int) []string {
	ret := make([]string, 0, n)
	dfs(n, 0, 0, &ret, "")
	return ret
}

func dfs(target, left, right int, ret *[]string, res string) {
	if len(res) == target*2 {
		*ret = append(*ret, res)
		return
	}

	if left < target {
		res = res + "("
		dfs(target, left+1, right, ret, res)
		res = res[:len(res)-1]
	}

	if right < left {
		res = res + ")"
		dfs(target, left, right+1, ret, res)
		res = res[:len(res)-1]
	}
}
