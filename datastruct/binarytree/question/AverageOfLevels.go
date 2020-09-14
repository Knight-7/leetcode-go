/*
* @Author: Knight
* @Date:   2020/9/12 11:42
* @Email:  knight2347@163.com
* @Ideal:  leetcode637 二叉树的层平均值
 */

package question

import "leetcode/datastruct/binarytree"

func averageOfLevels(root *binarytree.TreeNode) []float64 {
	ans := make([]float64, 0)
	if root == nil {
		return ans
	}
	queue := []*binarytree.TreeNode{root}
	tmp, tmpSum := 0, 0
	for len(queue) > 0 {
		tmp = len(queue)
		t := tmp
		for tmp > 0 {
			top := queue[0]
			queue = queue[1:]
			if top.Left != nil {
				queue = append(queue, top.Left)
			}
			if top.Right != nil {
				queue = append(queue, top.Right)
			}
			tmpSum += top.Val
			tmp--
		}
		avg := float64(tmpSum) / float64(t)
		ans = append(ans, avg)
		tmpSum = 0
	}
	return ans
}
