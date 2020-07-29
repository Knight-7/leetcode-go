//Author  :     knight
//Date    :     2020/06/25 18:02:36
//Version :     1.0
//Email   :     knight2347@163.com
//idea    :     链表

package linkedlist

//LinkedList 链表
type LinkedList struct {
	Val  int
	Next *LinkedList
}

//ComplexLinkedList 复杂链表
type ComplexLinkedList struct {
	Val     int
	Next    *ComplexLinkedList
	Sibling *ComplexLinkedList
}

//GenerateList 根据切片生成一个链表
func GenerateList(nums []int) *LinkedList {
	if len(nums) == 0 {
		return nil
	}
	head := &LinkedList{nums[0], nil}
	tmp := head
	for i := 1; i < len(nums); i++ {
		node := &LinkedList{nums[i], nil}
		tmp.Next = node
		tmp = tmp.Next
	}
	return head
}
