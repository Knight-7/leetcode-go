//Author  :     knight
//Date    :     2020/06/08 21:39:29
//Version :     1.0
//Email   :     knight2347@163.com
//idea    :     并查集

package main

func equationsPossible(equations []string) bool {
	pre := make([]int, 26)
	for i := 0; i < len(pre); i++ {
		pre[i] = i
	}

	for _, str := range equations {
		if str[1] == '=' {
			x1 := int(str[0] - 'a')
			x2 := int(str[3] - 'a')
			union(x1, x2, pre)
		}
	}

	for _, str := range equations {
		if str[1] == '!' {
			x1 := int(str[0] - 'a')
			x2 := int(str[3] - 'a')
			if find(x1, pre) == find(x2, pre) {
				return false
			}
		}
	}

	return true
}

func find(x int, pre []int) int { //查找根节点
	for x != pre[x] {
		pre[x] = pre[pre[x]]
		x = pre[x]
	}
	return x
}

func union(x, y int, pre []int) {
	pre[find(x, pre)] = find(y, pre)
}
