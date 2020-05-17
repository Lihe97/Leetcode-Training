package main

import (
	"fmt"
	"strings"
)

func main() {


	a  := "abc"

	b := strings.ToTitle(a)
	a = strings.Title(a)
	fmt.Println(a,b)
}


