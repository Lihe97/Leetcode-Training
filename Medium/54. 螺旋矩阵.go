package main

import (
	"fmt"
	"math"
)


func spiralOrder(matrix [][]int) []int {
	res := []int{}
	height := len(matrix)
	width := len(matrix[0])
	if height==0{
		return nil
	}
	i , j := 0,0
	for i <= int(math.Ceil(float64(height)/2-1)) && j <=int(math.Ceil(float64(width)/2-1)){
		temp := matrix[i:height-i]
		OneLayer(&res,temp,i,width - 2*j)
		i ++
		j++
	}
	return  res
}
func OneLayer(res *[]int,m [][]int,h int,w int){
	height := len(m)
	width := len(m[0])
	if w == 1{
		for i := 0 ; i <height ; i++{
			*res = append(*res, m[i][h])
		}
	}else{
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
	}
}


func main() {
	//a := [][]int{{1,2,3},{4,5,6},{7,8,9}}
	//a := [][]int{{1,2,3,4},{5,6,7,8},{9,10,11,12},{13,14,15,16}}
	//a := [][]int{{2,3,4},{5,6,7},{8,9,10},{11,12,13},{14,15,16}}
	a:=[][]int{{1,2,3,4,5},{6,7,8,9,10},{11,12,13,14,15},{16,17,18,19,20},{21,22,23,24,25}}
	fmt.Println(spiralOrder(a))
	
}
