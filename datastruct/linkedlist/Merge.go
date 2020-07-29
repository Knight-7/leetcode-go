//Author  :     knight
//Date    :     2020/07/04 23:26:52
//Version :     1.0
//Email   :     knight2347@163.com
//idea    :     合并两个有序链表

package linkedlist

//Merge 合并两个有序的链表
func Merge(list1, list2 *LinkedList) *LinkedList {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	head := &LinkedList{-1, nil}
	mergeNode := head
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			mergeNode.Next = list1
			list1 = list1.Next
			mergeNode = mergeNode.Next
			if list1 == nil {
				mergeNode.Next = list2
				return head.Next
			}
		} else {
			mergeNode.Next = list2
			list2 = list2.Next
			mergeNode = mergeNode.Next
			if list2 == nil {
				mergeNode.Next = list1
				return head.Next
			}
		}
	}
	
	return head.Next
}

//MergeRecursive 合并有序链表(递归)
func MergeRecursive(list1, list2 *LinkedList) *LinkedList {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	var mergeNode *LinkedList
	if list1.Val < list2.Val {
		mergeNode = list1
		mergeNode.Next = MergeRecursive(list1.Next, list2)
	} else {
		mergeNode = list2
		mergeNode.Next = MergeRecursive(list1, list2.Next)
	}

	return mergeNode
}