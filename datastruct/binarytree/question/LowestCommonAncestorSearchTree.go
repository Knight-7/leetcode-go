package question

import "leetcode/datastruct/binarytree"

func lowestCommonAncestorSearchTree(root, p, q *binarytree.TreeNode) *binarytree.TreeNode {
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}

	if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestorSearchTree(root.Left, p, q)
	} else if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestorSearchTree(root.Right, p, q)
	} else {
		return root
	}
}
