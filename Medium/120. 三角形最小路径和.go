package main

import "fmt"

func minimumTotal(triangle [][]int) int {

	for i := 1 ; i < len(triangle) ; i ++{
		triangle[i][0] += triangle[i-1][0]
		triangle[i][i] += triangle[i-1][i-1]
	}
	for i := 2 ; i < len(triangle) ; i++{

		for j := 1;  j < i ; j ++{
			triangle[i][j] += mmmin(triangle[i-1][j],triangle[i-1][j-1])
		}

	}
	//fmt.Println(triangle)
	res := triangle[len(triangle)-1][0]
	for i := 1 ; i < len(triangle) ; i++{
		if triangle[len(triangle)-1][i] < res{
			res = triangle[len(triangle)-1][i]
		}
	}
	return res
}
func mmmin(a,b int)int{
	if a <= b {
		return a
	}else{
		return b
	}
}
func main() {
	a := [][]int{{2},{3,4},{6,5,7},{4,1,8,3}}
	fmt.Println(minimumTotal(a))
	
}
