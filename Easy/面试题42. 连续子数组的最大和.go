package main

import "fmt"

func MaxSubArray(nums []int) int {

	dp := make([]int,len(nums))
	dp[0] = nums[0]
	res := 0
	for i := 1 ; i < len(nums) ; i ++{
		if dp[i-1] > 0{
			dp[i] = dp[i-1] + nums[i]
		}else{
			dp[i] = nums[i]
		}
		if dp[i] > res{
			res = dp[i]
		}
	}
	//fmt.Println(dp)
	return res
}
func main() {
	a := []int{-2,1,-3,4,-1,2,1,-5,4}
	fmt.Println(MaxSubArray(a))
	
}
