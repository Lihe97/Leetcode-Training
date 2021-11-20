package main

func findTargetSumWays(nums []int, target int) int {

	dp := make([]int,len(nums))

	if nums[0] ==target || nums[0] == -target{
		dp[0] = 1
	}

	for i := 1 ; i < len(nums) ; i ++{
		dp[i]
	}

}