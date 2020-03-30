package main

import (
	"fmt"
	"strings"
)


func findWords(words []string) []string {

	dic1 := "asdfghjklASDFGHJKL"
	dic2 := "qwertyuiopQWERTYUIOP"
	dic3 := "zxcvbnmZXCVBNM"
	var result []string
	for i:=0; i<len(words); i++  {
		n1, n2, n3 := 0, 0, 0
		for j:=0; j<len(words[i]); j++  {
			if  strings.Contains(dic1, string(words[i][j])) == true{
				n1++
			}
			if  strings.Contains(dic2, string(words[i][j])) == true{
				n2++

			}
			if  strings.Contains(dic3, string(words[i][j])) == true{
				n3++

			}
		}
		if n1 == len(words[i]) || n2 == len(words[i]) || n3 == len(words[i]){
			result = append(result, string(words[i]))
		}
	}
	return result

}
func main() {
	a := "abcd"
	fmt.Println(strings.Contains(a,"bd"))
}
