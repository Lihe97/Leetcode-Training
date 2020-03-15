package main

import (
	"fmt"
	"sort"
)

func luckyNumbers (matrix [][]int) []int {
	height := len(matrix)
	width := len(matrix[0])
	res := []int{}
	for i:= 0 ; i <height ; i++{
		m := map[int]int{}
		temp := []int{}
		for j:= 0 ;j < width;j++{
			m[matrix[i][j]] = j
			temp = append(temp, matrix[i][j])
		}
		sort.Ints(temp)
		min := temp[0]
		minp := m[min]

		b := true

		for w := 0 ; w < height;w ++{
			if min < matrix[w][minp]{
				b = false
				break
			}
		}
		if b{
			res = append(res, min)
		}
	}

	return res

}
func main() {
	//a := [][]int{{1,10,4,2},{9,3,8,7},{15,16,17,12}}
	a := [][]int{{7,8},{1,2}}
	fmt.Println(luckyNumbers(a))
	
}
