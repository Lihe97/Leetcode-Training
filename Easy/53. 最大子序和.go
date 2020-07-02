package main

import "fmt"

//func maxSubArray(nums []int) int {
//	l := len(nums)
//	max := nums[0]
//	temp := nums[0]
//	for i:= 1 ; i < l ; i++{
//		if temp < 0{
//			temp = nums[i]
//		}else {
//			temp += nums[i]
//		}
//		if temp > max{
//			max = temp
//		}
//	}
//	return max
//}

func maxSubArray(nums []int) int{

	dp := make([]int,len(nums))
	dp[0] = nums[0]
	res := dp[0]
	for i := 1 ; i < len(nums) ; i ++{
		//if nums[i] >= 0 && dp[i-1] > 0{
		//	dp[i] = dp[i-1] + nums[i]
		//}else if nums[i] > 0 && dp[i-1] <= 0{
		//	dp[i] = nums[i]
		//}else if nums[i] < 0  && dp[i-1] > 0{
		//	dp[i] = dp[i-1] + nums[i]
		//}else if nums[i] < 0 && dp[i-1] < 0{
		//	dp[i] = nums[i]
		//}
		if dp[i-1] >= 0{
			dp[i] = dp[i-1] + nums[i]
		}else{
			dp[i] = nums[i]
		}
		if dp[i] > res{
			res = dp[i]
		}

	}
	return res
}

func main() {
	a := []int{-2,1,-3,4,-1,2,1,-5,4}
	fmt.Println(maxSubArray(a))
}
