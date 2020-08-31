/*
* @Author: Knight
* @Date:   2020/8/5 10:42
* @Email:  knight2347@163.com
* @Ideal:  将罗马数字转化为整数
 */

package string

func mdMap(m map[string]int) {
	m["I"] = 1
	m["V"] = 5
	m["X"] = 10
	m["L"] = 50
	m["C"] = 100
	m["D"] = 500
	m["M"] = 1000
	m["IV"] = 4
	m["IX"] = 9
	m["XL"] = 40
	m["XC"] = 90
	m["CD"] = 400
	m["CM"] = 900
}

func romanToInt(s string) int {
	m := make(map[string]int)
	mdMap(m)
	return getInt(s, m)
}

func getInt(s string, m map[string]int) int {
	if len(s) == 0 {
		return 0
	}
	ch := string(s[0])
	if ch == "I" || ch == "X" || ch == "C" {
		if len(s) > 1 {
			tmpCh := string(s[:2])
			if tmpCh == "IV" || tmpCh == "IX" || tmpCh == "XL" || tmpCh == "XC" ||
				tmpCh == "CD" || tmpCh == "CM" {
				return m[tmpCh] + getInt(s[2:], m)
			}
		}
	}
	return m[ch] + getInt(s[1:], m)
}
