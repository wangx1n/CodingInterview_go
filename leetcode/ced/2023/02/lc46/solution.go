package lc46

//
//func permute(nums []int) [][]int {
//	n := len(nums)
//	used := make([]bool, n)
//	var ans []int
//	var res [][]int
//	dfs(nums, ans, &res, 0, used)
//	return res
//}
//
//func dfs(nums, ans []int, res *[][]int, count int, used []bool) {
//	if count == len(nums) {
//		*res = append(*res, append([]int{}, ans...))
//	}
//	for i := 0; i < len(nums); i++ {
//		if !used[i] {
//			ans = append(ans, nums[i])
//			used[i] = true
//			dfs(nums, ans, res, count+1, used)
//			used[i] = false
//			ans = ans[:len(ans)-1]
//		}
//	}
//}
