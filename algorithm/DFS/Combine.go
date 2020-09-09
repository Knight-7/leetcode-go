/*
* @Author: Knight
* @Date:   2020/9/8 9:22
* @Email:  knight2347@163.com
* @Ideal:
 */

package DFS

func combine(n int, k int) [][]int {
	ans := make([][]int, 0)
	var dfs func(int, []int)
	dfs = func(curIndex int, tmp []int) {
		if len(tmp) == k {
			t := make([]int, k)
			copy(t, tmp)
			ans = append(ans, t)
			return
		}
		for i := curIndex; i <= n; i++ {
			tmp = append(tmp, i)
			dfs(i+1, tmp)
			tmp = tmp[:len(tmp)-1]
		}
	}
	dfs(1, []int{})
	return ans
}
