/*
* @Author: Knight
* @Date:   2020/8/28 20:06
* @Email:  knight2347@163.com
* @Ideal:  leetcode 657 机器人能否返回原点
 */

package string

func judgeCircle(moves string) bool {
	h, a := 0, 0
	for _, v := range moves {
		switch v {
		case 'U':
			a++
		case 'D':
			a--
		case 'L':
			h--
		case 'R':
			h++
		}
	}
	return h == 0 && a == 0
}
