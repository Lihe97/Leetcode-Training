package main

import (
	"fmt"
	"sort"
)

func permutation(s string) []string {

	flag := make([]bool,len(s))

	res := []string{}

	help(s,flag,"",&res)
	//fmt.Println(res)

	sort.Strings(res)
	//fmt.Println(res)

	ress := []string{}

	for i := 0; i < len(res); i++ {
		repeat := false
		for j := i + 1; j < len(res); j++ {
			if res[i] == res[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			ress = append(ress, res[i])
		}
	}

	return ress
}
func help(s string,flag []bool,cur string,res *[]string)  {

	if len(cur) == len(s){
		*res = append(*res,cur)
		return
	}
	for i := 0 ; i < len(s) ; i ++{
		if ! flag[i]{
			flag[i] = true
			help(s,flag,cur + string(s[i]),res)
			flag[i] = false
		}
	}


}

func main() {

	fmt.Print(permutation("abb"))
	
}
