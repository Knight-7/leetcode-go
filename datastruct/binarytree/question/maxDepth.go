//Author  :     knight
//Date    :     2020/06/26 09:32:20
//Version :     1.0
//Email   :     knight2347@163.com
//idea    :     求二叉树的最大深度 分治法

package question

import "leetcode/datastruct/binarytree"

//MaxDepth 二叉树的最大深度
func MaxDepth(root *binarytree.TreeNode) int {
	if root == nil {
		return 0
	}

	leftDepth := MaxDepth(root.Left)
	rightDepth := MaxDepth(root.Right)

	if leftDepth > rightDepth {
		return leftDepth + 1
	}
	return rightDepth + 1
}
