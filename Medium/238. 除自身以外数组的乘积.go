package main

import "fmt"

func productExceptSelf(nums []int) []int {

	dp := [][]int{}
	res := make([]int,len(nums))
	for i := 0 ; i < len(nums) ; i++{
		temp := make([]int,len(nums))
		dp = append(dp, temp)
	}
	for i := 0 ; i < len(nums) ; i++{
		dp[i][i] = nums[i]
	}
	for i := 0 ; i < len(nums) ; i++{

		for j := i + 1 ;j < len(nums) ; j++{
			dp[i][j] = nums[j] * dp[i][j-1]
		}
	}
	res[0] = dp[1][len(nums)-1]
	res[len(nums)-1] = dp[0][len(nums)-2]
	for i := 1 ; i < len(nums) - 1 ; i++{
		res[i] = dp[0][i-1] * dp[i+1][len(nums)-1]
	}

	//fmt.Println(res)
	return res
}
func main() {
	a := []int{1,2,3,4}
	fmt.Println(productExceptSelf(a))
	
}
