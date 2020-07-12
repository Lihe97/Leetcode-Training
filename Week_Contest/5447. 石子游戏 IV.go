package main

import "math"

func winnerSquareGame(n int) bool {

	dp := make([]bool,n+1)
	dp[0] = false




	for i := 1 ; i  <= n ; i ++{
		dp[i] = false

		for j := 1; j <= int(math.Sqrt(float64(i))) && i-j*j >= 0 && dp[i] == false; j ++{
			dp[i] = !dp[i-j*j]
		}
	}
	return dp[n]





	//if math.Sqrt(float64(n-1)) == 0{
	//	return false
	//}
	return false


}


func main() {
	
}
