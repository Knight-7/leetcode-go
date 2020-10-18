/*
* @Author  :     knight
* @Date    :     2020/10/17 16:35:06
* @Version :     1.0
* @Email   :     knight2347@163.com
* @idea    :     回溯法
 */

package DFS

func totalNQueens(n int) int {
	var ans int
	cols := make([]bool, n)
	dig1 := make([]bool, 2*n-1)
	dig2 := make([]bool, 2*n-1)
	var backtrace func(row int)
	backtrace = func(row int) {
		if row == n {
			ans++
			return
		}

		for col, hasVal := range cols {
			d1, d2 := row+n-1-col, row+col
			if hasVal || dig1[d1] || dig2[d2] {
				continue
			}
			cols[col] = true
			dig1[d1] = true
			dig2[d2] = true
			backtrace(row + 1)
			cols[col] = false
			dig1[d1] = false
			dig2[d2] = false
		}
	}
	backtrace(0)
	return ans
}
