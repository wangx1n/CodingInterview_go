package main

import "sort"

var ans [][]int

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	ans = [][]int{}
	var res []int
	used := make([]bool, len(nums))
	retrospect47(res, nums, used, len(nums))
	return ans
}

func retrospect47(res, nums []int, used []bool, length int) {
	if len(res) == length {
		ans = append(ans, append([]int{}, res...))
		return
	}
	for i := 0; i < length; i++ {
		if !used[i] {
			if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
				continue
			}
			res = append(res, nums[i])
			used[i] = true
			retrospect47(res, nums, used, length)
			used[i] = false
			res = res[:len(res)-1]
		}
	}
}
