//Author  :     knight
//Date    :     2020/07/28 12:02:29
//Version :     1.0
//Email   :     knight2347@163.com
//idea    :     堆的测试文件

package heap

import "testing"

func TestTopK(t *testing.T) {
	nums := []int{1, 1, 1, 2, 2, 3}
	ans := topKFrequent1(nums, 2)
	t.Log(ans)
}
