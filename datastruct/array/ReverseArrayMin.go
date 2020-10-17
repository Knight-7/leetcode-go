//Author  :     knight
//Date    :     2020/07/09 09:48:08
//Version :     1.0
//Email   :     knight2347@163.com
//idea    :     旋转数组(原数组是递增的)的最小值

package array

//ReverseArrayMin 旋转数组(原数组是递增的)的最小值
func ReverseArrayMin(numbers []int, length int) int {
	if numbers == nil || length <= 0 {
		return 0
	}

	if numbers[0] < numbers[length-1] {
		return numbers[0]
	}

	l, r := 0, length-1
	m := l
	for l < r {
		if r-l == 1 {
			return numbers[r]
		}

		if numbers[l] == numbers[m] && numbers[m] == numbers[r] {
			return minNumber(numbers, l, r)
		}

		m := (l + r) >> 1
		if numbers[m] >= numbers[l] {
			l = m
		} else if numbers[m] <= numbers[r] {
			r = m
		}
	}

	return numbers[r]
}

func minNumber(numbers []int, l, r int) int {
	ans := numbers[l]
	for i := l; i <= r; i++ {
		if ans > numbers[i] {
			ans = numbers[i]
		}
	}
	return ans
}
