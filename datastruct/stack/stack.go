//Author  :     knight
//Date    :     2020/06/26 15:10:00
//Version :     1.0
//Email   :     knight2347@163.com
//idea    :     栈

package stack

//Stack 栈
type Stack struct {
	data   []int
	length int
}

//GetStack 创建一个新的栈
func GetStack() Stack {
	return Stack{
		data:   make([]int, 0),
		length: 0,
	}
}

//Push 压入栈
func (s *Stack) Push(item int) {
	s.data = append(s.data, item)
	s.length++
}

//Top 获取栈顶元素
func (s *Stack) Top() int {
	if !s.Empty() {
		return s.data[s.length-1]
	}
	return -1
}

//Pop 弹出栈顶元素并返回
func (s *Stack) Pop() int {
	if !s.Empty() {
		tmp := s.data[0]
		s.data = s.data[:s.length-1]
		s.length--
		return tmp
	}
	return -1
}

//Empty 栈是否为空
func (s *Stack) Empty() bool {
	return s.length == 0
}
