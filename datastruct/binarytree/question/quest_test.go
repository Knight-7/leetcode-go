/*
* @Author: Knight
* @Date:   2020/8/5 11:34
* @Email:  knight2347@163.com
* @Ideal:
 */

package question

import (
	"leetcode/datastruct/binarytree"
	"testing"
)

func TestKthLargest(t *testing.T) {
	nums := []int{5, 3, 6, 2, 4, 1}
	bst := binarytree.SearchTree{}
	for _, v := range nums {
		bst.Insert(v)
	}
	ret := kthLargest(bst.Root, 3)
	t.Log(ret)
}
