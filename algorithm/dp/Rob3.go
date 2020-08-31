/*
* @Author: Knight
* @Date:   2020/8/5 9:48
* @Email:  knight2347@163.com
* @Ideal:  打家劫舍三。房屋的排列成了一个二叉树,所以使用dp+二叉树.
*          这道题实际上还是dp问题，状态转移还是两个选择：
*          1. 选择该节点， 返回的是该节点孙子节点和+该节点值
*          2. 不选择该节点，那么返回的是该节点左右子节点值的和
*          然后取最大值
*          这道题的本质还是dp[i] = max(dp[i-2]+nums[i], dp[i-1]，只不过二叉树中的i-2表现为孙子节点的和，i-1表现为子节点的和
 */

package dp

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob3(root *TreeNode) int {
	ans, _ := robTree(root)
	return ans
}

func robTree(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}

	leftChildVal, leftGrandChildVal := robTree(root.Left)
	rightChildVal, rightGrandChildVal := robTree(root.Right)

	return max(root.Val+leftGrandChildVal+rightGrandChildVal, leftChildVal+rightChildVal), leftChildVal + rightChildVal
}
