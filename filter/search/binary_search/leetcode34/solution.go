package leetcode34

func searchRange(nums []int, target int) []int {
	leftIdx := searchWithCase(nums, target, true)
	right := searchWithCase(nums, target, false) - 1
	if leftIdx <= right && right < len(nums) && nums[leftIdx] == target && nums[right] == target {
		return []int{leftIdx, right}
	}
	return []int{-1, -1}
}

func searchWithCase(nums []int, target int, withCase bool) int {
	n := len(nums)
	ans := n
	l, r := 0, n-1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] > target || (withCase && nums[mid] >= target) {
			r = mid - 1
			ans = mid
		} else {
			l = mid + 1
		}
	}
	return ans
}
