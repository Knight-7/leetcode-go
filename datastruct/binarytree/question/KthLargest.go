/*
* @Author: Knight
* @Date:   2020/8/5 11:17
* @Email:  knight2347@163.com
* @Ideal:  根据二叉搜索树的中序遍历是一个递增的序列
 */

package question

import "leetcode/datastruct/binarytree"

func kthLargest(root *binarytree.TreeNode, k int) int {
	midTraseNums := midTrans(root)
	return midTraseNums[len(midTraseNums)-k]
}

func midTrans(root *binarytree.TreeNode) []int {
	ret := make([]int, 0)

	if root == nil {
		return ret
	}

	stack := make([]*binarytree.TreeNode, 0)
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ret = append(ret, top.Val)
		root = top.Right
	}
	return ret
}
