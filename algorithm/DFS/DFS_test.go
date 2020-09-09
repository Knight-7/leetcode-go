/*
* @Author: Knight
* @Date:   2020/8/7 9:37
* @Email:  knight2347@163.com
* @Ideal:
 */

package DFS

import "testing"

func TestIslandPerimeter(t *testing.T) {
	grid := [][]int{{1, 1}}
	ret := islandPrimeter(grid)
	t.Log(ret)
}

func TestSolve(t *testing.T) {
	board := [][]byte{{'X', 'O', 'X'}, {'X', 'O', 'X'}, {'X', 'O', 'X'}}
	solve(board)
	t.Log(board)
}

func TestImage(t *testing.T) {
	image := [][]int{{0, 0, 0}, {0, 1, 1}}
	image = floodFill(image, 1, 1, 1)
	t.Log(image)
}

func TestDFS_Room(t *testing.T) {
	rooms := [][]int{{1}, {1, 1}}
	t.Log(canVisitAllRooms(rooms))
}

func TestCombine(t *testing.T) {
	ans := combine(4, 3)
	t.Log(ans)
}
