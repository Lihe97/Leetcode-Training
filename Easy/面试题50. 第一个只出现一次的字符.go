package main

import "fmt"

func firstUniqChar(s string) byte {

	mp := map[byte]int{}
	for i := 0 ; i < len(s) ; i ++{
		mp[s[i]] ++
	}
	for i := 0 ; i < len(s) ; i ++{
		if mp[s[i]] == 1{
			return s[i]
		}
	}
	return ' '
}

func main() {
	fmt.Println(ffirstUniqChar("aabd"))
	
}
