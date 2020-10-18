package linkedlist

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return head
	}
	preNode, backNode := &ListNode{Next: head}, &ListNode{Next: head}
	for n > 0 {
		preNode = preNode.Next
		n--
	}
	for preNode.Next != nil {
		preNode = preNode.Next
		backNode = backNode.Next
	}
	if backNode.Next == head {
		return head.Next
	}
	backNode.Next = backNode.Next.Next
	return head
}
