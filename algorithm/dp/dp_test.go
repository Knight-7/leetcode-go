//Author  :     knight
//Date    :     2020/07/09 09:24:12
//Version :     1.0
//Email   :     knight2347@163.com
//idea    :     dp测试

package dp

import (
	"testing"
)

func TestCuttingRope(t *testing.T) {
	ans := CuttingRope(4)
	if ans == 4 {
		t.Log("测试成功")
	} else {
		t.Log(ans)
		t.Error("测试失败")
	}
}

func TestRob(t *testing.T) {
	nums := []int{1, 2}
	ans := rob(nums)
	t.Log(ans)
}

func TestRob2(t *testing.T) {
	nums := []int{1}
	ans := rob2(nums)
	t.Log(ans)
}
