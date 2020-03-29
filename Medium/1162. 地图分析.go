package main

import (
	"fmt"
)

type point struct {
	x int
	y int
}

func maxDistance(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	count := 0
	group := []*point{}
	for i:= 0 ; i < len(grid) ; i++{
		for j := 0 ; j < len(grid[0]) ; j++{
			if grid[i][j] == 1{
				group = append(group,&point{i,j})
			}
		}
	}
	if len(group) == 0 || len(group) == len(grid)*len(grid[0]){
		return -1
	}
	dx := []int{1,0,-1,0}
	dy := []int{0,1,0,-1}
	for len(group) != 0{

		count++
		temp := group
		group = []*point{}
		for _,val := range temp {
			i := val.x
			y := val.y
			for q := 0; q < 4; q++ {
				newx := i + dx[q]
				newy := y + dy[q]

				if newx >= 0 && newx < len(grid) && newy >= 0 && newy < len(grid[0]) && grid[newx][newy] == 0 {
					//fmt.Println(newx,newy)
					grid[newx][newy] = 1
					group = append(group, &point{newx, newy})
				}
				//time.Sleep(2*time.Second)
			}
			//fmt.Println(len(group))
		}
	}
	return count -1
}
func main() {

	a := [][]int{{1,0,1},{0,0,0},{1,0,1}}
	fmt.Println(maxDistance(a))
}
