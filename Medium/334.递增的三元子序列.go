package main

import (
	"fmt"

)
//给定一个未排序的数组，判断这个数组中是否存在长度为 3 的递增子序列。
//区间划分，min最小值，mx第二大的值

func increasingTriplet(nums []int) bool {
	//min := 0
	mx := int(^uint(0) >> 1)
	min :=int(^uint(0) >> 1)
	for x := range nums{
		if nums[x]<=min {
			min = nums[x]
		}else if nums[x] <= mx{
			mx = nums[x]
		}else{
			return true
		}
	}
	return false
}
func main() {
	a:= []int{3,5,2,1}
	fmt.Println(increasingTriplet(a))
}
