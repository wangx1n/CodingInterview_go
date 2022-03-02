/**
 * @Author: luccifer
 * @Description:
 * @File:  leetcode11
 * @Version: 1.0.0
 * @Date: 2022/3/2 16:36
 */

package array


func maxArea(height []int) int {
	l ,r := 0, len(height) - 1
	res := 0
	for l < r {
		area := (r - l) * (height[l], height[r])
	}
}
