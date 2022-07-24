package main

var m = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	ans := 0
	length := len(s)
	for i := length - 1; i >= 0; i-- {
		if s[i] == 'V' || s[i] == 'X' {
			if i-1 >= 0 && s[i-1] == 'I' {
				ans += m[s[i]] - 1
				i--
			} else {
				ans += m[s[i]]
			}
		} else if s[i] == 'L' || s[i] == 'C' {
			if i-1 >= 0 && s[i-1] == 'X' {
				ans += m[s[i]] - 10
				i--
			} else {
				ans += m[s[i]]
			}
		} else if s[i] == 'D' || s[i] == 'M' {
			if i-1 >= 0 && s[i-1] == 'C' {
				ans += m[s[i]] - 100
				i--
			} else {
				ans += m[s[i]]
			}
		} else {
			ans += 1
		}
	}
	return ans
}
