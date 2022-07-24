/**
 * @Author: luccifer
 * @Description:
 * @File:  leetcode16_self
 * @Version: 1.0.0
 * @Date: 2022/3/4 23:57
 */

package array

import (
	"math"
	"sort"
)
// [-3, 0, 1, 2]
func threeSumClosest_self(nums []int, target int) int {
	sort.Ints(nums)
	res, n := math.MaxInt32, len(nums)
	for first := 0; first < n; first++ {
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		for second := first + 1; second < n; second++ {
			third := n - 1
			for second < third {
				cur := nums[first] + nums[second] + nums[third]
				if cur == target {
					return cur
				}
				if Abs(target - cur) < Abs(target - res) {
					res = cur
				}
				third--
			}

		}
	}
	return res
}

func Abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}
