//Author  :     knight
//Date    :     2020/07/27 13:04:05
//Version :     1.0
//Email   :     knight2347@163.com
//idea    :     堆
//堆的定义: 1、堆是一种完全二叉树，
//        : 2、堆要求孩子节点要小于等于或者大于等于父节点

package heap

//Heap 堆
//Data 堆中维护的数据
//Cmp 一个比较函数，cmp为 max(a, b) a > b 时为大根堆， min(a, b) a < b时为小根堆
type Head struct {
	Data []interface{}
	Cmp  func(interface{}, interface{}) bool
}

//上浮
func (h *Head) shiftUp(index int) {
	parent := h.parent(index)
	for index > 0 && h.Cmp(h.Data[index], h.Data[parent]) {
		h.swap(parent, index)
		index = parent
		parent = h.parent(index)
	}
}

//下沉
func (h *Head) shiftDown(index int) {
	length := len(h.Data)
	leftChild := h.leftChild(index)
	for leftChild < length {
		if leftChild+1 < length && h.Cmp(h.Data[leftChild+1], h.Data[leftChild]) {
			leftChild++
		}
		if !h.Cmp(h.Data[index], h.Data[leftChild]) && h.Data[index] != h.Data[leftChild] {
			h.swap(index, leftChild)
			index = leftChild
			leftChild = h.leftChild(index)
		} else {
			break
		}
	}
}

func (h *Head) swap(index1, index2 int) {
	tmp := h.Data[index1]
	h.Data[index1] = h.Data[index2]
	h.Data[index2] = tmp
}

func (h *Head) parent(child int) int {
	return (child - 1) / 2
}

func (h *Head) leftChild(parent int) int {
	return parent*2 + 1
}

func (h *Head) rightChild(parent int) int {
	return parent*2 + 2
}

//Insert 堆中插入数据
// Example:
func (h *Head) Push(val interface{}) {
	h.Data = append(h.Data, val)
	h.shiftUp(len(h.Data) - 1)
}

//Top 获取堆的顶部数据，获取之前，需要判断堆是否为空
func (h *Head) Top() interface{} {
	return h.Data[0]
}

//Size 堆的大小
func (h *Head) Size() int {
	return len(h.Data)
}

//Pop 弹出堆的顶部
func (h *Head) Pop() interface{} {
	top := h.Top()
	h.Data[0] = h.Data[len(h.Data)-1]
	h.Data = h.Data[:len(h.Data)-1]
	h.shiftDown(0)
	return top
}

func (h *Head) HeapSort() {

}
