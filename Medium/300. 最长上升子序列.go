package main

import "fmt"

func lengthOfLIS(nums []int) int {
	dp := make([]int,len(nums))
	dp[0] = 1
	res := 1
	for i := 1 ; i < len(nums) ; i++{
		max := 0
		for j := 0;  j < i ; j++{
			if nums[i] > nums[j]{
				if dp[j] > max{
					max = dp[j]
				}
			}
		}
		dp[i] = max+1
		if dp[i] > res{
			res = dp[i]
		}
	}

	return res
}

func main() {
	//a := []int{1,3,6,7,9,4,10,5,6}
	a := []int{10,9,2,5,3,7,101,18}
	fmt.Println(lengthOfLIS(a))
	
}
