package main

import (
	"fmt"
	"math"
)

func integerBreak(n int) int {
	if n == 0{
		return 0
	}
	if n == 1{
		return 1
	}
	if n == 2{
		return 2
	}
	a := n/3
	b := n%3
	if b == 1{
		return int(math.Pow(3, float64(a-1))*4)
	}else if b == 0{
		return int(math.Pow(3, float64(a)))
	}else {
		return int(math.Pow(3, float64(a))*2)
	}
}
func main() {
	fmt.Println(integerBreak(7))
	
}
