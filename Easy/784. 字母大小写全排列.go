package main

import (
	"fmt"
)

func letterCasePermutation(S string) []string {

	res := []string{}
	var temp string
	dfs(S,&res,temp)

	return res
}
func dfs(S string,res *[]string,temp string){

	if len(temp) == len(S){
		*res = append(*res, temp)
		return
	}
	temp += string(S[len(temp)])
	dfs(S,res,temp)
	temp = string(temp[0:len(temp)-1])
	if (S[len(temp)] >= 'a' && S[len(temp)] <= 'z') || (S[len(temp)] >= 'A' && S[len(temp)] <= 'Z'){
		temp += string(S[len(temp)] ^ ' ')
		dfs(S,res,temp)
	}


}

func main() {


	fmt.Println(letterCasePermutation("ab1c"))
	
}
