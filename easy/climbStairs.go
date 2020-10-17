//Author  :     knight
//Date    :     2020/06/13 07:40:02
//Version :     1.0
//Email   :     knight2347@163.com
//idea    :      爬楼梯，dp

package main

func climbStairs(n int) int {
	dp := make([]int, n+1)
	if n == 1 {
		return 1
	}
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}
