package main

import "fmt"

func hammingDistance(x int, y int) int {

	temp := x ^ y

	res := 0

	for temp > 0 {
		res += temp & 1
		temp = temp >> 1

	}
	return res
}

func main() {

	fmt.Println(hammingDistance(1,4))
	
}
