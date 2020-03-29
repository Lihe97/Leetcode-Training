package main

import "fmt"

func findTheDifference(s string, t string) byte {

	//s原来 t新

	mp := map[byte]int{}

	for i :=  0 ; i < len(s) ; i ++{
		mp[s[i]] ++
		mp[t[i]] --
	}
	mp[t[len(t)-1]]++
	for a,v := range mp{
		if v == 1{
			return a
		}
	}


	return 'a'
}

func main() {

	fmt.Println(findTheDifference("aabb","aabbb"))
}
