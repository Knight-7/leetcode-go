//Author  :     knight
//Date    :     2020/07/10 14:14:42
//Version :     1.0
//Email   :     knight2347@163.com
//idea    :     正则表达式的匹配

package string

//IsMatch 正则表达式的匹配
func IsMatch(str string, pattern string) bool {
	if len(pattern) == 0 {
		return len(str) == 0
	}

	if len(pattern) > 1 && pattern[1] == '*' {
		//第一种情况忽视*和它前面的要给字符
		return IsMatch(str, pattern[2:]) ||
			//第二种情况，当模式的第一个字符和字符串的字符匹配，则字符向后移动一位（不需要字符串向后移动以为同时模式移动两位，因为这种情况已经包含在字符串向后移动一位里面了）
			(len(str) > 0 && (str[0] == pattern[0] || pattern[0] == '.')) && IsMatch(str[1:], pattern)
	} else {
		//如果pattern[1] != '*", 那么如果str[0]和patter[0]匹配，则递归比较str[1:]和pattern[1:],否则就返回false
		return len(str) > 0 && (str[0] == pattern[0] || pattern[0] == '.') && IsMatch(str[1:], pattern[1:])
	}

}

func isMatch(s string, p string) bool {
	if len(p) == 0 {
		return len(s) == 0
	}

	if len(p) > 1 && p[1] == '*' {
		return isMatch(s, p[2:]) ||
			(len(s) > 0 && (s[0] == p[0] || p[0] == '.') && isMatch(s[1:], p))
	} else {
		return len(s) > 0 && (s[0] == p[0] || p[0] == '.') && isMatch(s[1:], p[1:])
	}
}
