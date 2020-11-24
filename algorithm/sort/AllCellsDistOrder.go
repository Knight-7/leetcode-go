/*
* @Author  :     knight
* @Date    :     2020/11/17 19:32:35
* @Version :     1.0
* @Email   :     knight2347@163.com
* @idea    :     leetcode 1030 直接进行排序
 */

package sort

import "sort"

type Node struct {
	x, y int
	dis  int
}

func allCellsDistOrder(R int, C int, r0 int, c0 int) [][]int {
	nodes := make([]Node, 0)
	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			nodes = append(nodes, Node{
				x:   i,
				y:   j,
				dis: abs(i-r0) + abs(j-c0),
			})
		}
	}
	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].dis < nodes[j].dis
	})
	ans := make([][]int, 0)
	for _, node := range nodes {
		ans = append(ans, []int{node.x, node.y})
	}
	return ans
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
