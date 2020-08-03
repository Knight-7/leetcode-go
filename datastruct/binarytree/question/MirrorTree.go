package question

import "leetcode/datastruct/binarytree"

//二叉树的镜像
func mirrorTree(root *binarytree.TreeNode) *binarytree.TreeNode {
	if root != nil {
		tmpNode := root.Right
		root.Right = root.Left
		root.Left = tmpNode
		root.Left = mirrorTree(root.Left)
		root.Right = mirrorTree(root.Right)
		return root
	}
	return nil
}
