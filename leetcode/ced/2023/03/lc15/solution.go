package lc15

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	var res [][]int

	for first := 0; first < len(nums); first++ {
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		target := -nums[first]
		second := first + 1
		third := len(nums) - 1
		for ; second < len(nums); second++ {
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			for second < third && nums[second]+nums[third] > target {
				third--
			}
			if second == third {
				break
			}
			if nums[second]+nums[third] == target {
				ans := []int{nums[first], nums[second], nums[third]}
				res = append(res, ans)
			}
		}
	}
	return res
}
