/**
 * @Author: luccifer
 * @Description:
 * @File:  leetcode16_doublePoint
 * @Version: 1.0.0
 * @Date: 2022/3/6 22:03
 */

package array

import (
	"math"
	"sort"
)

func main() {
	a := []int{-1,2,1,-4}
	threeSumClosest(a, 1)
}

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	res, n := math.MaxInt32, len(nums)
	for first := 0; first < n; first++ {
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		second, third := first+1, n-1
		for second < third {
			cur := nums[first] + nums[second] + nums[third]
			if cur == target {
				return cur
			}
			if Abs(target - cur) < Abs(target - res) {
				res = cur
			}
			if cur > target {
				tempThird := third - 1
				for second < tempThird && nums[third] == nums[tempThird] {
					tempThird--
				}
				third = tempThird
			} else {
				tempSecond := second + 1
				for tempSecond < third && nums[tempSecond] == nums[second] {
					tempSecond++
				}
				second = tempSecond
			}
		}
	}
	return res
}