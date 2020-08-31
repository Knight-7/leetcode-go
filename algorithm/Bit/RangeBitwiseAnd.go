/*
* @Author: Knight
* @Date:   2020/8/23 11:32
* @Email:  knight2347@163.com
* @Ideal:
 */

package Bit

func rangeBitwiseAnd(m, n int) int {
	cnt := 0
	for m != n {
		m >>= 1
		n >>= 1
		cnt++
	}
	return m << cnt
}
