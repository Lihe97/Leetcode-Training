package main

import (
	"fmt"
	"math"
)

type LRUCache struct {
	mp map[int][2]int
	length int
}


func Constructor(capacity int) LRUCache {
	//return LRUCache{mp: map[int][2]int{},capacity}
	return LRUCache{
		mp: map[int][2]int{},
		length: capacity,
	}
}
func (this *LRUCache) Get(key int) int {

	maxt := 0
	for _,a := range this.mp{
		if a[1] > maxt{
			maxt = a[1]
		}
	}
	for a,x := range this.mp{
		if a == key{
			temp := [2]int{x[0],maxt + 1}
			this.mp[a] = temp
			return x[0]
		}
	}
	return -1
}
func (this *LRUCache) Put(key int, value int)  {
	maxt := 0
	mint := math.MaxInt8
	var min int
	for b,a := range this.mp{
		if a[1] < mint{
			mint = a[1]
			min = b
		}
		if a[1] > maxt{
			maxt = a[1]
		}
	}
	if _,ok := this.mp[key];ok{
		temp := [2]int{value,maxt}
		this.mp[key] = temp
	}else{
	if len(this.mp) == 0{
		temp := [2]int{value,1}
		this.mp[key] = temp
	}else if len(this.mp) < this.length{
		temp := [2]int{value,maxt + 1}
		this.mp[key] = temp
	}else{
		delete(this.mp,min)
		temp := [2]int{value,maxt + 1}

		this.mp[key] = temp
	}
	}

}

func main() {

	obj := Constructor(2)
	fmt.Println("拿2",obj.Get(2))
	obj.Put(2,6)
	fmt.Println("拿1",obj.Get(1))
	obj.Put(1,5)
	obj.Put(1,2)
	fmt.Println("拿1",obj.Get(1))
	fmt.Println("拿2",obj.Get(2))

}
