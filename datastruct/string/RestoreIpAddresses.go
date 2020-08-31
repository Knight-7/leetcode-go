/*
* @Author: Knight
* @Date:   2020/8/9 9:27
* @Email:  knight2347@163.com
* @Ideal:  复原IP地址，IP地址每个整数只能以1或2开头，第三位最大值是5
*          回溯法
 */

package string

import "strconv"

var (
	ans    []string
	tmpNum []int
)

func restoreIpAddresses(s string) []string {
	ans = make([]string, 0)
	tmpNum = make([]int, 4)
	dfs(s, 0, 0)
	return ans
}

func dfs(s string, strIndex, tmpIndex int) {
	if tmpIndex == 4 {
		if strIndex == len(s) {
			addr := ""
			for k, v := range tmpNum {
				addr += strconv.Itoa(v)
				if k != len(tmpNum)-1 {
					addr += "."
				}
			}
			ans = append(ans, addr)
		}
		return
	}
	if strIndex == len(s) {
		return
	}

	if s[strIndex] == '0' {
		tmpNum[tmpIndex] = 0
		dfs(s, strIndex+1, tmpIndex+1)
	}

	addr := 0
	for i := strIndex; i < len(s); i++ {
		addr = addr*10 + int(s[i]-'0')
		if addr > 0 && addr <= 255 {
			tmpNum[tmpIndex] = addr
			dfs(s, i+1, tmpIndex+1)
		} else {
			break
		}
	}
}
