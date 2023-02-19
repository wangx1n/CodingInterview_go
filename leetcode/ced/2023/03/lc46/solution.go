package lc46

func permute(nums []int) [][]int {
	var ans []int
	var res [][]int
	used := make([]bool, len(nums))
	dfs(nums, ans, &res, 0, used)
	return res
}

func dfs(nums, ans []int, res *[][]int, begin int, used []bool) {
	if len(ans) == len(nums) {
		*res = append(*res, append([]int{}, ans...))
		return
	}

	for i := begin; i < len(nums); i++ {
		if !used[i] {
			used[i] = true
			ans = append(ans, nums[i])
			dfs(nums, ans, res, begin, used)
			used[i] = false
			ans = ans[:len(ans)-1]
		}
	}
}
