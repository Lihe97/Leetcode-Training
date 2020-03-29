package main

import (
	"fmt"
)

func singleNumber(nums []int) int {
	eor := 0
	for i := 0;i<len(nums);i++{
		eor ^= nums[i]
		fmt.Println(i,eor)
	}
	return eor
}
func main() {
	a := []int{3,3,1,4,4,5,5}
	fmt.Println(singleNumber(a))
	
}
