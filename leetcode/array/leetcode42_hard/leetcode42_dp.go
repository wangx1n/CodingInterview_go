package main

import (
	"fmt"
	"gitee.com/wangx1n9347/common_go/Array"
)

func main() {
	fmt.Println(trap([]int{4, 2, 0, 3, 2, 5}))
}

func trap(height []int) (res int) {
	n := len(height)
	leftMax := make([]int, n)
	leftMax[0] = height[0]
	for i := 1; i < n; i++ {
		leftMax[i] = Array.MinOrMax(height[i], leftMax[i-1], Array.GetMax)
	}
	rightMax := make([]int, n)
	rightMax[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
		rightMax[i] = Array.MinOrMax(height[i], rightMax[i+1], Array.GetMax)
	}
	for i, h := range height {
		res += Array.MinOrMax(leftMax[i], rightMax[i], Array.GetMin) - h
	}
	return
}
