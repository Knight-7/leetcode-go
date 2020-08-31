/*
* @Author: Knight
* @Date:   2020/8/12 21:16
* @Email:  knight2347@163.com
* @Ideal:
 */

package array

func runningSum(nums []int) []int {
	var tmp int
	ans := []int{}
	for _, v := range nums {
		tmp += v
		ans = append(ans, tmp)
	}
	return ans
}
