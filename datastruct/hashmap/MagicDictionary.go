/*
* @Author  :     knight
* @Date    :     2020/09/13 18:13:40
* @Email   :     knight2347@163.com
* @idea    :     leetcode 676 实现一个魔法字典
*/

package hashmap

type MagicDictionary struct {
	keys    map[int]bool
	buckets map[int][]string
}


/** Initialize your data structure here. */
func Constructor() MagicDictionary {
	return MagicDictionary{
		keys:    make(map[int]bool),
		buckets: make(map[int][]string),
	}
}


func (this *MagicDictionary) BuildDict(dictionary []string)  {
	for _, v := range dictionary {
		this.keys[len(v)] = true
		this.buckets[len(v)] = append(this.buckets[len(v)], v)
	}
}


func (this *MagicDictionary) Search(searchWord string) bool {
	length := len(searchWord)
	if _, ok := this.keys[length]; !ok {
		return false
	}
	for _, v := range this.buckets[length] {
		differences := 0
		for i := 0; i < length; i++ {
			if v[i] != searchWord[i] {
				differences++
			}
			if differences > 1 {
				break
			}
		}
		if differences == 1 {
			return true
		}
	}
	return false
}


/**
 * Your MagicDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.BuildDict(dictionary);
 * param_2 := obj.Search(searchWord);
 */
