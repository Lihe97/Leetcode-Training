package main

import "fmt"

func isSubsequence(s string, t string) bool {

	ss := []byte(s)
	tt := []byte(t)

	dp := make([]int,len(s) + 1)
	for i := 0 ; i < len(s); i ++{
		j:= 0
		temp := false
		if dp[i] != 0{
			j = dp[i] + 1
		}
		for   ; j < len(t) ; j ++{

			if tt[j] == ss[i]{
				temp = true
				dp[i + 1] = j
				break
			}
		}
		if temp == false{
			return false
		}
	}

	if dp[len(s)] != 0{
		return true
	}
	return false
}
func main() {
	fmt.Println(isSubsequence("abc",
	"agbdc"))
	
}
