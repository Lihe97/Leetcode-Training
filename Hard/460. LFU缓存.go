package main

import (
	"fmt"
	"math"
)

type LFUCache struct {
	mp map[int][]int
	length int

}
func Constructor(capacity int) LFUCache {
	mp := make(map[int][]int,capacity)
	return LFUCache{mp:mp,length:capacity}
}
func (this *LFUCache) Get(key int) int {
	for a,b := range this.mp{
		if a == key{

			b = append(b, 1)


			this.mp[a] = b

			return b[0]
		}
	}

	return -1
}
func (this *LFUCache) Put(key int, value int)  {
	if this.length == 0{
		return
	}
	if len(this.mp) < this.length{
		temp := []int{value}
		this.mp[key] = temp

	}else{
	min := math.MaxInt8
	var del int
	for a,b := range this.mp{
		if len(b) < min {
			min = len(b)
			del =  a
		}
	}

	delete(this.mp, del)
	temp := []int{value}
	this.mp[key] = temp
	}
}

func main() {

	obj := Constructor(2)
	obj.Put(1,1)
	obj.Put(2,2)
	fmt.Println(obj.Get(1))
	obj.Put(3,3)
	//fmt.Println(obj.mp)
	fmt.Println(obj.Get(2))
	fmt.Println(obj.Get(3))
	obj.Put(4,4)
	fmt.Println(obj.Get(1))
	fmt.Println(obj.Get(3))
	fmt.Println(obj.Get(4))

}
