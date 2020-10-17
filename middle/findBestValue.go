//Author  :     knight
//Date    :     2020/06/14 10:19:48
//Version :     1.0
//Email   :     knight2347@163.com
//idea    :     排序+枚举+二分查找

package main

import (
	"sort"
)

func findBestValue(arr []int, target int) int {
	sort.Ints(arr)
	n := len(arr)
	presum := make([]int, n)
	for i := 1; i < n; i++ {
		presum[i] = presum[i-1] + arr[i-1]
	}

	return 0
}
