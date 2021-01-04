package main

import (
	"fmt"
	"sort"
)

func sortByBits(arr []int) []int {


	sort.Slice(arr, func(i, j int) bool {
		if cal(arr[i]) == cal(arr[j]) {
			return arr[i] < arr[j]
		}
		return cal(arr[i]) < cal(arr[j])
	})
	return arr
}

func cal(x int) int{
	sum := 0
	//fmt.Print(x," ")
	for x != 0{
		sum += (x & 1)
		x = x >> 1
	}
	//fmt.Println(sum)
	return sum
}

func main() {

	fmt.Println(sortByBits([]int{0,1,2,3,4,5,6,7,8}))
}
