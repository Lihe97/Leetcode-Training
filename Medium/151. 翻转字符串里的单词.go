package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	s = strings.Trim(s," ")
	for i := 0 ; i < len(s) - 2; i++{
		for s[i] == ' ' && s[i+1] ==' '{
			s = s[:i] + s[i+1:]
		}
	}
	b := strings.Split(s," ")
	for i := 0 ; i < len(b)/2; i++{
		b[i],b[len(b)-i-1] = b[len(b)-i-1],b[i]
	}
	var res string
	for i := 0 ; i < len(b) ; i ++{
		res += b[i] + " "
	}
	res = res[:len(res)-1]

	return res
}
func main() {
	s := "   a     good     example  "

	//fmt.Println(len(s))
	//a = strings.ReplaceAll(a,"  ","")
	fmt.Println(reverseWords(s))
}
