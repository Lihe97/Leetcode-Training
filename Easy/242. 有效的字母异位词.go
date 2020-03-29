package main

import "fmt"

func isAnagram(s string, t string) bool {

	if len(s) != len(t){
		return false
	}
	mp := map[byte]int{}
	for i := 0 ; i < len(s) ; i++{
		mp[s[i]] ++
		mp[t[i]] --
	}
	for _,v := range mp{
		if v != 0{
			return false
		}
	}
	return true
}

func main() {

	fmt.Println(isAnagram("anagrem","angaram"))
}
