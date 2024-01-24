package leetcode039

import "sort"

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var ret [][]int
	var ans []int
	dfs(&ret, candidates, target, 0, ans)
	return ret
}

func dfs(ret *[][]int, candidates []int, target int, index int, ans []int) {
	if target < 0 || index >= len(candidates) {
		return
	}

	if target == 0 {
		*ret = append(*ret, append([]int{}, ans...))
		return
	}

	dfs(ret, candidates, target, index+1, ans)

	if target-candidates[index] >= 0 {
		ans = append(ans, candidates[index])
		dfs(ret, candidates, target-candidates[index], index, ans)
		ans = ans[:len(ans)-1]
	}
}
