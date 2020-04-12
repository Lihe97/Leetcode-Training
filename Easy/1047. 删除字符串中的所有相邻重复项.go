package main

import "fmt"

func removeDuplicates(S string) string {

	stack := []string{}

	for  i := 0 ; i < len(S) ; i ++{
		if len(stack) != 0 && stack[len(stack)-1] == string(S[i]){
			stack = stack[:len(stack)-1]
		}else{
			stack = append(stack, string(S[i]))
		}
	}
	var res string
	for i := 0 ; i < len(stack); i ++{
		res += stack[i]
	}
	return  res
}

func main() {

	fmt.Println(removeDuplicates("abbaca"))
}
