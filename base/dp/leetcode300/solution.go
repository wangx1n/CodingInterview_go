package leetcode300

// mid
// dp[i] = max(dp[j]+1, dp[i])
func lengthOfLIS(nums []int) int {
	dp := make([]int, 0, len(nums))
	for i := 0; i < len(nums); i++ {
		dp = append(dp, 1)
	}

	res := 1

	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] >= nums[i] {
				continue
			}
			// nums = [0,1,0,3,2,3], 如果没有max, 则index=4位置的dp值为2
			dp[i] = max(dp[j]+1, dp[i])

			if res <= dp[i] {
				res = dp[i]
			}
		}
	}

	return res
}

func max(int2 int, int3 int) int {
	if int2 > int3 {
		return int2
	}
	return int3
}
