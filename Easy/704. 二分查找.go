package main

import (
	"fmt"
	"time"
)

func search(nums []int, target int) int {
	res  := binary(0,len(nums)-1,nums,target)
	return res
}
func binary(left int ,right int ,nums []int,target int)int{
	//fmt.Println(left,right)

	for left <= right{
		mid := (left + right)/2
		time.Sleep(time.Second*2)

		if nums[mid] == target{
			return mid

		}else if nums[mid] > target{
			return binary(left,mid-1,nums,target)
		}else{
			return binary(mid+1,right,nums,target)
		}
	}
	return -1
}
func main() {
	a:= []int{-1,0,3,5,9,12}
	fmt.Println(search(a,9))
	
}
