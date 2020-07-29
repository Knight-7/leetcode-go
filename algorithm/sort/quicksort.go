//Author  :     knight
//Date    :     2020/06/25 20:10:54
//Version :     1.0
//Email   :     knight2347@163.com
//idea    :     快速排序

package sort

import "fmt"

//QuickSort 快速排序
func QuickSort(nums []int) []int {
	quicksort(nums, 0, len(nums) - 1)
	return nums
}

func quicksort(nums []int, l, r int) {
	if l >= r {
		return
	}

	pivot := partition(nums, l, r)
	quicksort(nums, l, pivot - 1)
	quicksort(nums, pivot + 1, r)
}

func partition(nums []int, l, r int) int {
	small := l - 1
	for i := l; i < r; i++ {
		if nums[i] < nums[r] {
			small++
			if i != small {
				swap(&nums[i], &nums[small])
			}
		}
	}
	small++
	swap(&nums[small], &nums[r])
	return small
}

func swap(a, b *int) {
	tmp := *a
	*a = *b
	*b = tmp
}

func main() {
	nums := []int{1, 5, 9, 13, 3, 7, 0, 9, 7}
	QuickSort(nums)
	for _, v := range nums {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
}