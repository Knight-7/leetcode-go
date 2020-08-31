/*
* @Author: Knight
* @Date:   2020/8/8 10:57
* @Email:  knight2347@163.com
* @Ideal:  恢复二叉搜索树
 */

package question

import "leetcode/datastruct/binarytree"

func recoverTree(root *binarytree.TreeNode) {
	if root == nil {
		return
	}

	//获取中序遍历
	nums := make([]int, 0)
	var inorder func(root *binarytree.TreeNode)
	inorder = func(root *binarytree.TreeNode) {
		if root != nil {
			inorder(root.Left)
			nums = append(nums, root.Val)
			inorder(root.Right)
		}
	}
	inorder(root)
	//找到两个节点被交换的值
	x, y := findSwapNum(nums)
	//交换节点的值
	recover(root, x, y, 2)
}

func findSwapNum(nums []int) (x, y int) {
	x, y = -1, -1
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			y = nums[i+1]
			if x == -1 {
				x = nums[i]
			} else {
				break
			}
		}
	}
	return
}

func recover(root *binarytree.TreeNode, x, y, count int) {
	if root == nil {
		return
	}

	if root.Val == x || root.Val == y {
		if root.Val == x {
			root.Val = y
		} else {
			root.Val = x
		}
		count--
		if count == 0 {
			return
		}
	}

	recover(root.Left, x, y, count)
	recover(root.Right, x, y, count)
}

//在中序遍历的时候，比较当前节点和前一个节点的大小关系
func recoverTree2(root *binarytree.TreeNode) {
	if root == nil {
		return
	}

	var x, y, pre *binarytree.TreeNode
	stack := []*binarytree.TreeNode{}
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if pre != nil && top.Val < pre.Val {
			y = top
			if x == nil {
				x = pre
			} else {
				break
			}
		}
		pre = top
		root = top.Right
	}
	x.Val, y.Val = y.Val, x.Val
}
