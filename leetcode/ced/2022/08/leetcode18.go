package _8

import "sort"

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	var res [][]int
	n := len(nums)
	for first := 0; first < n-3; first++ {
		if first != 0 && nums[first] == nums[first-1] {
			continue
		}
		for second := first + 1; second < n-2; second++ {
			if second != first+1 && nums[second] == nums[second-1] {
				continue
			}
			for third, forth := second+1, n-1; third < forth; {
				if sum := nums[first] + nums[second] + nums[third] + nums[forth]; sum == target {
					res = append(res, []int{nums[first], nums[second], nums[third], nums[forth]})
					for third++; third < forth && nums[third] == nums[third-1]; {
					}
					for forth--; third < forth && nums[forth] == nums[forth+1]; {
					}
				} else {
					if sum < target {
						third++
					} else {
						forth--
					}
				}
			}
		}
	}
	return res
}
