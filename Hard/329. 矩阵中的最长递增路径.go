package main

import (
	"fmt"
)

func longestIncreasingPath(matrix [][]int) int {
	if len(matrix) == 0{
		return 0
	}
	res := 0
	cache := [][]int{}
	for i := 0 ; i < len(matrix) ; i ++{
		temp2 := make([]int,len(matrix[0]))
		cache = append(cache,temp2)
	}
	//fmt.Println(cache)
	for i := 0 ; i < len(matrix) ; i ++{
		for j := 0 ; j < len(matrix[0]) ; j ++{

			res = max(res,t2(matrix,i,j,cache))

		}
	}
	//fmt.Println(cache)


	return res
}

func t2(matrix [][]int, i int, j int,cache [][]int) int {
	//fmt.Println(i,j)
	if 0 != cache[i][j] {
		return cache[i][j]
	}
	//fmt.Println(cache[i][j])

	mx := []int{1,0,-1,0}
	my := []int{0,1,0,-1}
	for k := 0 ; k < len(mx) ; k ++{
		x := i + mx[k]
		y := j + my[k]
		if x >= 0 && y >= 0 && x < len(matrix) && y < len(matrix[0]) && matrix[x][y] > matrix[i][j]{
			//fmt.Println(x,y)
			//time.Sleep(time.Second)
			cache[i][j] = max(cache[i][j],t2(matrix,x,y,cache))
		}
	}
	cache[i][j] += 1

	return cache[i][j]
}
func max(a,b int) int {
	if a > b{
		return a
	}else{
		return b
	}
}

func main() {

	m := [][]int{{9,9,4},{6,6,8},{2,1,1}}

	fmt.Println(longestIncreasingPath(m))
}