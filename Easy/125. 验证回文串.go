package main

import "fmt"

func isPalindrome(s string) bool {

	temp := []byte{}
	for i := 0 ; i < len(s) ; i ++{

		//fmt.Println(s[i])

		if s[i] >= 'a' && s[i] <= 'z'{
			temp = append(temp,s[i])
		}else if s[i] >= 'A' && s[i] <= 'Z'{
			temp = append(temp,s[i]+32)
		}else if s[i] >= '0' && s[i] <= '9'{
			temp = append(temp,s[i])
		}
	}
	//fmt.Println(temp)

	for i := 0 ; i < len(temp)/2 ; i++{
		//fmt.Println(temp[i],temp[len(temp)-i-1])
		if temp[i] != temp[len(temp) - i - 1]{
			return false
		}
	}


	return true
}

func main() {

	a := "0a"
	fmt.Println(isPalindrome(a))
	
}
