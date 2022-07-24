package main

func trap_MStack(height []int) (res int) {
	var stack []int
	for i, h := range height {
		for len(stack) > 0 && h > height[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			left := stack[len(stack)-1]
			weight := i - left - 1
			cheight := min(height[left], h) - height[top]
			res += weight * cheight
		}
		stack = append(stack, i)
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
