//Author  :     knight
//Date    :     2020/07/05 14:19:29
//Version :     1.0
//Email   :     knight2347@163.com
//idea    :     逆序打印链表

package linkedlist

import "fmt"

//PrintLikedListInReverseOrderWithStack 使用栈逆序打印链表
func PrintLikedListInReverseOrderWithStack(head *ListNode) {
	if head == nil {
		return
	}

	stack := make([]*ListNode, 0)
	for head != nil {
		stack = append(stack, head)
		head = head.Next
	}

	for len(stack) > 0 {
		top := stack[len(stack)-1]
		fmt.Printf("%d ", top.Val)
		stack = stack[:len(stack)-1]
	}
}

//PrintLikedListInReverseOrderRecurse 递归逆序打印链表
func PrintLikedListInReverseOrderRecurse(head *ListNode) {
	if head == nil {
		return
	}
	PrintLikedListInReverseOrderRecurse(head.Next)
	fmt.Printf("%d ", head.Val)
}
