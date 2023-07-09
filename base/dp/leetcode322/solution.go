package leetcode322

import "math"

// mid
func coinChange(coins []int, amount int) int {
	dp := make([]int, 0, amount+1)
	for i := 0; i < amount+1; i++ {
		dp = append(dp, math.MaxInt)
	}
	dp[0] = 0
	for i := 0; i < len(dp); i++ {
		for _, coin := range coins {
			if i-coin < 0 {
				continue
			}
			if dp[i-coin] == math.MaxInt {
				continue
			}
			dp[i] = min(dp[i], dp[i-coin]+1)
		}
	}
	if dp[amount] == math.MaxInt {
		return -1
	}
	return dp[amount]
}

func min(int2 int, int3 int) int {
	if int2 > int3 {
		return int3
	}
	return int2
}
