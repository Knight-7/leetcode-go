/*
* @Author: Knight
* @Date:   2020/9/1 8:49
* @Email:  knight2347@163.com
* @Ideal:
 */

package dp

func PredictTheWinner(nums []int) bool {
	return recursion(nums, 0, len(nums)-1, 1) >= 0
}

func recursion(nums []int, start, end, turn int) int {
	if start == end {
		return nums[start] * turn
	}
	startTurn := nums[start]*turn + recursion(nums, start+1, end, -turn)
	endTurn := nums[end]*turn + recursion(nums, start, end-1, -turn)
	return max(startTurn*turn, endTurn*turn) * turn
}
