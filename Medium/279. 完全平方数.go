package main

import "fmt"


func numSquares(n int) int {

	dp := make([]int,n+1)
	dp[0] = 0
	dp[1] = 1
	dp[2] = 2
	dp[3] = 3
	dp[4] = 1
	for i := 5 ; i <= n ; i ++{
		dp[i] = 100
		for j := 2; j * j <=i ; j++{
			if j * j == i{
				dp[i] = 1
			}else  if dp[i- j*j] +1 < dp[i]{

				dp[i] = dp[i- j*j] +1
			}
		}
	}
	//fmt.Println(dp)
	return dp[n]
}

func main() {
	fmt.Println(numSquares(12))
}
