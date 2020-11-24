/*
* @Author  :     knight
* @Date    :     2020/10/19 20:57:39
* @Version :     1.0
* @Email   :     knight2347@163.com
* @idea    :
 */

package string

func backspaceCompare(S, T string) bool {
	s, t := make([]byte, len(S)), make([]byte, len(T))
	i, p1, p2 := 0, 0, 0
	for i < len(S) || i < len(T) {
		if i < len(S) {
			if S[i] == '#' {
				reduce(&p1)
			} else {
				s[p1] = S[i]
				p1++
			}
		}
		if i < len(T) {
			if T[i] == '#' {
				reduce(&p2)
			} else {
				t[p2] = T[i]
				p2++
			}
		}
		i++
	}
	if len(s[:p1]) != len(t[:p2]) {
		return false
	}
	for i := 0; i < len(s[:p1]); i++ {
		if s[i] != t[i] {
			return false
		}
	}
	return true
}

func reduce(n *int) {
	if *n > 0 {
		*n--
	}
}
