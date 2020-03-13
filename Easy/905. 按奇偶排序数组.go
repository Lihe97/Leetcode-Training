package main

import "fmt"

//偶数在前，奇数在后
func sortArrayByParity(A []int) []int {
	a := make([]int,len(A))
	l := len(A)
	start , end := 0, l-1
	for i := 0 ; i < l; i++{
		if A[i]%2 == 0{
			a[start] = A[i]
			start++
		}else{
			a[end] = A[i]
			end --

		}
		//fmt.Println(a)
	}
	return a
}
func main() {
	b := []int{1,2,3,4,5,2,2,8,9,1}
	fmt.Println(sortArrayByParity(b))
	
}
