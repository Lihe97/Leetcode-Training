package main

import "fmt"

func llengthOfLongestSubstring(s string) int {

	mp := map[byte]int{}
	left , right,res := 0,0,0

	for right < len(s){
		if k,ok := mp[s[right]]; ok && k >= left{
			left = k+1
			mp[s[right]] = right
		}else{
			mp[s[right]] = right
		}
		if right - left + 1 > res{
			res = right - left + 1
		}
		right ++
		//fmt.Println(left,right,res)
	}
	return res
}

func main() {

	fmt.Println(llengthOfLongestSubstring("abcabcbb"))
	
}
