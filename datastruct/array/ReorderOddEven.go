//Author  :     knight
//Date    :     2020/07/10 16:49:15
//Version :     1.0
//Email   :     knight2347@163.com
//idea    :     调整数组顺序使奇数位于偶数前面

package array

//ReorderOddEven 调整数组顺序使奇数位于偶数前面
func ReorderOddEven(array []int) {
	if len(array) == 0 || array == nil {
		return
	}
	length := len(array)

	pBefore, pEnd := 0, length-1
	for pBefore < pEnd {
		// 前一个指针向后找第一个偶数
		for pBefore < length && array[pBefore]%2 == 1 {
			pBefore++
		}
		// 后一个指针向前找第一个奇数
		for pEnd > 0 && array[pEnd]%2 != 1 {
			pEnd--
		}

		if pBefore < pEnd {
			tmp := array[pBefore]
			array[pBefore] = array[pEnd]
			array[pEnd] = tmp
		}
	}
}
