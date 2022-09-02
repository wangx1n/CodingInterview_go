package main

func main() {
	isMatch("abbbbc", "ab*d*c")
}

func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	matches := func(i, j int) bool {
		if i == 0 {
			return false
		}
		if p[j-1] == '.' {
			return true
		}
		return s[i-1] == p[j-1]
	}

	f := make([][]bool, m+1)
	for i := 0; i < len(f); i++ {
		f[i] = make([]bool, n+1)
	}
	f[0][0] = true
	for i := 0; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == '*' {
				// *组合用0次
				f[i][j] = f[i][j-2]
				// 如果*前面的字符与s[i]相等，那么使用一次，i往前走一个，看用2次3次的情况
				if matches(i, j-1) {
					f[i][j] = f[i][j] || f[i-1][j]
				}
				// 如果s[i]==s[j]，那么...
			} else if matches(i, j) {
				f[i][j] = f[i][j] || f[i-1][j-1]
			}
		}
	}
	return f[m][n]
}
