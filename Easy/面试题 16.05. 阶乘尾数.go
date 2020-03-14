package main

import "fmt"

func trailingZeroes(n int) int {

	res := 0
	for n >0{
		res += n/5
		n = n/5
	}
	return  res
}
func main() {
	fmt.Println(trailingZeroes(130))
}
