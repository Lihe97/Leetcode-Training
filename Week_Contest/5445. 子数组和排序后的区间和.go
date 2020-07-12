package main

import (
	"fmt"
	"sort"
)

func rangeSum(nums []int, n int, left int, right int) int {

	temp := []int{}

	dp := make([][]int,len(nums))

	dp[0] = append(dp[0],nums[0])
	temp = append(temp,dp[0]...)
	for i := 1 ; i < len(nums) ; i ++{
		for j := 0 ; j  < len(dp[i-1]) ; j ++{
			dp[i] = append(dp[i],dp[i-1][j]+nums[i])

		}
		dp[i] = append(dp[i],nums[i])
		temp = append(temp,dp[i]...)
	}
	sort.Ints(temp)
	//fmt.Println(temp)
	res := 0
	for i := left - 1 ; i < right ; i ++{
		res += temp[i]
	}
	return res
}

func main() {
	fmt.Println(rangeSum([]int{1,2,3,4},4,1,5))
	
}
