//Author  :     knight
//Date    :     2020/07/10 11:44:02
//Version :     1.0
//Email   :     knight2347@163.com
//idea    :     用两个栈来实现队列

package stack

type queueStack struct {
	stack1 []int
	stack2 []int
}

func (q *queueStack) push(x int) {
	q.stack1 = append(q.stack1, x)
}

func (q *queueStack) pop() int {
	if len(q.stack2) == 0 {
		for len(q.stack1) > 0 {
			q.stack2 = append(q.stack2, q.stack1[len(q.stack1)-1])
			q.stack1 = q.stack1[:len(q.stack1)-1]
		}
	}
	if len(q.stack2) == 0 {
		return -1
	}
	tail := q.stack2[len(q.stack2)-1]
	q.stack2 = q.stack2[:len(q.stack2)-1]
	return tail
}