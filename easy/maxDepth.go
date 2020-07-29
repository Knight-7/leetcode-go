//Author  :     knight
//Date    :     2020/06/25 22:28:49
//Version :     1.0
//Email   :     knight2347@163.com
//idea    :     查找二叉树的最大深度

package main

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)

	if left > right {
		return left + 1
	}

	return right + 1
}
