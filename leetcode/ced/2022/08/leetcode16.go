package _8

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	res, n := math.MaxInt32, len(nums)
	var update func(a int)
	update = func(a int) {
		if abs(a, target) < abs(res, target) {
			res = a
		}
	}
	for first := 0; first < n-2; first++ {
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		second := first + 1
		third := n - 1
		for second < third {
			sum := nums[first] + nums[second] + nums[third]
			if sum == target {
				return sum
			}
			update(sum)
			if sum < target {
				for second < third && nums[second] == nums[second-1] {
					second++
				}
				second++
			} else {
				for second < third && nums[third] == nums[third-1] {
					third--
				}
				third--
			}
		}
	}
	return res
}

func abs(a, b int) int {
	if a-b < 0 {
		return -1 * (a - b)
	}
	return a - b
}
