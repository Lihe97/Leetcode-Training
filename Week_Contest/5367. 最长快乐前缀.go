package main

import "fmt"

func longestPrefix(s string) string {



	for i:= len(s) ; i > 0;  i --{
		if s[0:i-1] == s[len(s)-i+1:]{
			return s[0:i-1]
		}
	}
	return ""

}
func main() {

	fmt.Println(longestPrefix("leetcodeleet"))
}
