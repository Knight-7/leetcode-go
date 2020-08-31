/*
* @Author: Knight
* @Date:   2020/8/27 19:14
* @Email:  knight2347@163.com
* @Ideal:  leetcode332 重新安排行程
 */

package string

import "sort"

func findItinerary(tickets [][]string) []string {
	sort.Slice(tickets, func(i, j int) bool {
		if tickets[i][0] == tickets[j][0] {
			return tickets[i][1] < tickets[j][1]
		}
		return tickets[i][0] < tickets[j][0]
	})

	return ans
}
