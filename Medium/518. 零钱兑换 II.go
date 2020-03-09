package main

import "fmt"

func change(amount int, coins []int) int {

	dp := make([]int,amount+1)
	dp[0] = 1
	for x := range coins{
		for j := coins[x] ; j <= amount;j++{
			dp[j] += dp[j-coins[x]]
		}
	}
	return dp[amount]
}
func main() {
	fmt.Println()
}
