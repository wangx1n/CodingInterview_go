package leetcode011

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	maxValue := 0
	for left < right {
		tempMax := (right - left) * min(height[left], height[right])
		if tempMax > maxValue {
			maxValue = tempMax
		}
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return maxValue
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
