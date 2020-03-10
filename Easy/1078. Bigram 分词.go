package main

import (
	"fmt"
	"strings"
)

func findOcurrences(text string, first string, second string) []string {
	res := []string{}
	s := strings.Split(text, " ")
	for i:= 0 ; i < len(s) -2 ;{
		if s[i] == first && s[i+1] == second{
			res= append(res, s[i+2])
			i += 2
		}else{
			i++
		}
	}
	return res
}
func main() {
	fmt.Println(findOcurrences("alice is a good girl she is a good student","a","good"))
}
