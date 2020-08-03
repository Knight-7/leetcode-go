//Author  :     knight
//Date    :     2020/06/26 09:08:59
//Version :     1.0
//Email   :     knight2347@163.com
//idea    :

package main

import (
	"fmt"
	"leetcode/datastruct/head"
	"math/rand"
	"time"
)

func min(a, b int) bool {
	if a < b {
		return true
	}
	return false
}

func max(a, b int) bool {
	if a > b {
		return true
	}
	return false
}

func main() {
	rand.Seed(time.Now().Unix())
	maxHead := head.Head{
		Data: make([]int, 0),
		Cmp:  max,
	}
	minHead := head.Head{
		Data: make([]int, 0),
		Cmp:  min,
	}
	for i := 0; i < 7; i++ {
		a := rand.Intn(100)
		maxHead.Insert(a)
		minHead.Insert(a)
	}
	fmt.Println(maxHead.Data)
	fmt.Println(minHead.Data)
	for maxHead.Size() > 0 {
		fmt.Printf("%d ", maxHead.Pop())
	}
	fmt.Println()
	for minHead.Size() > 0 {
		fmt.Printf("%d ", minHead.Pop())
	}
	fmt.Println()
}
