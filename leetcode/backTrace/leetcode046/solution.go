package leetcode046

func permute(nums []int) [][]int {
	var ret [][]int
	dfs(nums, &ret, []int{}, make([]bool, len(nums)))
	return ret
}

func dfs(nums []int, ret *[][]int, ans []int, used []bool) {
	if len(ans) == len(nums) {
		*ret = append(*ret, append([]int{}, ans...))
		return
	}
	for i := 0; i < len(nums); i++ {
		if !used[i] {
			ans = append(ans, nums[i])
			used[i] = true
			dfs(nums, ret, ans, used)
			used[i] = false
			ans = ans[:len(ans)-1]
		}
	}
}
