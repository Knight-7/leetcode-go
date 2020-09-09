/*
* @Author: Knight
* @Date:   2020/8/31 20:52
* @Email:  knight2347@163.com
* @Ideal:  leetcode 941 钥匙和房间
*          使用深搜
 */

package DFS

func canVisitAllRooms(rooms [][]int) bool {
	cnt := 0
	n := len(rooms)
	if n == 0 || n == 1 {
		return true
	}
	visRoom := make([]bool, n)
	var DFS_Room func([][]int, int) bool
	DFS_Room = func(rooms [][]int, curRoom int) bool {
		if !visRoom[curRoom] {
			visRoom[curRoom] = true
			cnt++
		}
		if cnt == n {
			return true
		}
		for _, room := range rooms[curRoom] {
			if !visRoom[room] {
				if DFS_Room(rooms, room) {
					return true
				}
			}
		}
		return false
	}

	return DFS_Room(rooms, 0)
}
