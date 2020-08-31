/*
* @Author: Knight
* @Date:   2020/8/24 8:29
* @Email:  knight2347@163.com
* @Ideal:  leetcode 459 重复的子字符串
 */

package string

func repeatedSubstringPattern(s string) bool {
	length := len(s)
	for i := 1; i*2 <= length; i++ {
		if length%i == 0 {
			match := true
			for j := i; j < length; j++ {
				if s[j] != s[j-i] {
					match = false
					break
				}
			}
			if match {
				return true
			}
		}
	}
	return false
}
