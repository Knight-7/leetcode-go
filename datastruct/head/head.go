//Author  :     knight
//Date    :     2020/07/27 13:04:05
//Version :     1.0
//Email   :     knight2347@163.com
//idea    :     堆
//堆的定义: 1、堆是一种完全二叉树，
//        : 2、堆要求孩子节点要小于等于或者大于等于父节点

package head

//Header 堆
type Header interface {
	Insert(val int)
	Top() int
	Remove() int
	Size() int
}

//MaxHead 最大堆
type MaxHead struct {
	Data []int
}

//MinHead 最小堆
type MinHead struct {
	Data []int
}

func up(parentIndex, childIndex *int) {
	*childIndex = *parentIndex
	*parentIndex = (*parentIndex - 1) / 2
}

func down(parentIndex, nextChild1, nextChild2 *int) {
	*parentIndex = *nextChild1
	*nextChild1 = *parentIndex * 2 + 1
	*nextChild2 = *parentIndex * 2 + 2
}

//Insert 最大堆中插入数据（上浮）
func (h *MaxHead) Insert(val int) {
	h.Data = append(h.Data, val)
	childIndex := len(h.Data) - 1
	parentIndex := (childIndex - 1) / 2
	//上浮，当子节点的值大于父节点的值时，交换节点的值；
	//重复，直到子节点的值小于父节点的值
	for childIndex > 0 && h.Data[childIndex] > h.Data[parentIndex] {
		swap(&h.Data[childIndex], &h.Data[parentIndex])
		up(&parentIndex, &childIndex)
	}
}

//Top 获取最大堆的顶部最大值
func (h *MaxHead) Top() int {
	return 111
}

//Remove 移除并获取最大堆的顶部（下沉）
func (h *MaxHead) Remove() int {
	//获取最大节点，并将最后一个节点的值放到根节点，然后去掉最后一个节点
	top := h.Data[0]
	h.Data[0] = h.Data[len(h.Data)-1]
	h.Data = h.Data[:len(h.Data)-1]

	length := len(h.Data)
	curentIndex := 0
	leftIndex, rightIndex := curentIndex * 2 + 1, curentIndex * 2 + 2
	for leftIndex < length {
		if rightIndex < length {
			if h.Data[curentIndex] > max(h.Data[leftIndex], h.Data[rightIndex]) {
				break
			}
			if h.Data[curentIndex] == h.Data[leftIndex] || h.Data[curentIndex] == h.Data[rightIndex] {
				break
			}
			if h.Data[leftIndex] > h.Data[rightIndex] {
				if h.Data[curentIndex] < h.Data[leftIndex] {
					swap(&h.Data[curentIndex], &h.Data[leftIndex])
					down(&curentIndex, &leftIndex, &rightIndex)
				} else if h.Data[curentIndex] < h.Data[rightIndex] {
					swap(&h.Data[curentIndex], &h.Data[rightIndex])
					down(&curentIndex, &rightIndex, &leftIndex)
				}
			} else {
				if h.Data[rightIndex] > h.Data[leftIndex] {
					if h.Data[curentIndex] < h.Data[rightIndex] {
						swap(&h.Data[curentIndex], &h.Data[rightIndex])
					    down(&curentIndex, &rightIndex, &leftIndex)
					} else if h.Data[leftIndex] < h.Data[rightIndex] {
						swap(&h.Data[curentIndex], &h.Data[leftIndex])
						down(&curentIndex, &leftIndex, &rightIndex)
					}
				}
			}
		} else {
			if h.Data[curentIndex] < h.Data[leftIndex] {
				swap(&h.Data[curentIndex], &h.Data[leftIndex])
			}
			break
		}
	}

	return top
}

//Size 最大堆的大小
func (h *MaxHead) Size() int {
	return len(h.Data)
}

//Insert 最小堆中插入数据
func (h *MinHead) Insert(val int) {
	h.Data = append(h.Data, val)
	lastIndex := len(h.Data) - 1
	for lastIndex > 0 && h.Data[lastIndex] > h.Data[lastIndex/2] {
		tmp := h.Data[lastIndex]
		h.Data[lastIndex] = h.Data[lastIndex/2]
		h.Data[lastIndex/2] = tmp
		lastIndex = lastIndex / 2
	}
}

//Top 获取最小堆的顶部最大值
func (h *MinHead) Top() int {
	return 222
}

//Remove 移除并获取最小堆的顶部
func (h *MinHead) Remove() int {
	return 0
}

//Size 最小堆的大小
func (h *MinHead) Size() int {
	return len(h.Data)
}

func swap(a, b *int) {
	tmp := *a
	*a = *b
	*b = tmp
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
