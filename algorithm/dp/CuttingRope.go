//Author  :     knight
//Date    :     2020/07/09 09:15:11
//Version :     1.0
//Email   :     knight2347@163.com
//idea    :     剪绳子

package dp

//CuttingRope 剪绳子
func CuttingRope(n int) int {
	if n == 1 {
		return 0
	}
	if n == 2 {
		return 1
	}
	if n == 3 {
		return 2
	}

	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	dp[3] = 3
	for i := 4; i <= n; i++ {
		for j := 1; j <= i/2; j++ {
			dp[i] = max(dp[i], dp[j]*dp[i-j])
		}
	}

	return dp[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
