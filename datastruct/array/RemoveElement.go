/*
* @Author: Knight
* @Date:   2020/8/12 14:13
* @Email:  knight2347@163.com
* @Ideal:
 */

package array

func removeElement(nums []int, val int) int {
	lastIndex := len(nums) - 1
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] == val {
			nums[i], nums[lastIndex] = nums[lastIndex], nums[i]
			lastIndex--
		}
	}
	return lastIndex + 1
}
