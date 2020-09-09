package heap

import "container/heap"

type Node struct {
	val int
	cnt int
}

type HeapNode []Node

func (h HeapNode) Len() int { return len(h) }

func (h HeapNode) Less(i, j int) bool { return h[i].cnt < h[j].cnt }

func (h HeapNode) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *HeapNode) Push(x interface{}) {
	*h = append(*h, x.(Node))
}

//配合heap.Pop()使用，将堆顶元素置换到最后，并将最后的元素去掉
func (h *HeapNode) Pop() interface{} {
	tmp := *h
	top := tmp[0]
	*h = tmp[:len(tmp)-1]
	return top
}

func (h HeapNode) Size() int {
	return len(h)
}

//topK问题
//使用官方的最小堆
func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	ans := make([]int, 0)
	for _, v := range nums {
		m[v]++
	}
	h := &HeapNode{}
	for key, val := range m {
		if h.Size() < k {
			heap.Push(h, Node{
				val: key,
				cnt: val,
			})
		} else {
			if val > (*h)[0].cnt {
				heap.Pop(h)
				heap.Push(h, Node{
					val: key,
					cnt: val,
				})
			}
		}
	}
	for i := 0; i < k; i++ {
		ans = append(ans, (*h)[i].val)
	}
	return ans
}

//使用自己写的堆实现最小堆
func topKFrequent1(nums []int, k int) []int {
	m := make(map[int]int)
	ans := make([]int, 0)
	for _, v := range nums {
		m[v]++
	}
	h := Head{
		Data: make([]interface{}, 0),
		//小根堆
		Cmp: func(a, b interface{}) bool {
			if a.(Node).cnt < b.(Node).cnt {
				return true
			}
			return false
		},
	}
	for key, val := range m {
		if h.Size() < k {
			h.Push(Node{
				val: key,
				cnt: val,
			})
		} else {
			if val > h.Top().(Node).cnt {
				h.Pop()
				h.Push(Node{
					val: key,
					cnt: val,
				})
			}
		}
	}
	for i := 0; i < k; i++ {
		ans = append(ans, h.Data[i].(Node).val)
	}
	return ans
}
