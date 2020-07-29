//Author  :     knight
//Date    :     2020/07/09 10:31:28
//Version :     1.0
//Email   :     knight2347@163.com
//idea    :     数组的测试

package array

import "testing"

func TestReverseArrayMin(t *testing.T) {
	nums := []int{3, 4, 5, 1, 2}
	min := ReverseArrayMin(nums, 5)
	if min == 1 {
		t.Log(min)
		t.Log("测试成功")
	} else {
		t.Error(min)
		t.Error("测试失败")
	}
}