//Author  :     knight
//Date    :     2020/07/10 10:59:28
//Version :     1.0
//Email   :     knight2347@163.com
//idea    :     栈的测试

package stack

import "testing"

func TestIsPopOrder(t *testing.T) {
	pushOrder := []int{1, 2, 3, 4, 5}
	popOrder := []int{4, 5, 3, 1, 2}
	ok := IsPopOrder(pushOrder, popOrder)
	if !ok {
		t.Log("success")
	} else {
		t.Error("failed")
	}
}

func TestLongest(t *testing.T) {
	s := "()(()()"
	t.Log(longestValidParentheses(s))
}

func TestIsValid(t *testing.T) {
	s := "()[]{"
	ans := isValid(s)
	t.Log(ans)
}
