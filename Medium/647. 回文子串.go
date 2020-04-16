package main

import "fmt"

//func countSubstrings(s string) int {
//
//	dp := [][]bool{}
//	for i := 0 ; i < len(s); i ++{
//		temp := make([]bool,len(s))
//		dp = append(dp, temp)
//	}
//	res := 0
//	for j := 0 ; j < len(s) ; j++{
//		for i := 0 ; i <= j ; i ++{
//			if s[i] == s[j] && ( j - i < 2 || dp[i+1][j-1]){
//				dp[i][j] = true
//				res ++
//			}
//		}
//	}
//
//	//fmt.Println(dp)
//	return res
//}
func countSubstrings(s string) int {

	res := 0
	for i := 0 ; i < 2 * len(s) - 1; i++{
		left := i/2
		right := left + i % 2

		for left >= 0 && right < len(s) && s[left] == s[right]{
			res ++
			left --
			right ++
		}
	}
	return res
}

func main() {

	fmt.Println(countSubstrings("aaa"))
	
}
