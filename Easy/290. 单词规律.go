package main

import (
	"fmt"
	"strings"
)

func wordPattern(pattern string, str string) bool {

	b := strings.Split(str," ")
	if len(b) != len(pattern){
		return false
	}
	mp := map[byte]string{}

	for i := 0 ; i < len(pattern) ; i++{

		//fmt.Println('1',mp[pattern[i]])
		if mp[pattern[i]] == ""{
			mp[pattern[i]] = b[i]
		}else{
			//fmt.Println(mp)
			if mp[pattern[i]] != b[i]{
				return false
			}
		}
	}
	m := map[string]int{}
	for _,b := range mp{
		m[b] ++
		if m[b] > 1{
			return false
		}
	}
	//fmt.Println(mp)
	return true
}
func main() {
	//a := strings.Split("a b c"," ")
	fmt.Println(wordPattern("abba","aa v v aa"))
}
