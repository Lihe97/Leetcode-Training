package main

import "fmt"

func minPathSum(grid [][]int) int {

	height := len(grid)
	width := len(grid[0])

	for i := 1 ; i < height ; i ++{
		grid[i][0] += grid[i-1][0]
	}
	for i := 1 ; i < width ; i ++{
		grid[0][i] += grid[0][i-1]
	}
	for i := 1 ; i < height ; i ++{

		for j := 1 ; j < width ; j ++{
			grid[i][j] += mmin(grid[i-1][j],grid[i][j-1])
		}
	}
	return grid[height-1][width-1]
}
func mmin(a,b int)int{
	if a <= b {
		return a
	}else{
		return b
	}
}
func main() {

	a := [][]int{{1,3,1},{1,5,1},{4,2,1}}
	fmt.Println(minPathSum(a))
	
}
