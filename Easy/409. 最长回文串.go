package main

import "fmt"

func longestPalindrome(s string) int {
	mp := map[int32]int{}
	for _,s := range s{
		mp[s]++
	}
	one , two := 0,0
	for _,s := range mp{
		two += s/2
		one += s%2
	}
	if one > 0{
		return two*2+1
	}else {
		return two*2
	}
}
func main() {
	fmt.Println(longestPalindrome("abcccccddd"))
	
}
