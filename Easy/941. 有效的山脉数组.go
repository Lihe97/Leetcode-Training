package main

import "fmt"

func validMountainArray(A []int) bool {

	start , end := 0,len(A) - 1
	if len(A) < 3{
		return false
	}

	for start < len(A)-2 && A[start] < A[start+1]{
		start ++
	}
	for end > 1 && A[end-1] > A[end]{
		end --
	}


	return start == end

}

func main() {

	a := []int{0,2,3,3,2}
	fmt.Println(validMountainArray(a))
	
}
