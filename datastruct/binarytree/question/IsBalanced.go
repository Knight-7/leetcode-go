/*
* @Author: Knight
* @Date:   2020/8/5 11:47
* @Email:  knight2347@163.com
* @Ideal:
 */

package question

import "leetcode/datastruct/binarytree"

func isBalanced(root *binarytree.TreeNode) bool {
	_, ret := balance(root)
	return ret
}

func balance(root *binarytree.TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}

	leftDepth, leftBalance := balance(root.Left)
	rightDepth, rightBalance := balance(root.Right)

	if !leftBalance || !rightBalance {
		return 0, false
	}

	if abs(leftDepth-rightDepth) > 1 {
		return 0, false
	}

	return max(leftDepth, rightDepth) + 1, true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
