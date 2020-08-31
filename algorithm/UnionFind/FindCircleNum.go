/*
* @Author: Knight
* @Date:   2020/8/21 17:10
* @Email:  knight2347@163.com
* @Ideal:  朋友圈:找到有多少个朋友圈
 */

package UnionFind

import "fmt"

var pre []int

func findCircleNum(M [][]int) int {
	ret := 0
	pre = make([]int, len(M))
	for i := 0; i < len(pre); i++ {
		pre[i] = i
	}
	for i := 0; i < len(M); i++ {
		for j := 0; j < len(M); j++ {
			if M[i][j] == 1 {
				union(i, j)
			}
		}
	}
	for i := 0; i < len(M); i++ {
		if i == pre[i] {
			ret++
		}
	}
	fmt.Println(pre)
	return ret
}

//递归法
func find(x int) int {
	if x == pre[x] {
		return x
	}
	pre[x] = find(pre[x])
	return pre[x]
}

//非递归法
func findx(x int) int {
	r := x
	for r != pre[r] {
		r = pre[r]
	}
	i, j := x, 0
	for i != r {
		j = pre[i]
		pre[i] = r
		i = j
	}
	return r
}

func union(x, y int) {
	xx := find(x)
	yy := find(y)
	if xx != yy {
		pre[xx] = yy
	}
}
