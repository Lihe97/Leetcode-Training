package main

import (
	"fmt"
)

func reversePairs(nums []int) int {

	if len(nums) < 2{
		return 0
	}
	count := 0
	mergesort(nums,&count)
	return count
}
func mergesort(nums []int,count *int) []int {
	if len(nums) < 2{
		return nums
	}
	i := len(nums) / 2
	left := mergesort(nums[:i],count)
	right := mergesort(nums[i:],count)

	return merge(left,right,count)
}
func merge(left,right []int,count *int) []int{
	res := []int{}
	m,n := 0,0
	for m < len(left) && n < len(right){
		if left[m] > right[n]{
			res = append(res, right[n])
			n++
			*count  += len(left) - m
		}else{
			res = append(res, left[m])
			m++
		}
	}
	res = append(res, left[m:]...)
	res = append(res,right[n:]...)
	return res
}
func main() {

	a := []int{2,4,5,7,6,3}
	fmt.Println(reversePairs(a))
	
}
