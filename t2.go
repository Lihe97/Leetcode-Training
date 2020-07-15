package main

import (
	"fmt"
)
var max int = 1

func t2(matrix [][]int,curx,cury int,temp int,flag [][]bool)  {
	mx := []int{1,0,-1,0}
	my := []int{0,1,0,-1}
	//fmt.Println(curx,cury,flag,temp)

	flag[curx][cury] = true

	for i := 0 ; i < 4 ; i ++{
		newx := curx + mx[i]
		newy := cury + my[i]
		if newx >= 0 && newx < len(matrix) && newy >= 0 && newy < len(matrix[0]) && !flag[newx][newy]{
			if matrix[newx][newy] > matrix[curx][cury]{
				temp ++
				if temp > max{
					max = temp
				}
				t2(matrix,newx,newy,temp,flag)
				flag[newx][newy] = false
				temp --
			}
		}
	}

}

func main() {
	var m int
	var n int
	fmt.Scan(&m)
	fmt.Scan(&n)

	matrix := [][]int{}
	flag := [][]bool{}
	for i := 0 ; i < m ; i ++{
		temp := make([]int,n)
		temp2 := make([]bool,n)
		matrix = append(matrix,temp)
		flag = append(flag,temp2)
	}


	for i := 0 ; i < m ;i ++{
		for j := 0 ; j < n ; j ++{
			fmt.Scan(&matrix[i][j])
		}
	}

	for i := 0 ; i < len(matrix) ; i ++{
		for j := 0 ; j < len(matrix[0]) ; j ++{
			t2(matrix,i,j,1,flag)
			flag[i][j] = false
			//fmt.Println("?")
		}
	}
	fmt.Println(max)
}