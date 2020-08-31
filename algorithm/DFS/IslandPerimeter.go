/*
* @Author: Knight
* @Date:   2020/8/7 9:27
* @Email:  knight2347@163.com
* @Ideal:  求岛屿的周长，直接遍历并计算每一个格子的周长
 */

package DFS

func islandPrimeter(grid [][]int) int {
	ret := 0
	for i, line := range grid {
		for j, v := range line {
			if v == 1 {
				for _, next := range [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
					x := i + next[0]
					y := j + next[1]
					if isOut(x, y, &grid) || grid[x][y] == 0 {
						ret++
					}
				}
			}
		}
	}
	return ret
}

func isOut(x, y int, grid *[][]int) bool {
	if x < 0 || x >= len(*grid) || y < 0 || y >= len((*grid)[0]) {
		return true
	}
	return false
}
