package main

import "fmt"

func reverseLeftWords(s string, n int) string {
	return s[n:]+s[:n]


}
func main() {
	fmt.Println(reverseLeftWords("abcde",2))
	
}
