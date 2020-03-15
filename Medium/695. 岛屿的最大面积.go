package main

import (
	"fmt"
)

var xx  = []int{-1,1,0,0}
var yy  = []int{0,0,-1,1}
var q int
func maxAreaOfIsland(grid [][]int) int {
	max := 0
	height := len(grid)
	if height == 0{
		return 0
	}
	width := len(grid[0])
	work := [][]int{}


	for i:= 0 ; i< height; i++{
		temp := []int{}
		for j:= 0 ; j < width;j++{
			temp =append(temp, 0)
		}
		work = append(work, temp)
	}
	for i := 0 ; i < height; i++{
		for j:= 0 ; j < width ; j++{
			if grid[i][j] == 1 && work[i][j] == 0{
				DFSS(work,grid,i,j)
				//fmt.Println(q)
				if q > max{
					max = q
				}
				q = 0
			}
		}

	}
	//fmt.Println(work)
	return max
}
func DFSS(work [][]int,grid [][]int,i int,j int){
	work[i][j] = 1
	q++
	for k := 0 ; k < 4; k ++{

		newx := i + xx[k]
		newy := j + yy[k]
		if newx < 0 || newx >= len(work) || newy < 0 || newy >= len(work[0]){
			continue
		}
		//fmt.Println(newx,newy)
		//fmt.Println("work:",work[newx][newy])
		//fmt.Println("grid:",grid[newx][newy])
		if work[newx][newy] == 0 && grid[newx][newy] == 1{

			DFSS(work,grid,newx,newy)
		}
	}
}
func main() {


	a := [][]int{{1,0,1,1,0},
		          {1,1,0,1,0},
				  {1,0,0,0,0},
				  {0,0,0,0,0}}
	fmt.Println(maxAreaOfIsland(a))

}