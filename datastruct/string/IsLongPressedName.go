/*
* @Author  :     knight
* @Date    :     2020/10/21 19:19:59
* @Version :     1.0
* @Email   :     knight2347@163.com
* @idea    :     leetcode 925. 长按键入 思路：双指针法
 */

package string

func isLongPressedName(name, typed string) bool {
	var p1, p2 int
	for p2 < len(typed) {
		if p1 < len(name) && name[p1] == typed[p2] {
			p1++
			p2++
		} else if p2 >= 1 && typed[p2] == typed[p2-1] {
			p2++
		} else {
			return false
		}
	}
	return p1 == len(name)
}
