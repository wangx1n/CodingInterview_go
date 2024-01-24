package leetcode30

func findSubstring_timeout(s string, words []string) []int {
	m := getGlobalCombination(words)
	res := []int{}

	stepSize := len(words[0]) * len(words)
	l := 0
	r := stepSize + l
	for r <= len(s) {
		t := s[l:r]
		if m[t] {
			res = append(res, l)
		}
		l++
		r++
	}
	return res
}

func getGlobalCombination(words []string) map[string]bool {
	ret := make(map[string]bool)
	used := make([]bool, len(words))
	dfs(ret, words, "", used)
	return ret
}

func dfs(ret map[string]bool, words []string, temp string, used []bool) {
	if len(temp) == len(words[0])*len(words) {
		ret[temp] = true
		return
	}
	for i := 0; i < len(words); i++ {
		if used[i] {
			continue
		}
		temp += words[i]
		used[i] = true
		dfs(ret, words, temp, used)
		used[i] = false
		temp = temp[:len(temp)-len(words[0])]
	}
}
