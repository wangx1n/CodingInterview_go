package _1

func isValid(s string) bool {
	var stack []byte

	for _, ss := range s {
		if ss == '(' || ss == '[' || ss == '{' {
			stack = append(stack, byte(ss))
			continue
		}
		if len(stack) == 0 && (ss == ')' || ss == ']' || ss == '}') {
			return false
		}
		switch ss {
		case ')':
			if stack[len(stack)-1] != '(' {
				return false
			}
		case ']':
			if stack[len(stack)-1] != '[' {
				return false
			}
		case '}':
			if stack[len(stack)-1] != '{' {
				return false
			}
		}
		stack = stack[:len(stack)-1]
	}
	return len(stack) == 0
}
