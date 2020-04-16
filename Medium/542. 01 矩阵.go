package main

import (
	"fmt"
	"math"
)

func updateMatrix(matrix [][]int) [][]int {

	m := [][]int{}
	fmt.Println(m)
	for i := 0 ; i < len(matrix) ; i ++{
		temp := []int{}
		for j := 0; j < len(matrix[0]) ; j ++{
			temp = append(temp, math.MaxInt32)
		}
		m = append(m, temp)
	}
	fmt.Println(m)

	for i := 0 ; i < len(matrix) ; i ++{
		for j := 0 ;  j < len(matrix[0]) ; j ++{
			if matrix[i][j] == 0{
				m[i][j] = 0
			}
			if matrix[i][j] != 0{
				if i > 0 {
					m[i][j] = min(m[i-1][j] + 1,m[i][j])
				}
				if j > 0 {
					m[i][j] = min(m[i][j],m[i][j-1] + 1)
				}
			}
		}
	}
	for i := len(matrix) - 1 ; i >= 0 ; i --{
		for j := len(matrix[0]) - 1 ;  j >= 0 ; j --{
			if matrix[i][j] != 0{
				if i < len(matrix)-1 {
					m[i][j] = min(m[i+1][j] + 1,m[i][j])
				}
				if j < len(matrix[0]) - 1 {
					m[i][j] = min(m[i][j],m[i][j+1] + 1)
				}
			}
		}
	}
	return m

}

func main() {

	a := [][]int{{0,0,0},{0,1,0},{0,0,0}}
	fmt.Println(updateMatrix(a))
	
}
