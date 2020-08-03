//Author  :     knight
//Date    :     2020/07/23 08:00:07
//Version :     1.0
//Email   :     knight2347@163.com
//idea    :     二叉查找树

package binarytree

import "fmt"

//SearchTree 二叉查找树
type SearchTree struct {
	Root *Node
}

//Node 节点
type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

//Insert 插入函数
func (bst *SearchTree) Insert(val int) {
	//如果根节点为空，那么改点就为根节点，否则插入一个点
	if bst.Root == nil {
		bst.Root = &Node{val, nil, nil}
	} else {
		insertNode(bst.Root, val)
	}
}

func insertNode(root *Node, val int) {
	if val < root.Val {
		if root.Left == nil {
			root.Left = &Node{val, nil, nil}
		} else {
			insertNode(root.Left, val)
		}
	} else if val > root.Val {
		if root.Right == nil {
			root.Right = &Node{val, nil, nil}
		} else {
			insertNode(root.Right, val)
		}
	}

	//当等于的时候，说明val已经在该二叉查找树了，不做任何处理
}

//Delete 删除节点值为val的节点
func (bst *SearchTree) Delete(val int) bool {
	_, ok := deleteNode(bst.Root, val)
	return ok
}

func deleteNode(root *Node, val int) (*Node, bool) {
	if root == nil {
		return root, false
	}

	existed := false

	//从左边查找，同时左子结点指向删除后的左子树
	if val < root.Val {
		root.Left, existed = deleteNode(root.Left, val)
		return root, existed
	}
	//从右边查找，同时右子节点指向删除后的右子树
	if val > root.Val {
		root.Right, existed = deleteNode(root.Right, val)
		return root, existed
	}

	existed = true

	//该节点是叶子节点
	if root.Left == nil && root.Right == nil {
		root = nil
		return root, existed
	}

	//当前节点的左子结点为空，那么抬高右子节点
	if root.Left == nil {
		root = root.Right
		return root, true
	}

	//当前节点的右子节点为空，那么抬高左子节点
	if root.Right == nil {
		root = root.Left
		return root, existed
	}

	//当左右子节点都非空时
	//使用左子树最大的值或者右子树最小的值当作当前节点的值
	leftMaxNum, _ := max(root.Left)
	root.Val = leftMaxNum
	root.Left, _ = deleteNode(root.Left, leftMaxNum)

	return root, existed
}

//Min 查找最小的值
func (bst *SearchTree) Min() (int, bool) {
	return min(bst.Root)
}

func min(root *Node) (int, bool) {
	if root == nil {
		return 0, false
	}

	n := root
	for {
		if n.Left == nil {
			return n.Val, true
		}
		n = n.Left
	}
}

//Max 查找最大的值
func (bst *SearchTree) Max() (int, bool) {
	return max(bst.Root)
}

func max(root *Node) (int, bool) {
	if root == nil {
		return 0, false
	}

	n := root
	for {
		if n.Right == nil {
			return n.Val, true
		}
		n = n.Right
	}
}

//PreOrderTravse 二叉查找树的先序遍历
func (bst *SearchTree) PreOrderTravse() {
	preOrderTravse(bst.Root)
}

func preOrderTravse(root *Node) {
	if root != nil {
		fmt.Printf("%d ", root.Val)
		preOrderTravse(root.Left)
		preOrderTravse(root.Right)
	}
}

//InOrderTravse 二叉查找树的中序遍历
func (bst *SearchTree) InOrderTravse() {
	inOrderTravse(bst.Root)
}

func inOrderTravse(root *Node) {
	if root != nil {
		inOrderTravse(root.Left)
		fmt.Printf("%d ", root.Val)
		inOrderTravse(root.Right)
	}
}

//PostOrderTravse 二叉查找树的后序遍历
func (bst *SearchTree) PostOrderTravse() {
	postOrderTravse(bst.Root)
}

func postOrderTravse(root *Node) {
	if root != nil {
		postOrderTravse(root.Left)
		postOrderTravse(root.Right)
		fmt.Printf("%d ", root.Val)
	}
}

//LevelTravse 二叉查找树的层次遍历
func (bst *SearchTree) LevelTravse() {
	if bst.Root == nil {
		return
	}
	levelTravse(bst.Root)
}

func levelTravse(root *Node) {
	q := make([]*Node, 0)
	q = append(q, root)
	levelCount := 1     //记录一层中节点的个数
	nextLevelCount := 0 //记录下一层中节点的个数

	for len(q) > 0 {
		for levelCount > 0 {
			top := q[0]
			q = q[1:]
			levelCount--
			fmt.Printf("%d ", top.Val)

			//判断左子结点
			if top.Left != nil {
				q = append(q, top.Left)
				nextLevelCount++
			}
			//判断右子节点
			if top.Right != nil {
				q = append(q, top.Right)
				nextLevelCount++
			}
		}
		fmt.Println()
		levelCount = nextLevelCount
		nextLevelCount = 0
	}
}
