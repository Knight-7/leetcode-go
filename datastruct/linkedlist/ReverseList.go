//Author  :     knight
//Date    :     2020/07/04 16:04:31
//Version :     1.0
//Email   :     knight2347@163.com
//idea    :     翻转链表

package linkedlist

//ReverseLinkedList 翻转一个链表（使用栈）
func ReverseLinkedList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	stack := make([]*ListNode, 0)
	for head.Next != nil {
		stack = append(stack, head)
		head = head.Next
	}

	ans := head
	for len(stack) > 0 {
		head.Next = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		head = head.Next
	}
	//将头结点的next指向nil，否则的话头结点指向第二个节点，第二个节点指向头结点，会造成死循环
	head.Next = nil

	return ans
}

func reverseLinkedListWithoutStack2(head *ListNode) (*ListNode, *ListNode) {
	if head != nil && head.Next != nil {
		pre, ans := reverseLinkedListWithoutStack2(head.Next)
		pre.Next = head
		head.Next = nil
		return head, ans
	}
	return head, head
}

//ReverseLinkedListWithoutStack 翻转一个链表(递归)
func ReverseLinkedListWithoutStack(head *ListNode) *ListNode {
	_, ans := reverseLinkedListWithoutStack2(head)
	return ans
}

//ReverseLinkedListIterate 翻转链表(迭代)
func ReverseLinkedListIterate(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var reverseHead *ListNode
	var preNode *ListNode
	node := head
	for node != nil {
		nextNode := node.Next

		if node.Next == nil {
			reverseHead = node
		}

		node.Next = preNode
		preNode = node
		node = nextNode
	}

	return reverseHead
}
