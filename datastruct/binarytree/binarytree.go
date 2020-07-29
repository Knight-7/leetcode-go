//Author  :     knight
//Date    :     2020/06/25 21:11:04
//Version :     1.0
//Email   :     knight2347@163.com
//idea    :     二叉树

package binarytree

//BinaryTree 二叉树
type BinaryTree struct {
	Val   int
	Left  *BinaryTree
	Right *BinaryTree
}