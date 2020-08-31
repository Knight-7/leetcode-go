/*
* @Author: Knight
* @Date:   2020/8/12 8:02
* @Email:  knight2347@163.com
* @Ideal:  克隆图，给定一个图，然后clone它，clone意味着，图中的节点值是相同的，但是节点却是新的节点
 */

package BFS

type Node struct {
	Val       int
	Neighbors []*Node
}

//bfs版
func cloneGraph(node *Node) *Node {
	if node == nil {
		return node
	}
	cloneNode := &Node{
		Val:       node.Val,
		Neighbors: make([]*Node, 0),
	}
	vis := make(map[*Node]*Node)
	vis[node] = cloneNode
	queue := make([]*Node, 0)
	queue = append(queue, node)
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		for _, n := range cur.Neighbors {
			if _, ok := vis[n]; !ok {
				vis[n] = &Node{Val: n.Val, Neighbors: make([]*Node, 0)}
				queue = append(queue, n)
			}
			vis[cur].Neighbors = append(vis[cur].Neighbors, vis[n])
		}
	}
	return cloneNode
}

//dfs版
func cloneGraph2(node *Node) *Node {
	vis := make(map[*Node]*Node)
	var clone func(node *Node) *Node

	clone = func(node *Node) *Node {
		if node == nil {
			return node
		}

		if _, ok := vis[node]; ok {
			return vis[node]
		}

		cloneNode := &Node{node.Val, []*Node{}}
		vis[node] = cloneNode

		for _, n := range node.Neighbors {
			cloneNode.Neighbors = append(cloneNode.Neighbors, clone(n))
		}

		return cloneNode
	}
	return clone(node)
}
