package _8

func isValid(s string) bool {
	n := len(s)
	if n%2 != 0 {
		return false
	}
	m := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := make([]byte, 0)
	for _, c := range s {
		if m[byte(c)] > 0 {
			if len(stack) == 0 || m[byte(c)] != stack[len(stack)-1] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, byte(c))
		}
	}
	if len(stack) > 0 {
		return false
	}
	return true
}
