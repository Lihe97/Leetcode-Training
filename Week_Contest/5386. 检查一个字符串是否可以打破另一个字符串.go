package main

import "fmt"

func checkIfCanBreak(s1 string, s2 string) bool {

	s1 = so(s1)
	s2 = so(s2)
	one ,two := true,true
	//fmt.Println(s1,s2)
	for i := 0 ; i < len(s1) ; i ++{
		if s2[i] < s1[i] && one == true{
			one = false
		}
		if s2[i] > s1[i] && two == true{
			two = false
		}
	}
	return one||two

}
func so(s string) string {
	mp := map[int]int{}
	res := []byte{}
	for i:= 0; i < len(s); i++ {
		mp[ int(s[i]-'a')]++
	}

	k := 0
	for i:=0; i<=25; i++ {
		if v,ok := mp[i]; ok {
			for j:=0; j < v; j++ {
				res = append(res,byte(i + 'a'))
				k++
			}
		}
	}
	return string(res)
}

func main() {

	//a := "cba"
	fmt.Println(checkIfCanBreak("abe","acd"))
}
