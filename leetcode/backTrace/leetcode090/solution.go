package leetcode090

import "sort"

func subsetsWithDup(nums []int) [][]int {
	var ret [][]int
	sort.Ints(nums)
	dfs(&ret, nums, []int{}, 0, make([]bool, len(nums)))
	return ret
}

func dfs(ret *[][]int, nums, ans []int, index int, used []bool) {
	*ret = append(*ret, append([]int{}, ans...))

	for i := index; i < len(nums); i++ {
		if i > 0 && !used[i-1] && nums[i] == nums[i-1] {
			continue
		}
		used[i] = true
		ans = append(ans, nums[i])
		dfs(ret, nums, ans, i+1, used)
		ans = ans[:len(ans)-1]
		used[i] = false
	}
}
