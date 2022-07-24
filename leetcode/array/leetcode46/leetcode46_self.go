package array

var ans [][]int

/**
全排列
*/

func permute(nums []int) [][]int {
	ans = [][]int{}
	var res []int
	length := len(nums)
	used := make([]bool, length)
	retrospect46(res, nums, used, length)
	return ans
}

func retrospect46(res, nums []int, used []bool, length int) {
	if len(res) == length {
		ans = append(ans, append([]int(nil), res...))
		return
	}
	for i := 0; i < length; i++ {
		if !used[i] {
			res = append(res, nums[i])
			used[i] = true
			retrospect46(res, nums, used, length)
			used[i] = false
			res = res[:len(res)-1]
		}
	}
}
