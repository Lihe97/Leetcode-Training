package main

import (
	"fmt"
	"math"
)
type LFUCache struct {
	mp map[int][3]int
	length int
}
func Constructor(capacity int) LFUCache {
	return LFUCache{
		mp: map[int][3]int{},
		length: capacity,
	}
}
func (this *LFUCache) Get(key int) int {
	if this.length == 0{
		return -1
	}
	//for a,b := range this.mp{
	//	fmt.Println(a)
	//	fmt.Println(b)
	//}

	maxt := 0

	for _,b := range this.mp{
		if b[2] > maxt{
			maxt = b[2]
		}
	}
	for a,b := range this.mp{
		if a == key{
			temp := [3]int{b[0],b[1]+1,maxt+1}
			this.mp[a] = temp
			return this.mp[key][0]
		}
	}
	return -1
}


func (this *LFUCache) Put(key int, value int)  {
	maxt := 0
	minuse := math.MaxInt8
	q := []int{}
	for _,b := range this.mp{
		if b[2] > maxt{
			maxt = b[2]
		}
		if b[1] < minuse{
			minuse = b[1]
		}
	}
	for a,b := range this.mp{
		if b[1] == minuse{
			q = append(q, a)
		}
	}
	flag := true
	for a,b  := range this.mp{
		if a == key{
			temp := [3]int{value,b[1]+1,maxt+1}
			this.mp[a] = temp
			flag = false
		}
	}
	if flag{
	if len(this.mp) == 0{
		this.mp[key] = [3]int{value,0,1}
	}else if len(this.mp) < this.length{
		this.mp[key] = [3]int{value,0,maxt+1}
	}else{


		if len(q) == 1{
			delete(this.mp,q[0])
		}else{
			min := math.MaxInt8
			del := 0
			for _,b := range q{
				if this.mp[b][2] < min{
					min = this.mp[b][2]
					del = b
				}
			}
			delete(this.mp,del)


		}
		this.mp[key] = [3]int{value,0,maxt+1}
	}
	}


}


/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
func main() {

	obj := Constructor(2)
	obj.Put(1,1)
	obj.Put(2,2)
	fmt.Println("拿1",obj.Get(1))
	obj.Put(3,3)

	fmt.Println("拿2",obj.Get(2))
	fmt.Println("拿3",obj.Get(3))
	obj.Put(4,4)
	fmt.Println("拿1",obj.Get(1))
	fmt.Println("拿3",obj.Get(3))
	fmt.Println("拿4",obj.Get(4))

}
