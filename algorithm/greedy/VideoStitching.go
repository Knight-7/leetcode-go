/*
* @Author  :     knight
* @Date    :     2020/10/24 12:26:05
* @Version :     1.0
* @Email   :     knight2347@163.com
* @idea    :     贪心算法
 */

package greedy

import (
	"fmt"
	"sort"
)

func videoStitching(clips [][]int, T int) int {
	ans := -1
	if len(clips) == 0 && T > 0 {
		return ans
	}
	if len(clips) == 0 && T == 0 {
		return 0
	}
	sort.Slice(clips, func(i, j int) bool {
		if clips[i][0] == clips[j][0] {
			return clips[i][1] > clips[j][1]
		}
		return clips[i][0] < clips[j][0]
	})
	maxEnd := make([]int, T+1)
	for i := 0; i <= T; i++ {
		maxEnd[i] = i
	}
	fmt.Println(clips)
	start, end := clips[0][0], 0
	if start != 0 {
		return -1
	}
	i := 0
	for {
		if start >= clips[i][0] {
			end = max(end, clips[i][1])
		} else {
			if end >= clips[i][0] {
				start = clips[i][0]
				end = max(end, clips[i][1])
				ans++
			} else {
				return -1
			}
		}
		if end == T {
			break
		}
		i++
		if i == len(clips) {
			return -1
		}
	}
	return ans
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
