package leetcode040

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	var ret [][]int
	sort.Ints(candidates)
	dfs(&ret, []int{}, target, 0, make([]bool, len(candidates)), candidates)
	return ret
}

func dfs(ret *[][]int, ans []int, target int, index int, used []bool, candidates []int) {
	if target == 0 {
		*ret = append(*ret, append([]int{}, ans...))
		return
	}

	for i := index; i < len(candidates); i++ {
		if used[i] {
			continue
		}
		if target-candidates[i] >= 0 {
			ans = append(ans, candidates[i])
			used[i] = true
			dfs(ret, ans, target-candidates[i], i+1, used, candidates)
			used[i] = false
			ans = ans[:len(ans)-1]
		}
	}
}
