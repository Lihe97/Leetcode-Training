package main

import "fmt"

func canPermutePalindrome(s string) bool {

	mp := map[byte]int{}
	for i:= 0 ; i < len(s) ; i ++{
		mp[s[i]] ++
	}
	ji := 0
	for _,v := range  mp{
		if v % 2 != 0{
			ji ++
		}
	}
	if ji  > 1{
		return false
	}
	return true
}

func main() {
	fmt.Println(canPermutePalindrome("tactcoa"))
}
