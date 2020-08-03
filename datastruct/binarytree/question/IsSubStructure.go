package question

import "leetcode/datastruct/binarytree"

//判断树A是否是树B的子树
func isSubStructure(A *binarytree.TreeNode, B *binarytree.TreeNode) bool {
	result := false

	if A != nil && B != nil {
		//当两个节点的值相同，就判断他们的子节点是否相同
		if A.Val == B.Val {
			result = hasSameStructure(A, B)
		}

		//当没找到子树时，遍历两颗树
		if !result {
			result = isSubStructure(A.Left, B)
		}
		if !result {
			result = isSubStructure(A.Right, B)
		}
	}

	return result
}

func hasSameStructure(root1 *binarytree.TreeNode, root2 *binarytree.TreeNode) bool {
	if root2 == nil {
		return true
	}

	if root1 == nil {
		return false
	}

	if root1.Val != root2.Val {
		return false
	}

	return hasSameStructure(root1.Left, root2.Left) &&
		hasSameStructure(root1.Right, root2.Right)
}
