//Author  :     knight
//Date    :     2020/06/09 20:49:09
//Version :     1.0
//Email   :     knight2347@163.com
//idea    :    	动态规划

package main

import "strconv"

func translateNum(num int) int {
	src := strconv.Itoa(num)
	p, q, r := 0, 0, 1
	for i := 0; i < len(src); i++ {
		p, q, r = q, r, 0
		r += q
		if i == 0 {
			continue
		}
		pre := src[i-1 : i+1]
		if pre <= "25" && pre >= "10" {
			r += p
		}
	}
	return r
}
