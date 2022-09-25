package _9

import "math"

func divide__(dividend int, divisor int) int {
	if dividend == 0 || divisor == 0 {
		return 0
	}
	if dividend == math.MinInt32 {
		if divisor == -1 {
			return math.MaxInt32
		}
		if divisor == 1 {
			return math.MinInt32
		}
	}
	if divisor == math.MinInt32 {
		if dividend == math.MinInt32 {
			return 1
		}
		return 0
	}

	flag := false
	if dividend < 0 {
		flag = !flag
		dividend = -dividend
	}
	if divisor < 0 {
		flag = !flag
		divisor = -divisor
	}

	left, right := 0, math.MaxInt32
	res := 0
	for left <= right {
		mid := left + (right-left)>>1
		if fastAdd(divisor, mid, dividend) {
			right = mid - 1
		} else {
			res = mid
			if res == math.MaxInt32 {
				break
			}
			left = mid + 1
		}
	}

	if flag {
		res = -res
	}
	return res
}

func fastAdd(a, b, y int) bool {
	res := 0
	for b != 0 {
		if b&1 != 0 {
			res += a
		}
		a += a
		b >>= 1
	}
	return res > y
}
