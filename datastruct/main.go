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

func main() {
	rand.Seed(time.Now().Unix())
	maxHead := head.MaxHead{Data: make([]int, 0)}
	for i := 0; i < 7; i++ {
		maxHead.Insert(rand.Intn(100))
	}
	fmt.Println(maxHead.Data)
	for maxHead.Size() > 0 {
		fmt.Printf("%d %v\n", maxHead.Remove(), maxHead.Data)
	}
}
