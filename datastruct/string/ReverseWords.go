/*
* @Author: Knight
* @Date:   2020/8/30 19:26
* @Email:  knight2347@163.com
* @Ideal:  leetcode557 翻转字符串中的单词
 */

package string

func reverseWords(s string) string {
	preIndex := 0
	var ret string
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			ret += reverse(s[preIndex:i]) + " "
			preIndex = i + 1
		}
		if i == len(s)-1 {
			ret += reverse(s[preIndex:])
		}
	}
	return ret
}

func reverse(s string) string {
	var reversed string
	for _, v := range s {
		reversed = string(v) + reversed
	}
	return reversed
}
