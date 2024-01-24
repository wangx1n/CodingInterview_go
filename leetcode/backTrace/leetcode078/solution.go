package leetcode078

func subsets(nums []int) [][]int {
	var ret [][]int
	dfs(&ret, nums, []int{}, 0)
	return ret
}

func dfs(ret *[][]int, nums, ans []int, index int) {
	*ret = append(*ret, append([]int{}, ans...))

	for i := index; i < len(nums); i++ {
		ans = append(ans, nums[i])
		dfs(ret, nums, ans, i+1)
		ans = ans[:len(ans)-1]
	}
}
