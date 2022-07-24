/**
 * @Author: luccifer
 * @Description:
 * @File:  leetcode18_self
 * @Version: 1.0.0
 * @Date: 2022/3/6 22:44
 */

package main

import (
	"sort"
)

func main() {
	nums := []int{2,2,2,2,2}
	fourSum(nums, 8)
}

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	var res [][]int
	n := len(nums)
	for first := 0; first < n; first++ {
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		tmp := threeSum(nums[first + 1:], target - nums[first])
		for i := 0; i < len(tmp); i++ {
			tmpCol := tmp[i]
			tmpCol = append(tmpCol, nums[first])
			sort.Ints(tmpCol)
			res = append(res, tmpCol)
		}
	}
	return res
}

func threeSum(nums []int, target int) (res [][]int) {
	n := len(nums)
	if n < 3 {
		return
	}
	for first := 0; first < n; first++ {
		if first > 0 && nums[first] == nums[first - 1] {
			continue
		}
		third := n - 1
		for second := first+1; second < n; second++ {
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			for second < third && nums[first] + nums[second] + nums[third] > target{
				third--
			}
			if second == third {
				break
			}
			if nums[first] + nums[second] + nums[third] == target {
				res = append(res, []int{nums[first], nums[second], nums[third]})
			}
		}
	}
	return
}
