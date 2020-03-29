package main

import (
	"fmt"
	"strings"
)

func firstUniqChar(s string) int {

	mp := map[byte]int{}
	res := len(s)


	for i := 0 ; i < len(s) ; i++{
		mp[s[i]] ++
	}
	for a,v := range mp{
		if v == 1{
			if strings.Index(s,string(a)) < res{
				res = strings.Index(s,string(a))
			}
		}
	}
	if res == len(s){
		return -1
	}
	return res

}

func main() {

	fmt.Println(firstUniqChar("aaccdd"))
	
}
