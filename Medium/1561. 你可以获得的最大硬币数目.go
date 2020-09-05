package main

import "sort"

func maxCoins(piles []int) int {

	sort.Ints(piles)
	n := len(piles)
	res := 0
	for j := n - 1 ; j > n/3 ; j -= 2{
		res += piles[j-1]
	}

	return res
}


func main() {
	
}
