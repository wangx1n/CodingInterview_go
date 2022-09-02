package _022_07_30

func isPalindrome(x int) bool {
	if x < 0 || x != 0 && x%10 == 0 {
		return false
	}

	reverseNum := 0
	for reverseNum < x {
		reverseNum = reverseNum*10 + x/10
		x /= 10
	}

	return reverseNum == x || x == reverseNum/10
}
