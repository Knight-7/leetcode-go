/*
* @Author: Knight
* @Date:   2020/8/6 11:32
* @Email:  knight2347@163.com
* @Ideal:
 */

package question

import "leetcode/datastruct/binarytree"

func pathSum(root *binarytree.TreeNode, sum int) [][]int {
	curPath := make([]int, 0)
	ret := make([][]int, 0)
	if root == nil {
		return ret
	}

	path(root, &curPath, 0, sum, &ret)
	return ret
}

func path(root *binarytree.TreeNode, curPath *[]int, curSum, sum int, ret *[][]int) {
	if root.Left == nil && root.Right == nil {
		if curSum+root.Val == sum {
			tmp := make([]int, len(*curPath))
			copy(tmp, *curPath)
			tmp = append(tmp, root.Val)
			*ret = append(*ret, tmp)
			return
		} else {
			return
		}
	}

	curSum += root.Val
	*curPath = append(*curPath, root.Val)
	if root.Left != nil {
		path(root.Left, curPath, curSum, sum, ret)
	}

	if root.Right != nil {
		path(root.Right, curPath, curSum, sum, ret)
	}
	*curPath = (*curPath)[:len(*curPath)-1]
}
