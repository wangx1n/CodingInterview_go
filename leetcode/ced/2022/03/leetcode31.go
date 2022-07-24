/**
 * @Author: luccifer
 * @Description:
 * @File:  leetcode31
 * @Version: 1.0.0
 * @Date: 2022/3/17 19:13
 */

package _3

func nextPermutation(nums []int) {
	r := func(a []int) {
		n := len(a)
		for i := 0; i < n/2;i++ {
			a[n-1-i], a[i] = a[i], a[n-1-i]
		}
	}
	n := len(nums)
	i := n - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}
	if i >= 0 {
		j := n - 1
		for j > i && nums[i] >= nums[j] {
			j--
		}
		nums[j], nums[i] = nums[i], nums[j]
	}
	r(nums[i+1:])
}

