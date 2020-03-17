package main

import "fmt"

func minCostClimbingStairs(cost []int) int {
	_len := len(cost)
	if _len < 2 {
		return 0
	} else {
		dp := make([]int, _len)
		dp[0] = cost[0]
		dp[1] = cost[1]
		for i := 2; i < _len; i++ {
			if dp[i - 1] < dp[i - 2] {
				dp[i] = dp[i - 1] + cost[i]
			} else {
				dp[i] = dp[i - 2] + cost[i]
			}
		}
		if dp[_len - 1] > dp[_len - 2] {
			return dp[_len - 2]
		} else {
			return dp[_len - 1]
		}
	}
}
func main() {
	a:= []int{0,1,2,2}
	fmt.Println(minCostClimbingStairs(a))
	
}
