/*
* @Author: Knight
* @Date:   2020/8/5 9:32
* @Email:  knight2347@163.com
* @Ideal:  打家劫舍二。所有的房子都围成了一圈，意味着第一个和最后一个是无法同时选择的。那问题就转换到从去掉第一个房屋或者去掉第最后一个房屋中
*          进行打家劫舍一操作，并取最大值
 */

package dp

func rob2(nums []int) int {
	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return nums[0]
	}
	return max(rob2_1(nums[1:]), rob2_1(nums[:len(nums)-1]))
}

func rob2_1(nums []int) int {
	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return nums[0]
	}

	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-2]+nums[i], dp[i-1])
	}
	return dp[len(nums)-1]
}
