/*
* @Author: Knight
* @Date:   2020/9/6 9:11
* @Email:  knight2347@163.com
* @Ideal:
 */
package question

import "leetcode/datastruct/binarytree"

func levelOrderBottom(root *binarytree.TreeNode) [][]int {
	ans := make([][]int, 0)
	if root == nil {
		return ans
	}
	queue := make([]*binarytree.TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		length := len(queue)
		tmp := make([]int, 0)
		for _, v := range queue {
			tmp = append(tmp, v.Val)
		}
		cur := append([][]int{}, tmp)
		ans = append(cur, ans...)
		for length > 0 {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			length--
		}
	}
	return ans
}
