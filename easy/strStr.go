//Author  :     knight
//Date    :     2020/06/21 16:11:29
//Version :     1.0
//Email   :     knight2347@163.com
//idea    :     

package main

func strStr(haystack string, needle string) int {
	haystackLen, needleLen := len(haystack), len(needle)

	if needleLen == 0 {
		return 0
	}

	for i := 0; i < haystackLen - needleLen + 1; i++ {
		if haystack[i:i+needleLen] == needle {
			return i
		}
	}
	return -1
}