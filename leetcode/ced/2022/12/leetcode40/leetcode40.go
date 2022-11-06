package leetcode40

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	var ans [][]int
	var t_ans []int
	sort.Ints(candidates)
	dfs(candidates, t_ans, &ans, target, 0, 0)
	return ans
}

func dfs(cand, t_ans []int, ans *[][]int, target, sum, begin int) {
	if sum == target {
		*ans = append(*ans, append([]int{}, t_ans...))
		return
	}
	for i := begin; i < len(cand); i++ {
		if i > begin && cand[i] == cand[i-1] {
			continue
		}
		s := sum + cand[i]
		if s <= target {
			t_ans = append(t_ans, cand[i])
			dfs(cand, t_ans, ans, target, s, i+1)
			t_ans = t_ans[:len(t_ans)-1]
		} else {
			break
		}
	}
}
