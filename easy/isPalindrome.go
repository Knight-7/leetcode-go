//Author  :     knight
//Date    :     2020/06/10 21:27:59
//Version :     1.0
//Email   :     knight2347@163.com
//idea    :     leetcode判断是否是回文数

package main

import (
	"fmt"
	"strconv"
)

//方法一：将数字转换为字符串，然后将前后对应的数字一一比较
func isPalindrome(x int) bool {
	str := strconv.Itoa(x)
	l, r := 0, len(str) - 1
	for {
		fmt.Println(l, r)
		if str[l] != str[r] {
			return false
		}
		l++
		r--
		if l >= r {
			break
		}
	}
	return true
}

//方法二：翻转一般的数字，观察是否相等，然后比较
//相比于方法一；方法二减少了转化字符串的消耗
func isPalindrome2(x int) bool {
	if x == 0 {
		return true
	}
	//负数和个位数为0的，肯定不是回文串
	if x < 0 || x % 10 == 0 {
		return false
	}
	
	num := 0
	for x > num {
		num = num * 10 + x % 10
		x /= 10
	}

	return x == num || x == num / 10
}

func main() {
	fmt.Println(isPalindrome2(0))
}