package linkedlist

//K个一组翻转链表,当节点数不是K个整数，那么剩余节点保持原来的顺序
func reverseKGroup(head *ListNode, k int) *ListNode {
	node := &ListNode{Val: -1, Next: head}
	lastNode := node
	tmp := 0
	fitst, ans := false, head
	stack := make([]*ListNode, k)
	for node.Next != nil {
		for tmp < k && node.Next != nil {
			node = node.Next
			stack[tmp] = node
			tmp++
		}
		if tmp < k {
			break
		}
		if !fitst {
			ans = node
			fitst = true
		} else {
			lastNode.Next = node
		}
		next := node.Next
		for tmp > 1 {
			tmp--
			stack[tmp].Next = stack[tmp-1]
		}
		tmp--
		stack[tmp].Next = next
		node = stack[tmp]
		lastNode = node
	}
	return ans
}
