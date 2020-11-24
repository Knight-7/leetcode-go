package question

import "leetcode/datastruct/binarytree"

func countNodes(root *binarytree.TreeNode) int {
	ans := 0
	if root == nil {
		return ans
	}
	if root.Left != nil {
		ans += countNodes(root.Left)
	}
	if root.Right != nil {
		ans += countNodes(root.Right)
	}
	return ans + 1
}
