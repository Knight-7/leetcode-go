//Author  :     knight
//Date    :     2020/06/25 22:35:24
//Version :     1.0
//Email   :     knight2347@163.com
//idea    :     判断一个二叉树是否是高度平衡二叉树，分治法

package main

//TreeNode 二叉树
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	isBalanc, _ := getHeight(root)
	return isBalanc
}

func getHeight(root *TreeNode) (bool, int) {
	if root == nil {
		return true, 0
	}
	l, leftH := getHeight(root.Left)
	r, rightH := getHeight(root.Right)

	if !(l && r) {
		return false, 0
	}
	if leftH-rightH > 1 || rightH-leftH > 1 {
		return false, 0
	}
	if leftH > rightH {
		return true, leftH + 1
	}
	return true, rightH + 1
}
