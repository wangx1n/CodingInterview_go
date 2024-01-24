package leetcode077

func combine(n int, k int) [][]int {
	nums := make([]int, 0, n)
	for i := 1; i <= n; i++ {
		nums = append(nums, i)
	}
	var ret [][]int

	dfs(&ret, []int{}, 0, k, nums)
	return ret
}

func dfs(ret *[][]int, ans []int, index int, numsLen int, nums []int) {
	if len(ans) == numsLen {
		*ret = append(*ret, append([]int{}, ans...))
		return
	}

	for i := index; i < len(nums); i++ {
		ans = append(ans, nums[i])
		dfs(ret, ans, i+1, numsLen, nums)
		ans = ans[:len(ans)-1]
	}
}
