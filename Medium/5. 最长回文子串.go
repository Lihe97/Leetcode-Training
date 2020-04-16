package main

import "fmt"

//func longestPalindrome(s string) string {
//
//	dp := make([][]bool,len(s))
//	if len(s) == 1 {
//		return s
//	}
//	for i:= 0 ; i < len(s); i++{
//		dp[i] = make([]bool,len(s))
//		dp[i][i] = true
//	}
//	start,max := 0,1
//	for j := 1 ; j < len(s);j++{
//
//		for i := 0 ; i < j ; i++{
//			if s[i] == s[j]{
//
//				if j - i < 3{
//					dp[i][j] = true
//				}else{
//					dp[i][j] = dp[i+1][j-1]
//				}
//			}else{
//				dp[i][j] = false
//			}
//		if dp[i][j]{
//			if j-i+1>max{
//				max = j-i+1
//				start = i
//			}
//
//		}
//		}
//	}
//	return s[start:start+max]
//
//}

func longestPalindrome(s string) string {
	var res string
	for i := 0 ; i < len(s)*2 -1 ; i++{
		left := i/2
		right := left + i % 2

		for left >= 0 && right < len(s) && s[left] == s[right]{
			temp := s[left:right+1]
			if len(temp) > len(res){
				res = temp
			}
			left --
			right ++
		}
	}
	return res
}

func main() {
	//dp := make([][]bool,3)
	//arr := make([]bool,3)
	//for i := 0 ; i < len(dp); i ++{
	//	dp = append(dp, arr)
	//}
	//fmt.Println(dp)
	fmt.Println(longestPalindrome("abbb"))
	
}
