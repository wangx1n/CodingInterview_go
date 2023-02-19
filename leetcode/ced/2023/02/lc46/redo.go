package lc46

func permute(nums []int) [][]int {
	var ans []int
	var res [][]int
	used := make([]bool, len(nums))
	dfs(nums, ans, &res, 0, used)
	return res
}

func dfs(nums, ans []int, res *[][]int, count int, used []bool) {
	if count == len(nums) {
		*res = append(*res, append([]int{}, ans...))
	}

	for i := 0; i < len(nums); i++ {
		if !used[i] {
			used[i] = true
			ans = append(ans, nums[i])
			dfs(nums, ans, res, count+1, used)
			ans = ans[:len(ans)-1]
			used[i] = false
		}
	}
}
