//Author  :     knight
//Date    :     2020/06/20 14:41:39
//Version :     1.0
//Email   :     knight2347@163.com
//idea    :

package main

import "fmt"

func main() {
	var m map[string]int
	m = make(map[string]int)
	mdMap(m)
	fmt.Printf("%p %v\n", &m, m)
}

func mdMap(m map[string]int) {
	//m = make(map[string]int)
	m["a"] = 1
	m["b"] = 2
	fmt.Printf("%p %v\n", &m, m)
}