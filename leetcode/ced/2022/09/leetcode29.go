package _9

import "math"

// 快速乘 + 二分查找

// a * b >= x 找到乘积大于x的最小数

func quickAdd(a, b, x int) bool {
	res := 0
	for b != 0 {
		if b&1 != 0 {
			res += a
		}
		a += a
		b >>= 1
	}
	if res >= x {
		return true
	}
	return false
}

func divide(dividend int, divisor int) int {
	if dividend == 0 || divisor == 0 {
		return 0
	}

	if dividend == math.MinInt32 {
		if divisor == 1 {
			return math.MinInt32
		}
		if divisor == -1 {
			return math.MaxInt32
		}
	}

	if divisor == math.MinInt32 {
		if dividend == math.MinInt32 {
			return 1
		}
		return 0
	}

	res := 0
	revert := false

	if dividend > 0 {
		dividend = -dividend
		revert = !revert
	}
	if divisor > 0 {
		divisor = -divisor
		revert = !revert
	}

	left, right := 1, math.MaxInt32
	for left <= right {
		mid := left + (right-left)>>1
		if quickAdd(divisor, mid, dividend) {
			res = mid
			if mid == math.MaxInt32 {
				break
			}
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	if revert {
		res = -res
	}
	return res
}
