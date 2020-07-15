package main

import "fmt"

func letterCombinations(digits string) []string {

	res := []string{}
	if len(digits) == 0{
		return res
	}



	mp := map[byte][]byte{}
	mp['2'] = []byte{'a','b','c'}
	mp['3'] = []byte{'d','e','f'}
	mp['4'] = []byte{'g','h','i'}
	mp['5'] = []byte{'j','k','l'}
	mp['6'] = []byte{'m','n','o'}
	mp['7'] = []byte{'p','q','r','s'}
	mp['8'] = []byte{'t','u','v'}
	mp['9'] = []byte{'w','x','y','z'}


	dd(mp,0,"",digits,&res)




	return res
}
func dd(mp map[byte][]byte,cur int,temp string,digits string,res *[]string)  {
	if cur == len(digits) {
		*res = append(*res,temp)
		return
	}

	for i := 0 ; i < len(mp[digits[cur]]) ; i ++{
		dd(mp,cur+1,temp + string(mp[digits[cur]][i]),digits,res)
	}
}

func main() {

	fmt.Println(letterCombinations("2345"))
	
}
