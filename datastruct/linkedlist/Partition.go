//Author  :     knight
//Date    :     2020/07/26 10:45:37
//Version :     1.0
//Email   :     knight2347@163.com
//idea    :     分隔链表

package linkedlist

//要保留两个节点的相对初始位置
func partition(head *ListNode, x int) *ListNode {
	hBig, hSmall := &ListNode{-1, nil}, &ListNode{-1, nil}
	headBig, headSamll := hBig, hSmall
	for head != nil {
		if head.Val < x {
			hSmall.Next = head
			hSmall = hSmall.Next
		} else {
			hBig.Next = head
			hBig = hBig.Next
		}
		head = head.Next
	}

	hSmall.Next = headBig.Next

	return headSamll.Next
}
