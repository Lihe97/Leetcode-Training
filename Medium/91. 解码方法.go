package main

import "fmt"

func numDecodings(s string) int {
	if s == "" || s == "0"{
		return 0
	}
	dp := make([]int,len(s))

	if s[0] == '0'{
		return 0
	}
	dp[0] = 1
	for i := 1 ; i < len(s) ; i++{
		if s[i]!= '0' {
			dp[i] = dp[i-1]
		}
		if s[i-1] == '1'  || (s[i-1] == '2') && s[i] <= '6'{
			if i >=2 {
			dp[i] += dp[i-2]
			}else{
				dp[i]++
			}
		}
	}
	return dp[len(s)-1]
}
func main() {

	fmt.Println(numDecodings("12"))
	
}
