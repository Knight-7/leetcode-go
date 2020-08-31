/*
* @Author: Knight
* @Date:   2020/8/21 17:27
* @Email:  knight2347@163.com
* @Ideal:
 */

package UnionFind

import "testing"

func TestFindCircle(t *testing.T) {
	m := [][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}}
	ret := findCircleNum(m)
	t.Log(ret)
}
