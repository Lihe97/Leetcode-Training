package main

import "fmt"

func maxSubArray(nums []int) int {
	l := len(nums)
	max := nums[0]
	temp := nums[0]
	for i:= 1 ; i < l ; i++{
		if temp < 0{
			temp = nums[i]
		}else {
			temp += nums[i]
		}
		if temp > max{
			max = temp
		}
	}
	return max
}
func main() {
	a := []int{-2,1,-3,4,-1,2,1,-5,4}
	fmt.Println(maxSubArray(a))
}
