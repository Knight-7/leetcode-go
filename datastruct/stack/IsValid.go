package stack

//判断括号是否有效
func isValid(s string) bool {
	stack := make([]uint8, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			stack = append(stack, s[i])
		} else {
			if len(stack) > 0 && (s[i] == ')' && stack[len(stack)-1] == '(' ||
				s[i] == '}' && stack[len(stack)-1] == '{' ||
				s[i] == ']' && stack[len(stack)-1] == '[') {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	if len(stack) > 0 {
		return false
	}
	return true
}
