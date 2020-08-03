//Author  :     knight
//Date    :     2020/07/04 16:42:54
//Version :     1.0
//Email   :     knight2347@163.com
//idea    :     查找链表中环的入口

package linkedlist

//MeetNode 判断链表中有没有环，如果有环，那么MeetNode就不为空，否则MeetNode就为空(双指针)
func MeetNode(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	slow, fast := head, head.Next
	for slow != nil && fast != nil {
		if slow == fast {
			return fast
		}

		slow = slow.Next
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
		}
	}

	return nil
}

//EntryNodeOfLoop 查找链表中的环的入口节点(双指针)
func EntryNodeOfLoop(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	meetNode := MeetNode(head)
	if meetNode == nil {
		return nil
	}
	sumNodeOfLoop := 1
	tmp := meetNode.Next
	for tmp != meetNode {
		sumNodeOfLoop++
		tmp = tmp.Next
	}

	slow, fast := head, head
	for sumNodeOfLoop > 0 {
		fast = fast.Next
		sumNodeOfLoop--
	}
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}

	return fast
}
