//Author  :     knight
//Date    :     2020/06/25 21:14:54
//Version :     1.0
//Email   :     knight2347@163.com
//idea    :     根据树的遍历构建二叉树

package question

import "leetcode/datastruct/binarytree"

//ConstructWithPreorderAndInorder 根据前序遍历和中序遍历构建二叉树
func ConstructWithPreorderAndInorder(preOrder, inOrder []int) (root *binarytree.TreeNode) {
	//判断遍历数组是否符合规范
	if preOrder == nil || inOrder == nil || len(preOrder) < 0 || len(preOrder) != len(inOrder) {
		return
	}

	root = &binarytree.TreeNode{preOrder[0], nil, nil}

	//递归终止条件，即当只有一个节点的时候
	if len(preOrder) == 1 {
		if preOrder[0] == inOrder[0] {
			return
		}
	}

	index := 0
	for inOrder[index] != preOrder[0] {
		index++
	}
	if index == len(preOrder) {
		return
	}

	//递归构建左子树
	if index > 0 {
		root.Left = ConstructWithPreorderAndInorder(preOrder[1:1+index], inOrder[0:index])
	}

	//递归构建右子树
	if index+1 < len(preOrder) {
		root.Right = ConstructWithPreorderAndInorder(preOrder[index+1:], inOrder[index+1:])
	}

	return
}

//ConstructWithPostorderAndInorder 根据后序遍历和中序遍历构建二叉树
func ConstructWithPostorderAndInorder(postorder []int, inorder []int) (root *binarytree.TreeNode) {
	//判断遍历数组是否符合规范
	if postorder == nil || inorder == nil || len(postorder) < 0 || len(postorder) != len(inorder) {
		return
	}

	root = &binarytree.TreeNode{postorder[len(postorder)-1], nil, nil}

	//递归终止条件，即只有一个节点的时候
	if len(postorder) == 1 {
		if postorder[0] == inorder[0] {
			return
		}
	}

	index := 0
	for inorder[index] != postorder[len(postorder)-1] {
		index++
	}
	if index == len(postorder) {
		return
	}

	//构造左子树
	if index > 0 {
		root.Left = ConstructWithPostorderAndInorder(postorder[:index], inorder[:index])
	}

	//构造右子树
	if index+1 < len(postorder) {
		root.Right = ConstructWithPostorderAndInorder(postorder[index:len(postorder)-1], inorder[index+1:])
	}

	return
}
