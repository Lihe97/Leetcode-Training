package main

import "fmt"

//func findMaxConsecutiveOnes(nums []int) int {
//
//	res := 0
//	temp := 0
//	for i := 0 ; i < len(nums) ; i++{
//		if nums[i] == 0{
//			if temp > res{
//				res = temp
//			}
//			temp = 0
//		}
//		if nums[i] == 1{
//			temp++
//		}
//	}
//	if temp > res {
//		res = temp
//	}
//	return res
//}
func main() {

	a := []int{1,1,0,1,1,1}
	fmt.Println(findMaxConsecutiveOnes(a))

}
