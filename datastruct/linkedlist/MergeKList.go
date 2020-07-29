//Author  :     knight
//Date    :     2020/07/09 08:09:56
//Version :     1.0
//Email   :     knight2347@163.com
//idea    :     合并K个有序链表

package linkedlist

//MergeKList 合并K个有序链表(采用将K个链表转换为K-1个链表两两合并)
func MergeKList(lists []*LinkedList) *LinkedList {
	if lists == nil || len(lists) == 0 {
		return nil
	}

	mergeHead := &LinkedList{-1, nil}
	for _, v := range lists {
		mergeHead.Next = Merge(mergeHead.Next, v)
	}
	return mergeHead.Next
}

//MergeKListRecursive 合并K个有序链表(采用分治法)
func MergeKListRecursive(lists []*LinkedList) *LinkedList {
	if lists == nil || len(lists) <= 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}

	var mergeNode *LinkedList
	m := len(lists) >> 1
	mergeNode = Merge(MergeKListRecursive(lists[0:m]), MergeKListRecursive(lists[m:]))

	return mergeNode
}