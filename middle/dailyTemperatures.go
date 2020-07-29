//Author  :     knight
//Date    :     2020/06/11 08:39:54
//Version :     1.0
//Email   :     knight2347@163.com
//idea    :     方法一：暴力求解；方法二，单调栈

package main

import "fmt"

//暴力求解法
func dialyTemperatures1(T []int) []int {
	ans := make([]int, 0)

	current := 0
	for i := 0; i < len(T); i++ {
		current = T[i]
		for j := i + 1; j < len(T); j++ {
			if T[j] > current {
				ans = append(ans, j-i)
				break
			}
			if j == len(T)-1 {
				ans = append(ans, 0)
			}
		}
		if i == len(T)-1 {
			ans = append(ans, 0)
		}
	}
	return ans
}

type stack struct {
	nums   []int
	length int
}

func (s *stack) getLength() int {
	return s.length
}

func (s *stack) push(x int) {
	s.nums = append(s.nums, x)
	s.length++
}

func (s *stack) pop() int {
	topNum := s.nums[s.length-1]
	s.nums = s.nums[:s.length-1]
	s.length = len(s.nums)
	return topNum
}

func (s *stack) top() int {
	return s.nums[s.length-1]
}

func (s *stack) empty() bool {
	return s.length == 0
}

//单调栈
func dialyTemperatures2(T []int) []int {
	length := len(T)
	ans := make([]int, length)
	s := stack{nums: make([]int, 0), length: 0}

	for i := 0; i < length; i++ {
		if s.empty() || T[i] <= T[s.top()] {
			s.push(i)
			continue
		}
		for !s.empty() && T[i] > T[s.top()] {
			top := s.pop()
			ans[top] = i - top
		}
		s.push(i)
	}
	return ans
}

func main0() {
	t := []int{73, 74, 75, 71, 69, 72, 76, 73}
	fmt.Println(dialyTemperatures2(t))
}
