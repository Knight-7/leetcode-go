//Author  :     knight
//Date    :     2020/07/03 20:54:03
//Version :     1.0
//Email   :     knight2347@163.com
//idea    :     删除链表的节点

package linkedlist

//DeleteNode 删除链表中的节点
func DeleteNode(head, BeDeleteNode *ListNode) {
	if BeDeleteNode == nil || head == nil {
		return
	}

	if head == BeDeleteNode {
		head = nil
	}

	if BeDeleteNode.Next != nil {
		BeDeleteNode.Val = BeDeleteNode.Next.Val
		BeDeleteNode.Next = BeDeleteNode.Next.Next
	} else {
		tmp := head
		for tmp.Next != BeDeleteNode {
			tmp = tmp.Next
		}
		tmp.Next = nil
	}
}
