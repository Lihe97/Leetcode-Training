package main

import "fmt"

func arrangeCoins(n int) int {

	sum := 0
	res := 0
	for i := 1 ; sum + i <= n ; i ++{
		sum += i
		res ++
	}
	return res
}

func main() {
	fmt.Println(arrangeCoins(6))
	
}
