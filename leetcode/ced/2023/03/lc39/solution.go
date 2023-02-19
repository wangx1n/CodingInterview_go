package lc39

import "sort"

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var ans []int
	var res [][]int
	dfs(candidates, ans, &res, target, 0)
	return res
}

func dfs(nums, ans []int, res *[][]int, target, idx int) {
	if idx == len(nums) {
		return
	}
	if target == 0 {
		*res = append(*res, append([]int{}, ans...))
		return
	}

	dfs(nums, ans, res, target, idx+1)

	if target-nums[idx] >= 0 {
		ans = append(ans, nums[idx])
		dfs(nums, ans, res, target-nums[idx], idx)
		ans = ans[:len(ans)-1]
	}
}
