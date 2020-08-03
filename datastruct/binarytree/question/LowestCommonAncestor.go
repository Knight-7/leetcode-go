package question

import "leetcode/datastruct/binarytree"

//二叉树的最近公共祖先
func lowestCommonAncestor(root, p, q *binarytree.TreeNode) *binarytree.TreeNode {
	if root == nil || root == p || root == q {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left == nil {
		return right
	} else if right == nil {
		return left
	}
	return root
}
