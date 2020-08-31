/*
* @Author: Knight
* @Date:   2020/8/19 8:05
* @Email:  knight2347@163.com
* @Ideal:  回文子串：给定一个字符串，计算这个字符串中有多少个回文子串。
*          使用中心扩展法，中心分别是奇数长度和偶数长度的。
 */

package string

func countSubstrings(s string) int {
	ans := 0
	var check func(l, r int)
	check = func(l, r int) {
		for l >= 0 && r < len(s) && s[l] == s[r] {
			l--
			r++
			ans++
		}
	}
	for i := 0; i < len(s); i++ {
		check(i, i)
		check(i, i+1)
	}
	return ans
}
