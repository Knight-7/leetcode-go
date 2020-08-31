/*
* @Author: Knight
* @Date:   2020/8/25 17:14
* @Email:  knight2347@163.com
* @Ideal:  解数独
 */

package array

import "fmt"

var (
	nodes        []Node
	row, col, cr [9][10]bool
)

type Node struct {
	x, y int
}

func (n Node) getCr() int {
	return n.x/3*3 + n.y/3
}

func solveSudoku(board [][]byte) {
	nodes = make([]Node, 0)
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Println(board[i][j])
			if board[i][j] == '.' {
				nodes = append(nodes, Node{i, j})
			} else {
				v := int(board[i][j] - '0')
				row[i][v] = true
				row[j][v] = true
				cr[i/3*3][j/3] = true
			}
		}
	}
	DFS(board, 0)
}

func DFS(board [][]byte, n int) bool {
	if n >= len(nodes) {
		return true
	}
	node := nodes[n]
	for k := 1; k < 10; k++ {
		if !row[node.x][k] && !col[node.y][k] && !cr[node.getCr()][k] {
			row[node.x][k] = true
			col[node.y][k] = true
			cr[node.getCr()][k] = true
			board[node.x][node.y] = byte(k)
			fmt.Println(board[node.x][node.y])
			if DFS(board, n+1) {
				return true
			}
			row[node.x][k] = false
			col[node.y][k] = false
			cr[node.getCr()][k] = false
		}
	}
	return false
}
