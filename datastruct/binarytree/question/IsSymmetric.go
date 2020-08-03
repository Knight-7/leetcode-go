package question

import "leetcode/datastruct/binarytree"

//对称的二叉树
func isSymmetric(root *binarytree.TreeNode) bool {
	ans := true
	if root != nil {
		ans = symmetric(root.Left, root.Right)
	}
	return ans
}

func symmetric(leftNode, rightNode *binarytree.TreeNode) bool {
	if leftNode == nil && rightNode == nil {
		return true
	}

	//当其中的一个节点为nil时，返回false
	if (leftNode == nil && rightNode != nil) || (leftNode != nil && rightNode == nil) {
		return false
	}

	//当节点的值不相等时，返回false
	if leftNode.Val != rightNode.Val {
		return false
	}

	//对称的节点之间进行比较，即左节点的左节点和右节点的右节点是对称的节点，左节点的右节点和右节点的左节点是对称的节点
	return symmetric(leftNode.Left, rightNode.Right) &&
		symmetric(leftNode.Right, rightNode.Left)
}
