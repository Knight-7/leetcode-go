/*
* @Author: Knight
* @Date:   2020/09/11 09:00:58
* @Email:  knight2347@163.com
* @Ideal:  leetcode216 组合数三
 */

package DFS

func combinationSum3(k int, n int) [][]int {
	ans := make([][]int, 0)
	if n/k >= 9 {
		return ans
	}
	tmp := []int{}
	var dfs func(pos, reset int)
	dfs = func(pos, reset int) {
		if reset == 0 && len(tmp) == k {
			ans = append(ans, append([]int{}, tmp...))
			return
		}
		if len(tmp) == k {
			return
		}
		for i := pos; i < 10; i++ {
			if reset < i {
				return
			}
			tmp = append(tmp, i)
			dfs(i+1, reset-i)
			tmp = tmp[:len(tmp)-1]
		}
	}
	dfs(1, n)
	return ans
}
