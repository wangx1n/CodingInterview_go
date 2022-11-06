package lc39

import "sort"

func combinationSum(candidates []int, target int) (res [][]int) {
	var tres []int
	sort.Ints(candidates)
	dfs(candidates, tres, &res, target, 0)
	return
}

func dfs(cand, tres []int, res *[][]int, target, idx int) {
	if idx == len(cand) {
		return
	}
	if target == 0 {
		*res = append(*res, append([]int{}, tres...))
		return
	}

	dfs(cand, tres, res, target, idx+1)
	if target-cand[idx] >= 0 {
		tres = append(tres, cand[idx])
		dfs(cand, tres, res, target-cand[idx], idx)
		tres = tres[:len(tres)-1]
	}
}
