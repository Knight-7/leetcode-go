/*
* @Author: Knight
* @Date:   2020/8/11 11:15
* @Email:  knight2347@163.com
* @Ideal:  一个图由X和O组成，将所有被X包围的O改为X（包围是指O组成的所有区域都是被X包围的）
*          思路：遍历整个图的边界，将所有与之相连的的O标记
 */

package DFS

var vis [][]bool

func solve(board [][]byte) {
	if len(board) == 0 {
		return
	}
	vis = make([][]bool, len(board))
	for i := 0; i < len(board); i++ {
		vis[i] = make([]bool, len(board[0]))
	}
	for i := 0; i < len(board[0]); i++ {
		if board[0][i] == 'O' && !vis[0][i] {
			dfsBoard(board, 0, i)
		}
		if board[len(board)-1][i] == 'O' && !vis[len(board)-1][i] {
			dfsBoard(board, len(board)-1, i)
		}
	}
	for i := 0; i < len(board); i++ {
		if board[i][0] == 'O' && !vis[i][0] {
			dfsBoard(board, i, 0)
		}
		if board[i][len(board[0])-1] == 'O' && !vis[i][len(board[0])-1] {
			dfsBoard(board, i, len(board[0])-1)
		}
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == 'O' && !vis[i][j] {
				board[i][j] = 'X'
			}
		}
	}
}

func dfsBoard(board [][]byte, x, y int) {
	vis[x][y] = true
	for _, v := range [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
		nx := x + v[0]
		ny := y + v[1]
		if nx >= 0 && nx < len(board) && ny >= 0 && ny < len(board[0]) && board[nx][ny] == 'O' && !vis[nx][ny] {
			dfsBoard(board, nx, ny)
		}
	}
}
