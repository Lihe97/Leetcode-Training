package main

import "fmt"

func findLucky(arr []int) int {
	mp := map[int]int{}
	res := -1
	for i := 0 ; i < len(arr) ; i ++{
		mp[arr[i]]++
	}
	for a,b := range mp{
		if a == b && a > res{
			res = a
		}
	}
	return res
}
func main() {

	a := []int{2,2,3,4}
	fmt.Println(findLucky(a))
	
}
