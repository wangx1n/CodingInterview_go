package common

// stringAdd 字符串相加
func StringAdd(x, y string) string {
	res := []byte{}
	carry, cur := 0, 0
	for x != "" || y != "" || carry != 0 {
		cur = carry
		if x != "" {
			cur += int(x[len(x)-1] - '0')
			x = x[:len(x)-1]
		}
		if y != "" {
			cur += int(y[len(y)-1] - '0')
			y = y[:len(y)-1]
		}
		carry = cur / 10
		cur %= 10
		res = append(res, byte(cur)+'0')
	}
	for i, n := 0, len(res); i < n/2; i++ {
		res[i], res[n-1-i] = res[n-1-i], res[i]
	}
	return string(res)
}
