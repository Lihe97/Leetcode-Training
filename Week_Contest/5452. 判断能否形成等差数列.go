package main

import "sort"

func canMakeArithmeticProgression(arr []int) bool {

	sort.Ints(arr)
	temp := arr[1] - arr[0]
	for i := 2 ; i < len(arr) ; i ++{
		if arr[i] - arr[i-1] != temp{
			return false
		}
	}
	return true
}

func main() {
	
}
