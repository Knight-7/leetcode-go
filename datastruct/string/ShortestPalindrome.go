/*
* @Author: Knight
* @Date:   2020/8/29 11:52
* @Email:  knight2347@163.com
* @Ideal:  leetcode214 最短回文串
 */

package string

func shortestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	} else if len(s) == 1 {
		return s
	} else if len(s) == 2 {
		return string(s[1]) + s
	}
	mid := len(s) / 2
	for i := mid; i >= 0; i-- {
		if isEqual(s[:i], s[i+1:]) {
			ans := ""
			for j := i + 1; j < len(s); j++ {
				ans = string(s[j]) + ans
			}
			return ans + s[i:]
		}
		if i > 0 && s[i] == s[i-1] && isEqual(s[:i-1], s[i+1:]) {
			ans := ""
			for j := i + 1; j < len(s); j++ {
				ans = string(s[j]) + ans
			}
			return ans + s[i-1:]
		}
	}
	return ""
}

//判断s1的逆序和s2是否相等
func isEqual(s1, s2 string) bool {
	if len(s2) < len(s1) {
		return false
	}
	index1, index2 := len(s1)-1, 0
	for index1 >= 0 {
		if s1[index1] != s2[index2] {
			return false
		}
		index1--
		index2++
	}
	return true
}
