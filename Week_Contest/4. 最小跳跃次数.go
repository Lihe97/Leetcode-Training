package main

import (
	"fmt"
)

func minJump(jump []int) int {

	c := make([]int,len(jump))
	for i := 0 ; i < len(jump) ; i ++{
		c[i] = i + jump[i]
	}
	d := make([]int,len(jump))
	d[0] = c[0]
	for i := 1 ; i < len(jump) ; i ++{
		if c[i] > d[i-1]{
			d[i] = c[i]
		}else{
			d[i] = d[i-1]
		}
	}
	mp := map[int]int{}
	for i:= 0 ; i < len(d) ; i ++{
		mp[i] = d[i]
	}
	fmt.Println(d)
	cur := c[0]
	sum := 1
	//var max int
	fmt.Println(cur)
	for c[cur] < len(jump){

		fmt.Println("cur",cur)
		cur = d[cur]
		sum ++
		fmt.Println(cur)

	}
	return sum+1

}

func main() {

	//a := []int{1,2}
	a := []int{2,5,1,1,1,1}


	fmt.Println(minJump(a))
}
