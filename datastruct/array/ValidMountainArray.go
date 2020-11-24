/*
* @Author  :     knight
* @Date    :     2020/11/03 06:33:54
* @Version :     1.0
* @Email   :     knight2347@163.com
* @idea    :
 */

package array

func validMountainArray(A []int) bool {
	if len(A) < 3 {
		return false
	}
	p := 0
	for p < len(A)-1 && A[p] < A[p+1] {
		p++
	}
	if p == 0 || p == len(A)-1 {
		return false
	}
	for p < len(A)-1 && A[p] > A[p+1] {
		p++
	}

	return p == len(A)-1
}
