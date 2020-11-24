/*
* @Author  :     knight
* @Date    :     2020/10/20 20:24:41
* @Version :     1.0
* @Email   :     knight2347@163.com
* @idea    :     leetcode143 重排列表，使用双指针法，且第一个指针的速度是第二个的两倍
 */

package linkedlist

func reorderList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	preList, lastList := make([]*ListNode, 0), make([]*ListNode, 0)
	preNode, lastNode := &ListNode{Next: head}, &ListNode{Next: head}
	for preNode != nil && preNode.Next != nil {
		preNode = preNode.Next.Next
		lastNode = lastNode.Next
		preList = append(preList, lastNode)
	}
	for lastNode.Next != nil {
		lastNode = lastNode.Next
		lastList = append(lastList, lastNode)
	}
	if len(preList) == 1 && len(lastList) == 0 {
		return head
	}

	ans := &ListNode{Next: nil}
	tmp := ans
	p1, p2 := 0, len(lastList)-1
	for {
		tmp.Next = preList[p1]
		p1++
		tmp = tmp.Next

		tmp.Next = lastList[p2]
		p2--
		tmp = tmp.Next
		tmp.Next = nil
		if p2 < 0 {
			break
		}
	}
	if len(preList) != len(lastList) {
		tmp.Next = preList[p1]
		tmp.Next.Next = nil
	}
	return ans.Next
}
