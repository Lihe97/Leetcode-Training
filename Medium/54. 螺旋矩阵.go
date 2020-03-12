package main

import (
	"fmt"
)

func spiralOrder(matrix [][]int) []int {
	res := []int{}
	//
	height := len(matrix)
	if height==0{
		return nil
	}
	if height == 1{
		return matrix[0]
	}
	width := len(matrix[0])
	i , j := 0,0


	//fmt.Println(matrix[i:height-i])
	for i <= height-2 && j <=width-2{

		temp := matrix[i:height-i]

		OneLayer(&res,temp,i,width - 2*j)
		i ++
		j++
		//fmt.Println("resr",res)
	}




	return  res

}
func OneLayer(res *[]int,m [][]int,h int,w int){


	height := len(m)

	width := len(m[0])
	//fmt.Println(w,h)
	for i := h ; i <w+h ; i ++{
		*res = append(*res, m[0][i])
	}
	for i := 1  ; i < height ; i++{

		*res = append(*res, m[i][w+h-1])
	}
	if height > 1{
	for i := width - h-2 ; i >= h ; i--{
		*res = append(*res, m[height-1][i])
	}
	}
	for i := height - 2  ; i >=1 ; i--{
		*res = append(*res, m[i][h])
	}
	//fmt.Println(res)


}

func main() {
	//a := [][]int{{1,2,3,4},{2,3,4,5}}
	a := [][]int{{1},{5}}
	//fmt.Println(a)
	fmt.Println(spiralOrder(a))
	
}
