package _7

func romanToInt(s string) int {
	var res int
	m := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	n := len(s)
	for i := 0; i < n; i++ {
		if i == 0 {
			res += m[s[i]]
		}
		if m[s[i]] < m[s[i-1]] && i+1 < n {
			res += m[s[i+1]] - 1
			i++
		} else {
			res += m[s[i]]
		}
	}
	return res
}
