/*
* @Author: Knight
* @Date:   2020/8/10 8:50
* @Email:  knight2347@163.com
* @Ideal:
 */

package string

func countBinarySubstrings(s string) int {
	ret := 0
	if len(s) == 1 {
		return 0
	}

	lastCnt, preCnt := -1, 1
	for i := 1; i < len(s); i++ {
		if s[i] != s[i-1] {
			if lastCnt != -1 {
				ret += min(lastCnt, preCnt)
				lastCnt = preCnt
				preCnt = 0
			} else {
				lastCnt = preCnt
				preCnt = 0
			}
		}
		preCnt++

		if i == len(s)-1 && lastCnt != -1 {
			ret += min(lastCnt, preCnt)
		}
	}

	return ret
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
