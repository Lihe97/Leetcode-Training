package main

import "fmt"

func lengthOfLIS(nums []int) int {
	dp := make([]int,len(nums))
	if len(nums) == 0{
		return 0
	}
	dp[0] = 1
	res:=1
	for i := 1 ; i < len(dp) ; i++{

		m := 0
		for j:=0 ; j <= i-1 ; j++{
			if nums[i] > nums[j]{
				m = max(m,dp[j])
			}
		}
		dp[i] = m +1
		 res = max(res,dp[i])


	}
	return res
}
func max(a,b int)int{
	if a>b{
		return a
	}else{
		return b
	}
}
func main() {
	a := []int{1,3,6,7,9,4,10,5,6}
	fmt.Println(lengthOfLIS(a))
	
}
