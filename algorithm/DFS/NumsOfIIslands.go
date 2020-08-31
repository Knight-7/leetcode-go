/*
* @Author: Knight
* @Date:   2020/8/7 9:03
* @Email:  knight2347@163.com
* @Ideal:  求岛屿的数量，DFS深度优先搜索
 */

package DFS

func numIslands(grid [][]byte) int {
	ret := 0
	for k1, v1 := range grid {
		for k, v := range v1 {
			if v == '1' {
				ret++
				dfs(k1, k, grid)
			}
		}
	}
	return ret
}

func dfs(x, y int, grid [][]byte) {
	grid[x][y] = '0'
	for _, next := range [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
		nextX := x + next[0]
		nextY := y + next[1]
		if nextX >= 0 && nextX < len(grid) && nextY >= 0 && nextY < len(grid[0]) && grid[nextX][nextY] == '1' {
			dfs(nextX, nextY, grid)
		}
	}
}
