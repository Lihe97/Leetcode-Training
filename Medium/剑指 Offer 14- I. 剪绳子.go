package main

func cuttingRope(n int) int {

	dp := make([]int,n + 1)
	dp[2] = 1
	for i := 3 ; i <= n ; i ++{
		for j := 2 ; j < i ; j ++{
			dp[i] = max(dp[i],max(j * (i-j),dp[i-j] * j))
		}
	}
	return dp[n]
}
func max(a,b int) int {
	if a >= b{
		return a
	}else{
		return  b
	}

}