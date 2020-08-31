/*
* @Author: Knight
* @Date:   2020/8/6 10:54
* @Email:  knight2347@163.com
* @Ideal:  判断一个序列是否是二叉搜索树的后序遍历.
*          根据二叉搜索树后序遍历的特点，序列的最后一个值为根节点，
*          同时序列中从第一个大于最后一个节点开始，是属于根节点的的右子树并且其中所有的值都是大于根节点的，
*          而那个节点之前的都是根节点的左子树
 */

package question

func verifyPostorder(postorder []int) bool {
	return verify(postorder, 0, len(postorder)-1)
}

func verify(postorder []int, left, right int) bool {
	if left >= right {
		return true
	}

	p := left
	for postorder[p] < postorder[right] {
		p++
	}
	mid := p
	for postorder[p] > postorder[right] {
		p++
	}
	if p != right {
		return false
	}
	return verify(postorder, left, mid-1) && verify(postorder, mid, right-1)
}
