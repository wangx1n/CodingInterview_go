/**
 * @Author: luccifer
 * @Description:
 * @File:  leetcode33
 * @Version: 1.0.0
 * @Date: 2022/3/17 20:49
 */

package _3

func search(nums []int, target int) int {
	n := len(nums)
	l, r := 0, n-1
	if n == 0 {
		return -1
	}
	if n == 1 {
		if nums[0] != target {
			return -1
		}else {
			return 0
		}
	}
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			return mid
		}
		// 取有序部分二分
		if nums[0] <= nums[mid] {
			// 有序部分在mid左边，判断target是否在mid左边
			if nums[0] <= target && target < nums[mid] {
				// target在mid左边，置r = mid - 1
				r = mid - 1
			} else {
				// target不在mid左侧，在非有序部分，则l = mid + 1
				// 这个时候下一次循环的mid就不在左边了，在右边，进循环的时候进入下方的else,往右搜
				// 这里想多了，就是简单的分治
				l = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[n-1] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	return -1
}
