/**
 * @Author: luccifer
 * @Description:
 * @File:  leetcode015
 * @Version: 1.0.0
 * @Date: 2022/3/2 19:42
 */
package main

import (
	"fmt"
	carray "gitee.com/wangx1n9347/common_go/Array"
	"sort"
)

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(nums))
}

func threeSum(nums []int) [][]int {
	n := len(nums)
	res := [][]int{}
	for i := 0; i < n-1; i++ {
		target := 0 - nums[i]
		tmp := sum(nums[i+1:], target)
		for _, x := range tmp {
			if ok, _ := carray.ExistInTwoDimensional(x, res); !ok {
				res = append(res, x)
			}
		}
	}
	return res
}

func sum(nums []int, target int) (res [][]int) {
	if len(nums) < 2 {
		return
	}
	m := map[int]int{}
	for i, x := range nums {
		if p, ok := m[target-x]; ok {
			tmp := []int{0 - target, nums[p], x}
			sort.Slice(tmp, func(i, j int) bool {
				return tmp[i] < tmp[j]
			})
			res = append(res, tmp)
		} else {
			m[x] = i
		}
	}
	return
}
