//Author  :     knight
//Date    :     2020/06/25 20:27:10
//Version :     1.0
//Email   :     knight2347@163.com
//idea    :     排序测试

package sort

import (
	"testing"
)

//TestQuickSort 测试快排
func TestQuickSort(t *testing.T) {
	nums := []int{3, 4, 2, 3, 8, 6}
	nums2 := []int{2, 3, 3, 4, 6, 8}
	nums = QuickSort(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] != nums2[i] {
			t.Error("测试失败")
		}
	}
}

func TestMergeSort(t *testing.T) {
	nums := []int{3, 4, 2, 3, 8, 6}
	nums2 := []int{2, 3, 3, 4, 6, 8}
	nums = MergeSort(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] != nums2[i] {
			t.Error("测试失败")
		}
	}
}
