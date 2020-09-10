/*
* @Author: Knight
* @Date:   2020/09/10 08:56:18
* @Email:  knight2347@163.com
* @Ideal:
 */

package DFS

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	ans := make([][]int, 0)
	sort.Ints(candidates)
	var dfs func(curP, curSum int, curNum []int)
	dfs = func(curP, curSum int, curNum []int) {
		if curSum > target {
			return
		}
		if curSum == target {
			t := make([]int, len(curNum))
			copy(t, curNum)
			ans = append(ans, t)
			return
		}
		for i := curP; i < len(candidates); i++ {
			curNum = append(curNum, candidates[i])
			dfs(i + 1, curSum + candidates[i], curNum)
			curNum = curNum[:len(curNum) - 1]
		}
	}
	dfs(0, 0, []int{})
	return ans
}
