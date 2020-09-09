/*
* @Author: Knight
* @Date:   2020/9/4 7:44
* @Email:  knight2347@163.com
* @Ideal:
 */

package question

import (
	"leetcode/datastruct/binarytree"
	"strconv"
)

func binaryTreePaths(root *binarytree.TreeNode) []string {
	ans := make([]string, 0)
	if root == nil {
		return ans
	}
	var getPath func(*binarytree.TreeNode, string)
	getPath = func(root *binarytree.TreeNode, path string) {
		if root.Left == nil && root.Right == nil {
			path += strconv.Itoa(root.Val)
			ans = append(ans, path)
			return
		}
		if root.Left != nil {
			getPath(root.Left, path+strconv.Itoa(root.Val)+"->")
		}
		if root.Right != nil {
			getPath(root.Right, path+strconv.Itoa(root.Val)+"->")
		}
	}
	getPath(root, "")
	return ans
}
