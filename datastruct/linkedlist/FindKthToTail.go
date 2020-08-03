//Author  :     knight
//Date    :     2020/07/03 21:07:13
//Version :     1.0
//Email   :     knight2347@163.com
//idea    :     链表的倒数第K节点

package linkedlist

//FindKthToTail 输出链表的倒数第K个节点
func FindKthToTail(head *ListNode, k int) *ListNode {
	if head == nil || k < 0 {
		return nil
	}

	pre, back := head, head
	count := 0
	for count < k-1 {
		if pre.Next != nil {
			pre = pre.Next
			count++
		} else {
			return nil
		}
	}
	if pre == nil {
		return nil
	}
	for pre.Next != nil {
		pre = pre.Next
		back = back.Next
	}

	return back
}
