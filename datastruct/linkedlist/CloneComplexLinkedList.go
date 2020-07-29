//Author  :     knight
//Date    :     2020/07/05 14:35:28
//Version :     1.0
//Email   :     knight2347@163.com
//idea    :     复杂链表的复制

package linkedlist

//CloneComplexLinkedListWithMap 使用哈希表来复制 时间复杂度o(n) 空间复杂度o(n)
func CloneComplexLinkedListWithMap(head *ComplexLinkedList) *ComplexLinkedList {
	if head == nil {
		return nil
	}

	nodeMap := make(map[*ComplexLinkedList]*ComplexLinkedList)
	cloneList := &ComplexLinkedList{-1, nil, nil}
	var tmpNode *ComplexLinkedList
	var tmpCloneNode *ComplexLinkedList

	nodeMap[nil] = nil
	tmpNode = head
	tmpCloneNode = cloneList
	for tmpNode != nil {
		tmpCloneNode.Next = &ComplexLinkedList{
			Val:     tmpNode.Val,
			Next:    tmpNode.Next,
			Sibling: nil,
		}
		nodeMap[tmpNode] = tmpCloneNode.Next
		tmpNode = tmpNode.Next
		tmpCloneNode = tmpCloneNode.Next
	}

	tmpNode = head
	tmpCloneNode = cloneList.Next
	for tmpCloneNode != nil {
		siblingNode := nodeMap[tmpNode.Sibling]
		tmpCloneNode.Sibling = siblingNode
		tmpNode = tmpNode.Next
		tmpCloneNode = tmpCloneNode.Next
	}

	return cloneList.Next
}

//CloneComplexLinkedList 复杂链表的复制 时间复杂度o(n)
func CloneComplexLinkedList(head *ComplexLinkedList) *ComplexLinkedList {
	if head == nil {
		return nil
	}

	//在每个节点后面添加自己的复制节点
	tmpNode := head
	for tmpNode != nil {
		cloneNode := &ComplexLinkedList{tmpNode.Val, tmpNode.Next, nil}
		tmpNode.Next = cloneNode
		tmpNode = cloneNode.Next
	}

	//节点N的Sibling指向S，那么节点N的复制N'的Sibling指向S节点的复制S'
	tmpNode = head
	for tmpNode != nil {
		cloneNode := tmpNode.Next
		if tmpNode.Sibling != nil {
			cloneNode.Sibling = tmpNode.Sibling.Next
		}
		tmpNode = cloneNode.Next
	}

	//连接每个节点的复制节点，就是我们所复制的链表
	cloneHead := &ComplexLinkedList{-1, nil, nil}
	cloneNode := cloneHead
	tmpNode = head
	for tmpNode != nil {
		cloneNode.Next = tmpNode.Next
		cloneNode = cloneNode.Next
		tmpNode.Next = cloneNode.Next
		tmpNode = tmpNode.Next
	}

	return cloneHead.Next
}
