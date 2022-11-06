package lc47

import "sort"

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	var ans []int
	var res [][]int
	dfs(nums, ans, &res, 0)
	return res
}

func dfs(nums, ans []int, res *[][]int, count int) {
	if len(ans) == len(nums) {
		*res = append(*res, append([]int{}, ans...))
	}
	for i := count; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		ans = append(ans, nums[i])
		dfs(nums, ans, res, i+1)
		ans = ans[:len(ans)-1]
	}
}
