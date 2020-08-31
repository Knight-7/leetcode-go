/*
* @Author: Knight
* @Date:   2020/8/6 9:37
* @Email:  knight2347@163.com
* @Ideal:
 */

package string

func palindromePairs(words []string) [][]int {
	ret := make([][]int, 0)
	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			if isPalindrome(words[i] + words[j]) {
				ret = append(ret, []int{i, j})
			}
			if isPalindrome(words[j] + words[i]) {
				ret = append(ret, []int{j, i})
			}
		}
	}
	return ret
}

func isPalindrome(str string) bool {
	left, right := 0, len(str)-1
	for {
		if left >= right {
			break
		}
		if str[left] != str[right] {
			return false
		}
		left++
		right--
	}
	return true
}
