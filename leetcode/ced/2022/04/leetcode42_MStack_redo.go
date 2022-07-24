package _4

func _trap(height []int) int {
	var stack []int
	res := 0
	for i, h := range height {
		for len(stack) > 0 && h > height[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			left := stack[len(stack)-1]
			weight := i - left - 1
			height := min(height[left], h) - height[top]
			res += weight * height
		}
		stack = append(stack, i)
	}
	return res
}

func _min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
