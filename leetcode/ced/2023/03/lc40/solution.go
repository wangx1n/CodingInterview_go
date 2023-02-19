package lc40

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	var ans []int
	var res [][]int
	var used []bool
	used = make([]bool, len(candidates))

	sort.Ints(candidates)

	dfs(candidates, ans, &res, used, target, 0)
	return res
}

func dfs(nums, ans []int, res *[][]int, used []bool, target, begin int) {
	if target == 0 {
		*res = append(*res, append([]int{}, ans...))
		return
	}

	for i := begin; i < len(nums); i++ {
		if target-nums[i] >= 0 {
			if i > begin && nums[i] == nums[i-1] && !used[i] {
				continue
			}
			used[i] = true
			ans = append(ans, nums[i])
			dfs(nums, ans, res, used, target-nums[i], i+1)
			used[i] = false
			ans = ans[:len(ans)-1]
		}
	}
}
