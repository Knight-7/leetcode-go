//Author  :     knight
//Date    :     2020/07/10 14:46:41
//Version :     1.0
//Email   :     knight2347@163.com
//idea    :

package string

import "testing"

func TestIsMach(t *testing.T) {
	str := "aa"
	pattern := "a*"
	isManch := IsMatch(str, pattern)
	if isManch {
		t.Log("success")
	} else {
		t.Log("failed")
	}
}

func TestAddString(t *testing.T) {
	num1 := "13213546543"
	num2 := "78979"
	ans := addString(num1, num2)
	t.Log(ans)
}
