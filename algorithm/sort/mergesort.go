//Author  :     knight
//Date    :     2020/06/25 20:20:51
//Version :     1.0
//Email   :     knight2347@163.com
//idea    :     归并排序

package sort

//MergeSort 归并排序
func MergeSort(nums []int) []int {
	return mergersort(nums)
}

func mergersort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	mid := len(nums) / 2
	left := mergersort(nums[:mid])
	right := mergersort(nums[mid:])
	return merge(left, right)
}

func merge(nums1, nums2 []int) []int {
	nums := make([]int, 0)
	index1, index2 := 0, 0
	for index1 < len(nums1) && index2 < len(nums2) {
		if nums1[index1] < nums2[index2] {
			nums = append(nums, nums1[index1])
			index1++
		} else {
			nums = append(nums, nums2[index2])
			index2++
		}
	}
	nums = append(nums, nums1[index1:]...)
	nums = append(nums, nums2[index2:]...)
	return nums
}