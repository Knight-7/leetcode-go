//Author  :     knight
//Date    :     2020/06/26 09:36:56
//Version :     1.0
//Email   :     knight2347@163.com
//idea    :     队列

package queue

import "errors"

//Queue 队列
type Queue struct {
	Data   []interface{}
	length int
}

//GetQueue 返回一个空的队列
func GetQueue() Queue {
	return Queue{
		Data:   make([]interface{}, 0),
		length: 0,
	}
}

//Push 压入队列
func (q *Queue) Push(data interface{}) {
	q.Data = append(q.Data, data)
	q.length++
}

//Top 获取队头元素
func (q *Queue) Top() (interface{}, error) {
	if !q.Empty() {
		return q.Data[0], nil
	}
	return 0, errors.New("Queue is empty")
}

//Pop 弹出队头元素并返回
func (q *Queue) Pop() (interface{}, error) {
	if !q.Empty() {
		top := q.Data[0]
		q.Data = q.Data[1:]
		q.length--
		return top, nil
	}
	return 0, errors.New("Queue is empty")
}

//Empty 对了是否为空
func (q *Queue) Empty() bool {
	return q.length == 0
}

//Length 队列长度
func (q *Queue) Length() int {
	return q.length
}

func main() {
	q := Queue{
		Data:   make([]interface{}, 0),
		length: 0,
	}
	q.Push(1)
}
