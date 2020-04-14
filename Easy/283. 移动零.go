package main

import "fmt"

func moveZeroes(nums []int) []int{

	count := 0


	for i := 0 ;i < len(nums) - count; i++{
		if nums[i] == 0{
			nums = append(nums[0:i],nums[i+1:]...)
			nums = append(nums, 0)
			i --
			count ++
			//fmt.Println(nums,count,len(nums) - count)
		}
	}
	return nums
}
//func moveZeroes(nums []int)  {
//
//	i ,j := 0,0
//
//	for ; i < len(nums) ; i++{
//		if nums[i] != 0{
//			nums[i],nums[j] = nums[j],nums[i]
//			j++
//		}
//	}
//
//}
func main() {

	a := []int{2,1,0,0,0,0,0,0,3}
	fmt.Println(moveZeroes(a))
}
