/*
* @Author: Knight
* @Date:   2020/8/7 8:54
* @Email:  knight2347@163.com
* @Ideal:
 */

package question

import "leetcode/datastruct/binarytree"

func isSameTree(p, q *binarytree.TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if (p == nil && q != nil) || (p != nil && q == nil) {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
