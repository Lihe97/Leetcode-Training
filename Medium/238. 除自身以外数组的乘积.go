package main

import "fmt"

//func productExceptSelf(nums []int) []int {
//
//	left := make([]int,len(nums))
//	left[0] = 1
//	for  i := 1 ; i < len(nums);i++{
//		left[i] = nums[i-1] * left[i-1]
//	}
//	right := make([]int,len(nums))
//	right[len(nums)-1] = 1
//	for j := len(nums) -2 ; j >= 0 ; j --{
//		right[j] = nums[j+1] * right[j+1]
//	}
//	res := make([]int,len(nums))
//	for i := 0 ; i < len(nums) ; i++{
//		res[i] = left[i] * right[i]
//	}
//
//	return res
//
//}
func productExceptSelf(nums []int) []int {


	res := make([]int,len(nums))
	res[0] = 1
	for  i := 1 ; i < len(nums);i++{
		res[i] = nums[i-1] * res[i-1]
	}
	right := make([]int,len(nums))
	right[len(nums)-1] = 1
	for j := len(nums) -2 ; j >= 0 ; j --{
		right[j] = nums[j+1] * right[j+1]
		res[j] = res[j] * right[j]
	}
	return res
}
func main() {
	a := []int{1,2,3,4}
	fmt.Println(productExceptSelf(a))
	
}
