package main

import "fmt"

func peakIndexInMountainArray(A []int) int {

	left := 0
	right := len(A) -1

	for left < right {
		mid := (left + right)/2
		if A[mid] < A[mid+1]{
			left = mid+1
		}else{
			right = mid
		}
	}
	return left
}

func main() {

	a := []int{1,2,3,4,1}
	fmt.Println(peakIndexInMountainArray(a))
}
