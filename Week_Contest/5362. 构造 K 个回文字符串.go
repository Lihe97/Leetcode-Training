package main

import "fmt"

func canConstruct(s string, k int) bool {

	mp := map[byte]int{}

	for i := 0 ; i < len(s) ; i ++{
		mp[s[i]]++
	}
	ji := 0
	ou := 0
	for _,b := range mp{
		if b % 2 == 1{
			ji ++
		}else{
			ou ++
		}
	}
	if ji != 0{
		if k >= ji && k <= len(s){
			return true
		}else{
			return false
		}
	}
	if k <= len(s){
		return true
	}else{
		return false
	}


}

func main() {
		fmt.Println(canConstruct("yzyzyzyzyzyzyzy",2))
}
