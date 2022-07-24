package main

func maxSubArray(nums []int) int {
	max := nums[0]
	n := len(nums)
	for i := 1; i < n; i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] = nums[i] + nums[i-1]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}
