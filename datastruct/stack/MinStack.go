//Author  :     knight
//Date    :     2020/07/05 16:33:41
//Version :     1.0
//Email   :     knight2347@163.com
//idea    :     包含min函数的栈

package stack

//MinStack 包含min函数的栈
type MinStack struct {
	length  int
	data    []int
	mindata []int
}

//Push 压入栈
func (s *MinStack) Push(item int) {
	s.data = append(s.data, item)

	if s.length == 0 || s.mindata[s.length-1] > item {
		s.mindata = append(s.mindata, item)
	} else {
		s.mindata = append(s.mindata, s.mindata[s.length-1])
	}

	s.length++
}

//Pop 弹出栈
func (s *MinStack) Pop() int {
	if s.length <= 0 {
		return -1
	}

	s.mindata = s.mindata[:s.length-1]
	top := s.data[s.length]
	s.length--
	return top
}

//Min 获得栈中的最小元算
func (s *MinStack) Min() int {
	if s.length <= 0 {
		return -1
	}

	return s.mindata[s.length-1]
}
