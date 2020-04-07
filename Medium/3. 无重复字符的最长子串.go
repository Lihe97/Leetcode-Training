package main

import "fmt"

func lengthOfLongestSubstring(s string) int {

	max := 0

	temp := []byte{}
	temp = append(temp, s[0])
	for i := 1 ; i < len(s) ; i ++{
		for  j := len(temp)-1; j >= 0 ; j --{
			if temp[j] == s[i]{
				temp = temp[j+1:]
				break
			}
		}
		temp = append(temp, s[i])
		if len(temp) > max{
			max = len(temp)
		}
	}
	return max
}
func main() {

	fmt.Println(lengthOfLongestSubstring("pwwlew"))
	
}
