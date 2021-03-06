/**
 * @Author: luccifer
 * @Description:
 * @File:  leetcode15_doublePoint
 * @Version: 1.0.0
 * @Date: 2022/3/3 17:40
 */

package main

import "sort"

func threeSum_doublePoint(nums []int) (ans [][]int) {
	n := len(nums)
	sort.Ints(nums)
	for first := 0; first < n; first++ {
		if first > 0 && nums[first] == nums[first - 1] {
			continue
		}
		third := n - 1
		target := -1 * nums[first]
		for second := first + 1; second < n; second++ {
			if second > first + 1 && nums[second] == nums[second - 1] {
				continue
			}
			for  second < third && nums[second] + nums[third] > target {
				third--
			}
			if second == third {
				break
			}
			if (nums[second] + nums[third] == target) {
				ans = append(ans, []int{nums[first], nums[second], nums[third]})
			}
		}
	}
	return
}
