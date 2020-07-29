//Author  :     knight
//Date    :     2020/07/28 12:02:29
//Version :     1.0
//Email   :     knight2347@163.com
//idea    :     堆的测试文件

package head

import "testing"

func Test_top(t *testing.T) {
	var h Header
	h = &MinHead{make([]int, 0)}
	t.Log(h.Top())
}