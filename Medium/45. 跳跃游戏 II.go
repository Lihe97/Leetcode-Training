package main

import (
	"fmt"
	"time"
)

func jump(nums []int) int {

	c := make([]int,len(nums))
	step := 0
	p := 0
	for i := 0 ; i < len(nums) ; i++{
		c[i] = i + nums[i]
	}
	fmt.Println(c)

	for p < len(nums) - 1{
		max := 0
		temp := 0

		for i := 1 ; i <= nums[p] ; i ++{
			if p + i < len(nums) && c[p+i] > max {
				max = c[p + i]
				temp = p + i
			}
		}
		fmt.Println("max",max)
		p = temp
		step ++
		fmt.Println(p)
		time.Sleep(time.Second)
	}

	return step
}

func main() {

	a := []int{3,2,1}
	fmt.Println(jump(a))
	
}
