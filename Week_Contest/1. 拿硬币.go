package main

import "fmt"

func minCount(coins []int) int {

	sum := 0
	for i := 0 ; i < len(coins); i ++{
		sum += coins[i] / 2
		sum += coins[i] % 2
	}
	return sum

}

func main() {

	a := []int{4,3,2}
	fmt.Println(minCount(a))
	
}
