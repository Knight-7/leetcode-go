/*
* @Author: Knight
* @Date:   2020/8/21 10:06
* @Email:  knight2347@163.com
* @Ideal:  二叉树的最小深度
 */

package question

import "leetcode/datastruct/binarytree"

func minDepth(root *binarytree.TreeNode) int {
	depth := 0
	if root == nil {
		return depth
	}

	queue := make([]*binarytree.TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		depth++
		length := len(queue)
		for length > 0 {
			top := queue[0]
			queue = queue[1:]
			if top.Left == nil && top.Right == nil {
				return depth
			}
			if top.Left != nil {
				queue = append(queue, top.Left)
			}
			if top.Right != nil {
				queue = append(queue, top.Right)
			}
			length--
		}
	}
	return depth
}
