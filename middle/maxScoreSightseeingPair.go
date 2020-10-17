//Author  :     knight
//Date    :     2020/06/17 21:42:27
//Version :     1.0
//Email   :     knight2347@163.com
//idea    :     枚举

package main

func maxScoreSightseeingPair(A []int) int {
	maxScore, mx := 0, A[0]
	n := len(A)
	for i := 1; i < n; i++ {
		maxScore = max(maxScore, A[i]-i+mx)
		mx = max(mx, A[i]+i)
	}
	return maxScore
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
