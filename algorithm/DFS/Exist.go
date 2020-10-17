/*
* @Author: Knight
* @Date:   2020/9/13 00:32
* @Email:  knight2347@163.com
* @Ideal:  leetcode 79 单词搜索
 */

package DFS

func exist(board [][]byte, word string) bool {
	existed := false
	vis := make([][]bool, len(board))
	for i := range vis {
		vis[i] = make([]bool, len(board[0]))
	}
	var dfs func(pos, x, y int) bool
	dfs = func(pos, x, y int) bool {
		if x < 0 || x >= len(board) || y < 0 || y >= len(board[0]) || vis[x][y] {
			return false
		}
		if pos == len(word)-1 && board[x][y] == word[pos] {
			return true
		}
		if word[pos] == board[x][y] {
			vis[x][y] = true
			tmpExisted := dfs(pos+1, x+1, y) || dfs(pos+1, x-1, y) ||
				dfs(pos+1, x, y+1) || dfs(pos+1, x, y-1)
			vis[x][y] = false
			return tmpExisted
		}
		return false
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == word[0] {
				existed = dfs(0, i, j)
			}
			if existed {
				return existed
			}
		}
	}
	return existed
}
