package main

import "fmt"

func transpose(A [][]int) [][]int {
	height := len(A)
	width := len(A[0])
	res := [][]int{}

	for i := 0 ; i < width;i++{

		temp := []int{}
		for j := 0 ; j < height ; j++{
			temp = append(temp, A[j][i])
		}
		res = append(res, temp)

	}
	return res
}
func main() {
	a := [][]int{{1}}
	fmt.Println(transpose(a))
	
}
