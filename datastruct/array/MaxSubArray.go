/*
* @Author: Knight
* @Date:   2020/8/17 20:25
* @Email:  knight2347@163.com
* @Ideal:  连续子数组的最大和
*  if dp(i - 1) < 0, 那么dp[i] = nums[i]， nums[i]+dp[i-1]必然会比nums[i]小，
*  else dp[i-1] >= 0, dp[i] = dp[i-1]+nums[i],此时dp[i]肯定比nums[i]大，所以可能产生最大值
 */

package array

func maxSubArray(nums []int) int {
	tmp := 0
	maxAns := -1 << 63
	for i := 0; i < len(nums); i++ {
		if tmp < 0 {
			tmp = nums[i]
		} else {
			tmp += nums[i]
		}
		if tmp > maxAns {
			maxAns = tmp
		}
	}
	return maxAns
}
