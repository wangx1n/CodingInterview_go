package _02112
/**
hard
689. 三个无重叠子数组的最大和
给你一个整数数组 nums 和一个整数 k ，找出三个长度为 k 、互不重叠、且 3 * k 项的和最大的子数组，并返回这三个子数组。

以下标的数组形式返回结果，数组中的每一项分别指示每个子数组的起始位置（下标从 0 开始）。如果有多个结果，返回字典序最小的一个。



示例 1：

输入：nums = [1,2,1,2,6,7,5,1], k = 2
输出：[0,3,5]
解释：子数组 [1, 2], [2, 6], [7, 5] 对应的起始下标为 [0, 3, 5]。
也可以取 [2, 1], 但是结果 [1, 3, 5] 在字典序上更大。
示例 2：

输入：nums = [1,2,1,2,1,2,1,2,1], k = 2
输出：[0,2,4]


提示：

1 <= nums.length <= 2 * 104
1 <= nums[i] < 216
1 <= k <= floor(nums.length / 3)
 */
func maxSumOfThreeSubarrays(nums []int, k int) []int {
	sum1, sum2, sum3 := 0, 0, 0
	maxSum1, maxSum2, totalSum := 0, 0, 0
	maxSum1Index, maxSum2Index1, maxSum2Index2 := 0, 0, 0
	var ans []int
	for i := k * 2; i < len(nums); i++ {
		sum1 += nums[i - 2 * k]
		sum2 += nums[i - k]
		sum3 += nums[i]
		if i >= 3 * k - 1 {
			if sum1 > maxSum1 {
				maxSum1 = sum1
				maxSum1Index = i - 3 * k + 1
			}
 			if maxSum1 + sum2 > maxSum2 {
				maxSum2 = maxSum1 + sum2
				maxSum2Index2 = i - 2 * k + 1
				maxSum2Index1 = maxSum1Index
			}
			if maxSum2 + sum3 > totalSum {
				totalSum = maxSum2 + sum3
				ans = []int{maxSum2Index1, maxSum2Index2, i - k + 1}
			}
			sum1 -= nums[i - 3 * k + 1]
			sum2 -= nums[i - 2 * k + 1]
			sum3 -= nums[i - k + 1]
		}
	}
	return ans
}