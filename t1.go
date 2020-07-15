package main

import (
	"fmt"
)
func t1(n int) int{


	if n == 1 {
		return 1
	}else if n == 2{
		return 2
	}
	return t1(n-1) + t1(n-2)
}
func main() {
	var n int

	fmt.Scan(&n)

	fmt.Println(t1(n))
}