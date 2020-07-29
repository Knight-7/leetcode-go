//Author  :     knight
//Date    :     2020/06/26 09:29:28
//Version :     1.0
//Email   :     knight2347@163.com
//idea    :     二叉树的遍历

package binarytree

import (
	"fmt"
	"leetcode/datastruct/queue"
)

//PreorderTraverse 先序遍历
func PreorderTraverse(root *BinaryTree) {
	if root != nil {
		fmt.Printf("%d ", root.Val)
		PreorderTraverse(root.Left)
		PreorderTraverse(root.Right)
	}
}

//PreorderTravseWithoutRecursive 先序遍历非递归
func PreorderTravseWithoutRecursive(root *BinaryTree) []int {
	ans := make([]int, 0)
	stack := make([]*BinaryTree, 0)
	for len(stack) > 0 || root != nil {
		for root != nil {
			ans = append(ans, root.Val)
			stack = append(stack, root)
			root = root.Left
		}
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		root = top.Right
	}
	return ans
}

//InorderTraverse 中序遍历
func InorderTraverse(root *BinaryTree) {
	if root != nil {
		InorderTraverse(root.Left)
		fmt.Printf("%d ", root.Val)
		InorderTraverse(root.Right)
	}
}

//InorderTravseWithoutRecursive 中序遍历非递归
func InorderTravseWithoutRecursive(root *BinaryTree) []int {
	ans := make([]int, 0)
	stack := make([]*BinaryTree, 0)
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ans = append(ans, top.Val)
		root = top.Right
	}
	return ans
}

//PostorderTraverse 后序遍历
func PostorderTraverse(b *BinaryTree) {
	if b != nil {
		PostorderTraverse(b.Left)
		PostorderTraverse(b.Right)
		fmt.Printf("%d ", b.Val)
	}
}

//PostorderTravseWithoutRecursive 后序遍历非递归
func PostorderTravseWithoutRecursive(b *BinaryTree) []int {
	ans := make([]int, 0)
	stack := make([]*BinaryTree, 0)
	var lastvisited *BinaryTree
	for len(stack) > 0 || b != nil {
		for b != nil {
			stack = append(stack, b)
			b = b.Left
		}
		top := stack[len(stack)-1]
		if top.Right == nil || top.Right == lastvisited {
			ans = append(ans, top.Val)
			stack = stack[:len(stack)-1]
			lastvisited = top
		} else {
			b = top.Right
		}
	}
	return ans
}

//CengciTraverse 层次遍历
func CengciTraverse(root *BinaryTree) {
	q := queue.GetQueue()
	q.Push(root)
	count := 1
	for !q.Empty() {
		tmpCount := 0
		for count > 0 {
			now, _ := q.Pop()
			tmp := now.(*BinaryTree)
			fmt.Printf("%d ", tmp.Val)
			count--
			if tmp.Left != nil {
				q.Push(tmp.Left)
				tmpCount++
			}
			if tmp.Right != nil {
				q.Push(tmp.Right)
				tmpCount++
			}
		}
		count = tmpCount
	}
}
