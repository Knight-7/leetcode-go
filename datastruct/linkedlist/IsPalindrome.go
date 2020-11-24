/*
* @Author  :     knight
* @Date    :     2020/10/23 19:50:12
* @Version :     1.0
* @Email   :     knight2347@163.com
* @idea    :     判断数字链表是否回文
 */

package linkedlist

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	preNode, lastNode := &ListNode{Next: head}, &ListNode{Next: head}
	preNodes, lastNodes := make([]*ListNode, 0), make([]*ListNode, 0)
	for preNode != nil && preNode.Next != nil {
		preNode = preNode.Next.Next
		lastNode = lastNode.Next
		preNodes = append(preNodes, lastNode)
	}
	for lastNode.Next != nil {
		lastNode = lastNode.Next
		lastNodes = append(lastNodes, lastNode)
	}
	p2 := len(lastNodes) - 1
	for p1 := 0; p1 < len(lastNodes); p1++ {
		if preNodes[p1].Val != lastNodes[p2].Val {
			return false
		}
		p2--
	}
	return true
}
