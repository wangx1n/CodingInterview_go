/**
 * @Author: luccifer
 * @Description:
 * @File:  leetcode11
 * @Version: 1.0.0
 * @Date: 2022/3/2 16:36
 */

package main

import (
	carray "gitee.com/wangx1n9347/common_go/Array"
)

func maxArea(height []int) int {
	l ,r := 0, len(height) - 1
	res := 0
	for l < r {
		area := (r - l)  * carray.MinOrMax(height[l], height[r], carray.GetMin)
		res = carray.MinOrMax(area, res, carray.GetMax)
		if height[l] <= height[r] {
			l++
		} else {
			r--
		}
	}
	return res
}