//Author  :     knight
//Date    :     2020/07/05 16:08:15
//Version :     1.0
//Email   :     knight2347@163.com
//idea    :     查找第一个公共的节点

package linkedlist

//FindFirstCommonNode 查找第一个公共的节点
func FindFirstCommonNode(list1, list2 *LinkedList) *LinkedList {
	if list1 == nil || list2 == nil {
		return nil
	}

	//遍历获得两个链表，或得两个链表的长度
	nodesInList1, nodesInList2 := 0, 0
	tmpNode1, tmpNode2 := list1, list2
	for tmpNode1 != nil || tmpNode2 != nil {
		if tmpNode1 != nil {
			nodesInList1++
			tmpNode1 = tmpNode1.Next
		}
		if tmpNode2 != nil {
			nodesInList2++
			tmpNode2 = tmpNode2.Next
		}
	}

	tmpNode1, tmpNode2 = list1, list2
	for nodesInList1 > nodesInList2 {
		tmpNode1 = tmpNode1.Next
		nodesInList1--
	}
	for nodesInList1 < nodesInList2 {
		tmpNode2 = tmpNode2.Next
		nodesInList2--
	}

	for tmpNode1 != nil {
		if tmpNode1.Next == tmpNode2.Next {
			return tmpNode1.Next
		}
		tmpNode1 = tmpNode1.Next
		tmpNode2 = tmpNode2.Next
	}

	return nil
}