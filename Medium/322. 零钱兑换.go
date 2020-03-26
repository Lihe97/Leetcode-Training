package main

import (
	"fmt"
)
//dp问题
//根据样例[1,2,5] dp[i] = min(dp[i-1],dp[i-2],dp[i-5])+1
//test1
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 0; i <= amount; i++ {
		dp[i] = amount+1
	}
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		for j := 0; j < len(coins); j++ {
			if coins[j] <= i {
				dp[i] = mmmmin(dp[i], dp[i - coins[j]] + 1)
			}
		}
	}
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

func mmmmin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	a := []int{186,419,80,408}
	fmt.Println(coinChange(a,6249))
}
