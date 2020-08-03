package question

import "leetcode/datastruct/binarytree"

//将二叉树展开成一个链表
//二叉树的前序遍历，并将二叉树的右子节点设为空，右子节点指向下一个节点
func Flatten(root *binarytree.TreeNode) {
	nodes := preOrder(root)
	for i := 0; i < len(nodes)-1; i++ {
		nodes[i].Left = nil
		nodes[i].Right = nodes[i+1]
	}
}

func preOrder(root *binarytree.TreeNode) []*binarytree.TreeNode {
	ans := make([]*binarytree.TreeNode, 0)
	stack := make([]*binarytree.TreeNode, 0)
	for len(stack) > 0 || root != nil {
		for root != nil {
			ans = append(ans, root)
			stack = append(stack, root)
			root = root.Left
		}
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		root = top.Right
	}
	return ans
}

func inOrder(root *binarytree.TreeNode) []*binarytree.TreeNode {
	ans := make([]*binarytree.TreeNode, 0)
	stack := make([]*binarytree.TreeNode, 0)
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		top := stack[len(stack)-1]
		ans = append(ans, stack[len(stack)-1])
		stack = stack[:len(stack)-1]
		root = top.Right
	}
	return ans
}

func postOrder(root *binarytree.TreeNode) []*binarytree.TreeNode {
	ans := make([]*binarytree.TreeNode, 0)
	stack := make([]*binarytree.TreeNode, 0)
	var lastNode *binarytree.TreeNode
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		top := stack[len(stack)-1]
		if top.Right == nil || top.Right == lastNode {
			ans = append(ans, top)
			stack = stack[:len(stack)-1]
			lastNode = top

		} else {
			root = top.Right
		}
	}
	return ans
}
