package main

import "fmt"

func isMatch(s string, p string) bool {

	if s == p{
		return true
	}
	if len(s) != 0 && len(p) == 0{
		return false
	}
	dp := make([][]bool,len(s)+1)
	for i := 0 ; i < len(dp) ;i ++{
		temp := make([]bool,len(p)+1)
		dp[i] = temp
	}
	dp[0][0] = true

	for i := 0 ; i < len(p) ; i++{
		fmt.Println(string(p[i]))
		if p[i] == '*' && dp[0][i-1]{
			dp[0][i+1] = true
		}
	}
	fmt.Println(dp)

	for i := 0 ; i < len(s) ; i ++{
		for j := 0 ; j  < len(p) ; j ++{
			if p[j] == '.' || p[j] == s[i]{
				dp[i+1][j+1] = dp[i][j]
			}
			if p[j] == '*'{
				if p[j-1] != s[i] && p[j-1] != '.'{
					dp[i+1][j+1] = dp[i+1][j-1]
				}else{
					dp[i + 1][j + 1] = (dp[i + 1][j] || dp[i][j + 1] || dp[i + 1][j - 1]);
				}
			}
		}
	}
	return dp[len(s)][len(p)]
}
func main() {

	fmt.Println(isMatch("aab","c*a*b"))
	
}
