//Author  :     knight
//Date    :     2020/06/12 10:09:30
//Version :     1.0
//Email   :     knight2347@163.com
//idea    :     方法一：(三个数之和：将三个数之和转化为两个数之和，
//				即先确定一个数x，那么在剩余的数字中查找是否有两个数字之和为-x的就行（两数之和使用要给map就很方便))    失败了，因为去重问题
//    			方法二：排序+双指针

package main

import (
	"fmt"
	"sort"
)

func equal(a1, a2 []int) bool {
	less := func(i, j int) bool {
		return i < j
	}
	sort.Slice(a1, less)
	sort.Slice(a2, less)
	if a1[0] == a1[0] && a1[1] == a1[1] {
		return true
	}
	return false
}

func towSum(nums []int, index, target int) (bool, [][]int) {
	m, ans := make(map[int]int), make([][]int, 0)

	for k, v := range nums {
		if k != index {
			_, ok := m[target-v]
			if ok {
				ans = append(ans, []int{nums[index], v, target - v})
			}
		} else {
			m[v] = k
		}
	}

	if len(ans) > 0 {
		return true, ans
	}
	return false, nil
}

func threeSum1(nums []int) [][]int {
	ans, flag := make([][]int, 0), make(map[int]bool)

	for i := 0; i < len(nums); i++ {
		_, ok := flag[nums[i]]
		if !ok {
			flag[nums[i]] = true
			find, tmp := towSum(nums, i, 0-nums[i])
			if find {
				for _, v := range tmp {
					ans = append(ans, v)
				}
			}
		}
	}

	return ans
}

func threeSum(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	ans := make([][]int, 0)

	for first := 0; first < n; first++ {
		if first > 0 && nums[first] == nums[first - 1] {
			continue
		}

		target := 0 - nums[first]
		third := n - 1

		for second := first + 1; second < n - 1; second++ {
			if second > first + 1 && nums[second] == nums[second - 1] {
				continue
			}

			for second < third && nums[second] + nums[third] > target {
				third--
			}

			if second == third {
				break
			}

			if nums[second] + nums[third] == target {
				ans = append(ans, []int{nums[first], nums[second], nums[third]})
			}
		}
		
	}

	return ans
}

func main1() {
	nums := []int{-1, 0, 1, 2, -1, -4, 2}
	fmt.Println(threeSum(nums))
}
