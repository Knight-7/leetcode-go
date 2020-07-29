//Author  :     knight
//Date    :     2020/07/26 11:39:00
//Version :     1.0
//Email   :     knight2347@163.com
//idea    :     最长有效括号

package stack

func longestValidParentheses(s string) int {
	stack := make([]int, 1)
	stack[0] = -1
	sumLength := 0
	for k, v := range s {
		if v == '(' {
			stack = append(stack, k)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, k)
			} else {
				sumLength = max(sumLength, k - stack[len(stack)-1])
			}
		}
	}
	return sumLength
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}