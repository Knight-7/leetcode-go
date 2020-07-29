//Author  :     knight
//Date    :     2020/07/05 16:47:55
//Version :     1.0
//Email   :     knight2347@163.com
//idea    :     根据压入序列和弹出序列，判断弹出序列是否是弹出序列

package stack

//IsPopOrder 判断是否是弹出序列
//创建一个辅助栈，如果下一次要弹出的数字正好是栈顶元素，直接弹出，否则将没有压入栈的元素压入栈中，
//直到栈顶元素等于下一次要弹出的数字。如果所有的元素都压入栈中了，还没有找到下一个要弹出的元素，
//说明该序列不是一个弹出序列
func IsPopOrder(pushOrder, popOrder []int) bool {
	if len(pushOrder) == len(popOrder) && len(pushOrder) == 0 {
		return true
	}

	stack := Stack{data: make([]int, 0), length: 0}
	pushIndex, popIndex := 0, 0
	for popIndex < len(popOrder) {
		//如果栈为空，现将一个数字压入栈
		if stack.Empty() && pushIndex < len(pushOrder) {
			stack.Push(pushOrder[pushIndex])
			pushIndex++
		}
		//当数字没有全部压入栈且栈顶数字不等于下一个要弹出的数字，那么将数字压入栈中
		for pushIndex < len(pushOrder) && stack.Top() != popOrder[popIndex] {
			stack.Push(pushOrder[pushIndex])
			pushIndex++
		}
		//当栈非空且下一个弹出数字等于栈顶数字，那么讲栈顶数字弹出，并判断下一个要弹出的数字是否等于栈顶数字
		for popIndex < len(popOrder) && !stack.Empty() && stack.Top() == popOrder[popIndex] {
			stack.Pop()
			popIndex++
		}
		//如果所有的数字都压入栈中且栈顶元素没有找到下一个要弹出的数字，那么说明该序列不是弹出序列
		if pushIndex == len(pushOrder) && popIndex < len(popOrder) {
			return false
		}
	}

	stack1 := Stack{
		data:   make([]int, 0),
		length: 0,
	}
	stack1.Empty()
	return true
}
