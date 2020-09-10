/*
* @Author: Knight
* @Date:   2020/09/10 08:56:18
* @Email:  knight2347@163.com
* @Ideal:  leetcode40 组合数2
 */

package DFS

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	ans := make([][]int, 0)
	path, nums := []int{}, []int{}
	sort.Ints(candidates)
	cnt := map[int]int{}
	for _, v := range candidates {
		if _, ok := cnt[v]; !ok {
			nums = append(nums, v)
			cnt[v]++
		} else {
			cnt[v]++
		}
	}
	var dfs func(curP, reset int)
	dfs = func(curP, reset int) {
		if reset == 0 {
			ans = append(ans, append([]int{}, path...))
			return
		}
		if curP == len(nums) || reset < nums[curP] {
			return
		}

		dfs(curP+1, reset)

		most := min(reset/nums[curP], cnt[nums[curP]])
		for i := 1; i <= most; i++ {
			path = append(path, nums[curP])
			dfs(curP+1, reset-i*nums[curP])
		}
		for i := 1; i <= most; i++ {
			path = path[:len(path)-1]
		}
	}
	dfs(0, target)
	return ans
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
