package main

import "fmt"

func isHuiwen(s string)bool{

	left := 0
	right := len(s) - 1
	for left < right{
		if s[left] != s[right]{
			return false
		}else{
			left++
			right--
		}
	}
	return true
}
func validPalindrome(s string) bool {

	left ,right := 0,len(s)-1
	for left < right{
		if s[left] != s[right]{
			return isHuiwen(s[left+1:right+1]) || isHuiwen(s[left:right])
		}else{
			right--
			left++
		}

	}
	return true
}

func main() {

	fmt.Println(isHuiwen("abcba"))
	
}
