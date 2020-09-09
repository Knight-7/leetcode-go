/*
* @Author: Knight
* @Date:   2020/9/9 9:34
* @Email:  knight2347@163.com
* @Ideal:  leetcode39 组合总和
 */

package DFS

func combinationSum(candidates []int, target int) [][]int {
	ans := make([][]int, 0)
	var dfs func(index, curSum int, curNum []int)
	dfs = func(index, curSum int, curNum []int) {
		if curSum > target {
			return
		}
		if curSum == target {
			t := make([]int, len(curNum))
			copy(t, curNum)
			ans = append(ans, t)
			return
		}
		for i := index; i < len(candidates); i++ {
			curNum = append(curNum, candidates[i])
			dfs(i, curSum+candidates[i], curNum)
			curNum = curNum[:len(curNum)-1]
		}
	}
	dfs(0, 0, []int{})
	return ans
}
