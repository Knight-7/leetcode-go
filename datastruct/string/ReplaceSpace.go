/*
* @Author: Knight
* @Date:   2020/8/17 20:14
* @Email:  knight2347@163.com
* @Ideal:
 */

package string

import "strings"

func replaceSpace(s string) string {
	preIndex := 0
	ans := []string{}
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			ans = append(ans, s[preIndex:i], "%20")
			preIndex = i + 1
		}
		if i == len(s)-1 && s[i] != ' ' {
			ans = append(ans, s[preIndex:i+1])
		}
	}
	return strings.Join(ans, "")
}
