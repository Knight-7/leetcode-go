//Author  :     knight
//Date    :     2020/07/04 16:56:06
//Version :     1.0
//Email   :     knight2347@163.com
//idea    :     链表的测试

package linkedlist

import "testing"

func TestEntryNodeOfLoop(t *testing.T) {
	nums := []int{1}
	head := GenerateList(nums)
	tmp := head
	for tmp.Next != nil {
		tmp = tmp.Next
	}
	tmp.Next = head
	entryNode := EntryNodeOfLoop(head)
	if entryNode == head.Next.Next {
		t.Log("测试成功")
		t.Log(entryNode.Val)
	} else if entryNode == nil {
		t.Log("当前链表没有环")
	} else {
		t.Error("测试失败")
		t.Log(entryNode.Val)
	}
}

func TestReverseNode(t *testing.T) {
	nums := []int{7, 2, 3, 4}
	head := GenerateList(nums)
	reverseHead := ReverseLinkedListIterate(head)
	for reverseHead != nil {
		t.Logf("%d ", reverseHead.Val)
		reverseHead = reverseHead.Next
	}
}

func TestMerge(t *testing.T) {
	nums1 := []int{1, 3, 5, 7}
	nums2 := []int{1, 4, 8}
	list1 := GenerateList(nums1)
	list2 := GenerateList(nums2)
	mergeList := MergeRecursive(list1, list2)
	for mergeList != nil {
		t.Logf("%d ", mergeList.Val)
		mergeList = mergeList.Next
	}
}

func TestCloneComplexLinkedList(t *testing.T) {
	c := &ComplexLinkedList{2, nil, nil}
	c1 := &ComplexLinkedList{3, c, nil}
	c2 := &ComplexLinkedList{1, c1, c}
	clone := CloneComplexLinkedList(c2)
	clone2 := CloneComplexLinkedListWithMap(c2)
	for clone != nil {
		if clone.Sibling != nil && clone2.Sibling != nil {
			t.Logf("%d %d %d %d", clone.Val, clone.Sibling.Val, clone2.Val, clone2.Sibling.Val)
		} else {
			t.Logf("%d %d", clone.Val, clone2.Val)
		}
		clone = clone.Next
		clone2 = clone2.Next
	}
}

func TestFindFirstCommonNode(t *testing.T) {
	l1 := &ListNode{4, nil}
	l2 := &ListNode{3, l1}
	l3 := &ListNode{7, l2}
	l4 := &ListNode{2, l2}
	l5 := &ListNode{7, l4}
	commonNode := FindFirstCommonNode(l3, l5)
	t.Log(commonNode.Val)
}

func TestPartition(t *testing.T) {
	nums := []int{1, 4, 3, 2, 5, 2}
	head := GenerateList(nums)
	head = partition(head, 3)
	for head != nil {
		t.Log(head.Val)
		head = head.Next
	}
}

func TestReverstKGroup(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	list := GenerateList(nums)
	ans := reverseKGroup(list, 5)
	for ans != nil {
		t.Logf("%d ", ans.Val)
		ans = ans.Next
	}
}
