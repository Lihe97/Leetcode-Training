package main

import "fmt"

func reverse(x int) int {

	res := 0
	a := 2<<30 -1

	for x != 0{
		res = res * 10 + x % 10
		x = x / 10
	}

	if res > a  || res < -a + 1{
		return 0
	}
	return res
}
func main() {
	fmt.Println(reverse(1534236469))
	
}
