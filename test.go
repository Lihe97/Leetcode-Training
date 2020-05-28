package main

import (
	"fmt"
	"strings"
)

func canWin(n int) bool {

	flag := make([]bool,n+1)
	flag[0] = false
	for i := 1 ; i <= n ; i++{
		flag[i] = false
		for j := 1 ; j <= 3 && i - j >= 0 && flag[i] == false; j++{
			flag[j] = !flag[i-j]
		}
	}
	return flag[n]
}

func main() {


	a  := "abc"

	b := strings.ToTitle(a)
	a = strings.Title(a)
	fmt.Println(a,b)
}


