//Author  :     knight
//Date    :     2020/06/21 16:23:03
//Version :     1.0
//Email   :     knight2347@163.com
//idea    :     回溯算法

package main

import "fmt"

func subsets(nums []int) [][]int {
	ans := make([][]int, 0)
	tmp := make([]int, 0)
	backtrack(nums, 0, tmp, &ans)
	return ans
}

func backtrack(nums []int, pos int, tmp []int, ans *[][]int) {
	t := make([]int, len(tmp))
	copy(t, tmp)
	*ans = append(*ans, t)
	for i := pos; i < len(nums); i++ {
		tmp = append(tmp, nums[i])
		backtrack(nums, i + 1, tmp, ans)
		tmp = tmp[:len(tmp) - 1]
	}
}

func main(){
	nums := []int{1, 2, 3}
	fmt.Println(subsets(nums))
}