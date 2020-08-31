package string

//dp
/*
状态定义： 设动态规划列表 dp ，dp[j] 代表以字符 s[j] 为结尾的 “最长不重复子字符串” 的长度。

转移方程： 固定右边界 j ，设字符 s[j] 左边距离最近的相同字符为 s[i] ，即 s[i] = s[j] 。

当 i < 0 ，即 s[j] 左边无相同字符，则 dp[j] = dp[j-1] + 1 ；
当 dp[j - 1] < j - i ，说明字符 s[i] 在子字符串 dp[j-1] 区间之外 ，则 dp[j] = dp[j - 1] + 1 ；
当 dp[j−1]≥j−i ，说明字符 s[i]在子字符串 dp[j-1] 区间之中 ，则 dp[j] 的左边界由 s[i] 决定，即 dp[j] = j - i ；
*/
func LongestSubstringWithoutDuplication(str string) int {
	curLength := 0                 //表示以当前节点作为最后节点时子串的长度
	maxLength := 0                 //最长子串的长度
	position := make(map[rune]int) //记录节点上一次出现的位置
	for k, v := range str {
		preIndex, ok := position[v] //获取上一次出现的位置
		//preIndex<0表示改字符前面没出现过
		//k-preIndex>curIndex表示该字符上一次出现的在当前子串的前面，即当前子串不包含该字符
		if !ok || k-preIndex > curLength {
			curLength++
		} else {
			//表示当前字符在当前子串中已经出现过了，那么判断它是否为最长子串
			if curLength > maxLength {
				maxLength = curLength
			}
			//更新当前子串的长度
			curLength = k - preIndex
		}
		//更新当前字符的位置
		position[v] = k
	}
	if curLength > maxLength {
		maxLength = curLength
	}

	return maxLength
}
