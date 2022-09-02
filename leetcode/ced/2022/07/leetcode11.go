package _7

func maxArea(height []int) int {
	l, r := 0, len(height)-1
	maxSize := 0
	for l < r {
		size := (r - l + 1) * min_(height[l], height[r])
		if size > maxSize {
			maxSize = size
		}
		if height[l] <= height[r] {
			l++
		} else {
			r--
		}
	}
	return maxSize
}
