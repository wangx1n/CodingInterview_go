package _4

func trap(height []int) int {
	res := 0
	n := len(height)
	maxLeft := make([]int, n)
	maxLeft[0] = height[0]
	for i := 1; i < n; i++ {
		maxLeft[i] = max(height[i], maxLeft[i-1])
	}
	maxRight := make([]int, n)
	maxRight[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
		maxRight[i] = max(height[i], maxRight[i+1])
	}
	for i, h := range height {
		res += min(maxRight[i], maxLeft[i]) - h
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
