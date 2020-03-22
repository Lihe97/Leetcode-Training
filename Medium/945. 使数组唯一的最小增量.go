package main

import "sort"

func minIncrementForUnique(A []int) int {

	sort.Ints(A)
	res := 0
	for i := 1 ; i < len(A) ; i++{
		if A[i] <= A[i-1]{
			res += A[i-1] - A[i] + 1
			A[i] = A[i-1] + 1
		}
	}
	return res
}
func main() {
	
}
