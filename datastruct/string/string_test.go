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