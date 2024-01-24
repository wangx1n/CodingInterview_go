package leetcode047

import "sort"

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	var ret [][]int
	dfs(nums, []int{}, &ret, make([]bool, len(nums)))
	return ret
}

func dfs(nums, ans []int, ret *[][]int, used []bool) {
	if len(ans) == len(nums) {
		*ret = append(*ret, append([]int{}, ans...))
		return
	}

	for i := 0; i < len(nums); i++ {
		if !used[i] {
			if i > 0 && !used[i-1] && nums[i] == nums[i-1] {
				continue
			}
			used[i] = true
			ans = append(ans, nums[i])
			dfs(nums, ans, ret, used)
			ans = ans[:len(ans)-1]
			used[i] = false
		}
	}
}
